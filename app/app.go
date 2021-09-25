package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"bitbucket.org/matchmove/go-tools/secure"
	"bitbucket.org/matchmove/integration-svc-aub/app/errs"
	"bitbucket.org/matchmove/integration-svc-aub/app/resource"
	"bitbucket.org/matchmove/integration-svc-aub/app/resource/api"
	"bitbucket.org/matchmove/integration-svc-aub/modules/constant"
	"bitbucket.org/matchmove/integration-svc-aub/modules/container"

	"/home/pooja/git/pismo-service/modules/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/streadway/amqp"
	rest "gopkg.in/matchmove/rest.v2"

	"github.com/newrelic/go-agent/_integrations/nrgorilla/v1"
)

const (
	AuthHeader = "Authorization"
)

// App represent the application
type App struct {
	Server       *rest.Server
	DB           *sqlx.DB
	Refs         Refs
	Domain       Domain
	Environment  Environment
	AccessLog    AccessLog
	Routes       rest.Routes
	Cont         *container.Container
	Driver, Open string
	Connection   *sqlx.DB
}

// Link ...
type Link struct {
	Rel  string `json:"rel"`
	HREF string `json:"href"`
}

// Refs represents the application references
type Refs struct {
	Docs string
}

// Domain usually host and port for this server
type Domain struct {
	Host string
}

// AccessLog path of access log file
type AccessLog struct {
	Name string
}

// Environment environment variables
type Environment struct {
	Env string
}

var (
	app          App
	connPublish  *amqp.Connection
	connConsumer *amqp.Connection
)

const (
	// Realm defines this App's Authentication realm
	Realm = "MatchMove Pay %s"

	// WelcomeMessage defines the App's root message
	WelcomeMessage = "Welcome to %s!"

	LogConnection = "RabbitMq_Connections"
)

// RequiredEnvironmentVarNotDefined checking for required environment that is either empty or has not been defined
func RequiredEnvironmentVarNotDefined(key string) (string, error) {
	var err error
	var val string

	if val = os.Getenv(key); val == "" {
		err = fmt.Errorf("Failed to initialize application server: %s is missing", key)
	}

	return val, err
}

// NewApplicationServer creates new application server instance
func NewApplicationServer() (App, error) {
	var (
		domain, env, accessLog, docs string

		err error
	)

	a := App{}

	for _, key := range constant.RequiredEnvironmentVars {
		var value string
		if value, err = RequiredEnvironmentVarNotDefined(key); err != nil {
			return a, err
		}

		value = strings.TrimSpace(value)
		if value == "" {
			return a, fmt.Errorf("%s should not be empty", key)
		}

		switch key {
		case constant.EnvAppDomain:
			domain = value
			break
		case constant.EnvAppEnvironment:
			env = value
			break
		case constant.EnvAppRefDocs:
			docs = value
			break
		case constant.EnvAppAccessLog:
			accessLog = value
			break
		}
	}

	a.Refs = Refs{
		Docs: docs,
	}

	a.Domain = Domain{
		Host: domain,
	}

	a.Connect()

	a.AccessLog = AccessLog{
		Name: accessLog,
	}

	a.Environment = Environment{
		Env: env,
	}

	logger.GetInstance(a.Cont, LogConnection).WriteLog([]byte("\nRabbitMq publisher and consumer connections Established"))

	return a, nil
}

// Migrate ...
func (a *App) Migrate() (err error) {

	return
}

// Connect establishes a database connection using the set credentials
// from the DB struct
func (a *App) Connect() error {
	db, err := sqlx.Connect(os.Getenv(constant.EnvDbDriver), os.Getenv(constant.EnvDbOpen))

	if err != nil {
		return fmt.Errorf("DB Connect error: %s", err)
	}

	// using Unsafe so destination struct fields
	// don't have to be exactly equal to all the fields in the table being queried
	a.Connection = db.Unsafe()
	return nil
}

// GetApplication get current application server instance
func GetApplication() App {
	return app
}

// GetDomain gets the app's domain
func (a *App) GetDomain() string {
	return a.Domain.Host
}

