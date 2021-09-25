package api_test

import (
	"os"

	db "bitbucket.org/matchmove/go-database"
	"bitbucket.org/matchmove/integration-svc-aub/modules/constant"
	"bitbucket.org/matchmove/integration-svc-aub/modules/container"
	"bitbucket.org/matchmove/integration-svc-aub/modules/curl"

	"github.com/jmoiron/sqlx"
)

var (
	conn *sqlx.DB
)

// GetEnvironmentDB ...
func GetEnvironmentDB() (mysqlConn *sqlx.DB, cont *container.Container) {
	var (
		mysqlDB *db.DB
		err     error
	)

	if conn == nil {

		if mysqlDB, err = db.New(os.Getenv(constant.EnvDbDriver), os.Getenv(constant.EnvDbOpen)); err != nil {
			return
		}

		conn = mysqlDB.Connection
	}

	mysqlConn = conn

	mockedGlobalContainer := container.New().
		Register(curl.GetRegistry())
	cont = mockedGlobalContainer.Duplicate()

	return
}
