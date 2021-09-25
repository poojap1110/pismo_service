package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"bitbucket.org/matchmove/go-tools/array"
	"bitbucket.org/matchmove/integration-svc-aub/modules/constant"

	"bitbucket.org/liamstask/goose/lib/goose"
	"bitbucket.org/matchmove/logs"
)

var (
	driver string
	dbstr  string
	migdir string
	cmd    string
	bin    []string
	log    *logs.Log
)

func main() {

	log = logs.New()
	cmds := []string{"migrate", "create", "status"}
	driver = os.Getenv(constant.EnvDbDriver)
	migdir = os.Getenv(constant.EnvDbMigrationDir)
	dbstr = os.Getenv(constant.EnvDbOpen)
	migdir = migdir + string(os.PathSeparator)

	if migdir == "" || driver == "" {
		log.Fatal("Missing Env variable")
	}

	if len(os.Args) <= 1 {
		log.Fatal("Missing comment")
	}

	cmd = os.Args[1]

	if exist, _ := array.InArray(cmd, cmds); !exist {
		log.Fatal("Invalid comment")
	}

	if cmd == "create" {
		CreateMigration()
	}

	if cmd == "migrate" {
		RunMigration()
	}

	if cmd == "status" {
		StatusRun()
	}

}

// CreateMigration - function to create the migration file
func CreateMigration() {
	args := os.Args
	if len(args) <= 2 {
		log.Fatal("goose create: migration name required")
	}
	migrationType := "sql" // default to Go migrations
	if len(args) > 3 {
		migrationType = args[3]
	}

	if err := os.MkdirAll(migdir, 0777); err != nil {
		log.Print(migdir)
		log.Fatal(err)
	}

	n, err := goose.CreateMigration(args[2], migrationType, migdir, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	a, e := filepath.Abs(n)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println("goose: created", a)
}

// RunMigration - Run the migrations script availbale inside the migration folder
func RunMigration() {

	var target int64

	conf, err := GetDbConnections()
	if err != nil {
		log.Fatal(err)
	}

	current, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
	if err != nil {
		log.Fatal(err)
	}

	target = current

	args := os.Args
	if len(args) > 2 && args[2] == "down" {
		previous, err := goose.GetPreviousDBVersion(conf.MigrationsDir, target)
		if err != nil {
			log.Fatal(err)
		}
		target = previous
	}

	if err := goose.RunMigrations(conf, conf.MigrationsDir, target); err != nil {
		log.Fatal(err)
	}

}

// StatusRun - Get the current status of the database migration
func StatusRun() {
	conf, err := GetDbConnections()
	if err != nil {
		log.Fatal(err)
	}

	// collect all migrations
	min := int64(0)
	max := int64((1 << 63) - 1)
	migrations, e := goose.CollectMigrations(conf.MigrationsDir, min, max)
	if e != nil {
		log.Fatal(e)
	}

	db, e := goose.OpenDBFromDBConf(conf)
	if e != nil {
		log.Fatal("couldn't open DB:", e)
	}
	defer db.Close()

	// must ensure that the version table exists if we're running on a pristine DB
	if _, e := goose.EnsureDBVersion(conf, db); e != nil {
		log.Fatal(e)
	}

	fmt.Printf("goose: status for environment '%v'\n", conf.Env)
	fmt.Println("    Applied At                  Migration")
	fmt.Println("    =======================================")
	for _, m := range migrations {
		PrintMigrationStatus(db, m.Version, filepath.Base(m.Source))
	}
}

// GetDbConnections - get the db connection based on the db information from the env variable
func GetDbConnections() (*goose.DBConf, error) {
	d := goose.DBDriver{
		Name:    driver,
		OpenStr: dbstr,
		Dialect: &goose.MySqlDialect{},
	}
	d.Import = "github.com/go-sql-driver/mysql"
	if !d.IsValid() {
		return nil, fmt.Errorf("Invalid DBConf: %v", d)
	}

	conf := &goose.DBConf{
		Env:           "default",
		MigrationsDir: migdir,
		Driver:        d,
	}
	return conf, nil
}

// PrintMigrationStatus ...
func PrintMigrationStatus(db *sql.DB, version int64, script string) {
	var row goose.MigrationRecord
	q := fmt.Sprintf("SELECT tstamp, is_applied FROM goose_db_version WHERE version_id=%d ORDER BY tstamp DESC LIMIT 1", version)
	e := db.QueryRow(q).Scan(&row.TStamp, &row.IsApplied)

	if e != nil && e != sql.ErrNoRows {
		log.Fatal(e)
	}

	var appliedAt string

	if row.IsApplied {
		appliedAt = row.TStamp.Format(time.ANSIC)
	} else {
		appliedAt = "Pending"
	}

	fmt.Printf("    %-24s -- %v\n", appliedAt, script)
}