// RootHandler response to the URI request "/"
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", rest.ContentTypeJSON)

	json.NewEncoder(w).Encode(struct {
		Documentation string `json:"ref"`
	}{
		app.Refs.Docs,
	})
}

// SetRoutes sets the application's route config
func (a *App) SetRoutes() {
	// a.Routes = rest.NewRoutes()
}

// Jobs run background task
func (a *App) Jobs() {
	logger.GetInstance(a.Cont, LogConnection).WriteLog([]byte("\nConsumer initialization started"))
}

// PrepareRoutes populate all routes including root handler and default
func (a *App) PrepareRoutes() {

	if a.Server != nil {
		docsPrefix := resource.DocPath
		docsHandler := http.StripPrefix(docsPrefix, http.FileServer(http.Dir(os.Getenv(constant.EnvSwaggerPath))))
		a.Server.SetRoutes(mux.NewRouter().StrictSlash(true), rest.NewRoutes().
			Add(api.NewAccountsResource(a.DB, a.Cont)).
			Add(api.NewTransactionResource(a.DB, a.Cont)).
			Add(api.NewHeartbeat(a.DB, a.Cont)).
			Add(api.NewPismoErrors(a.DB, a.Cont)).
			Root(RootHandler).
			NotFound(a.DefaultNotFoundRouteHandler))

		if os.Getenv("APP_ENV") != "PRODUCTION" {
			a.Server.Router.PathPrefix(docsPrefix).Handler(docsHandler)
		}
		nrgorilla.InstrumentRoutes(a.Server.Router, nil)
	}
}

// DefaultNotFoundRouteHandler is the initial 404 route handler
func (a *App) DefaultNotFoundRouteHandler(w http.ResponseWriter, r *http.Request) {

	path := os.Getenv("APP_EXTERNAL_DOMAIN") + "/v1/error/"
	w.Header().Set("Content-Type", rest.ContentTypeJSON)
	json.NewEncoder(w).Encode(struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Link    []Link `json:"link"`
	}{
		http.StatusNotFound,
		"Resource Not found",
		[]Link{
			{
				Rel:  errs.ErrResourceNotFound,
				HREF: path + errs.ErrResourceNotFound,
			},
		},
	})
}

// Run start and serve
func Run(a *App) error {
	var (
		err error
	)

	aLog, _ := ioutil.TempFile("", "")
	defer func() {
		aLog.Close()
		os.Remove(aLog.Name())
	}()

	if a.Server, err = rest.NewServer(a.GetDomain()); err != nil {
		return err
	}

	log.Println("Starting server ", a.Server.URL.String())
	log.Println("Enviroment ", os.Getenv(constant.EnvAppEnvironment))

	a.Jobs()
	a.PrepareRoutes()

	a.Server.Handler = handlers.LoggingHandler(
		aLog,
		func(m *mux.Router) http.Handler {
			return m
		}(a.Server.Router),
	)

	app = *a

	if err = a.Server.Listen(); err != nil {
		return err
	}

	return nil
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

// GetClientIP returns the client's IP
func GetClientIP(r *http.Request) (string, error) {
	ipAddress := r.RemoteAddr
	if ip := r.Header.Get("X-Forwarded-For"); "" != ip {
		ipAddress = ip
		// X-Forwarded-For might contain multiple IPs. Get the last one.
		if strings.Contains(ipAddress, ",") {
			ips := strings.Split(ipAddress, ",")
			ipAddress = strings.Trim(ips[len(ips)-1], " ")
		}
	}
	var (
		ip  net.IP
		err error
	)
	if -1 != strings.Index(ipAddress, ":") {
		if ipAddress, _, err = net.SplitHostPort(ipAddress); nil != err {
			return "", fmt.Errorf("GetClientIP SplitHostPort Error: %v", err)
		}
	}
	if err := ip.UnmarshalText([]byte(ipAddress)); nil != err {
		return "", fmt.Errorf("GetClientIP UnmarshalText Error: %v", err)
	}
	return ipAddress, nil
}

// Identify identifier for logs
func Identify(r *http.Request) (identifier string) {

	identifier = fmt.Sprintf("%s\t", secure.MD5(time.Now().UTC().Format("2021-09-25 15:04:05")+
		r.URL.String()+
		r.Method))
	return identifier
}
