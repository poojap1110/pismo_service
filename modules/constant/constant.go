package constant

// Required environment variables
const (
	// EnvAppDomain application domain
	// APP_DOMAIN - port to listen to. example: https://0.0.0.0:433
	EnvAppDomain = "APP_DOMAIN"

	// APP_NAME application name
	EnvAppName = "APP_NAME"

	// EnvAppRefDocs application document URL
	// APP_REF_DOCS - URL for references documentation when an error occurs
	EnvAppRefDocs = "APP_REF_DOCS"

	// EnvAppEnvironment application environment
	// APP_ENV - server environment configuration. [DEVELOPMENT, TESTING, PRODUCTION]
	EnvAppEnvironment = "APP_ENV"

	// EnvAppAccessLog access log filename
	// APP_ACCESS_LOG - access log file path.
	EnvAppAccessLog = "APP_ACCESS_LOG"

	// EnvAppDefaultBucket - Bucket where the config are kept.
	// AWS_DEFAULT_BUCKET - Bucket where the config are kept.
	EnvAppDefaultBucket = "AWS_DEFAULT_BUCKET"

	// EnvAppDefaultFolder - Folder where the config are kept locally.
	// AWS_DEFAULT_BUCKET - Folder where the config are kept locally.
	EnvAppDefaultFolder = "APP_CONFIG_FOLDER"

	// EnvAppRegion ...
	EnvAppRegion = "AWS_REGION"

	// EnvAppAmazonAccessKey ...
	EnvAppAmazonAccessKey = "AWS_SECRET_ACCESS_KEY"

	// EnvAppAmazonAccessID ...
	EnvAppAmazonAccessID = "AWS_ACCESS_KEY_ID"

	EnvDbMigrationDir = "DB_MIGRATION_DIR"

	EnvDbDriver = "DB_DRIVER"

	EnvDbOpen = "DB_OPEN"

	EnvCache = "CACHE_SERVER"

	EnvAppVersion = "APP_VERSION"

	// EnvMemcachedServer ...
	EnvMemcachedServer = "MEMCACHED_SERVER"

	// EnvMemcachedConfig ...
	EnvMemcachedConfig = "MEMCACHED_CONFIG"

	// EnvApiGateway ...
	EnvApiGateway = "API_GATEWAY"

	// EnvAppLogFolder ...
	EnvAppLogFolder = "LOG_FOLDER"

	// === AUB - OP PROXY === //
	EnvOPProxyURL                  = "OP_PROXY_URL"
	ENVOPADDRProxyURL              = "OP_ADDR_PROXY_URL"
	EnvOPProxyVersion              = "OP_PROXY_VERSION"
	EnvOPRestDocumentationErrorURL = "OP_REST_DOCUMENTATION_ERROR_URL"
	EnvProductAndServices          = "PRODUCT_AND_SERVICES"
	EnvMonthlyVolumeTxn            = "MONTHLY_VOLUME_TXN"
	EnvDefaultMethod               = "DEFAULT_ACCT_METHOD"
	EnvDefaultRemarks              = "DEFAULT_REMARKS"
)

const (
	OpProxyKey = "OP_KEY"

	OpProxySecret = "OP_SECRET"

	OpProxyServiceHost = "OP_HOST"

	OpProxyAddrHost = "OP_PUT"

	OPADDRENPOINT = "OP_ADDR"

	AUBHost = "AUB_HOST"

	EnvSwaggerPath = "SWAGGER_PATH"

	EnvRabbitServer = "RABBITMQ_SERVER"

	EnvMaxRetries        = "MAX_RETRIES"
	EnvMaxTopicDelayTIME = "MAX_DELAY"
	EnvMaxRetriesTopic   = "MAX_RETRIES_TOPIC"
	EnvMaxDelayTIME      = "MAX_DELAY_TIME"
	EnvExchange          = "EXCHANGE_NAME"
	EnvDelayExchange     = "DELAY_EXCHANGE_NAME"
	EnvOpQueue           = "OP_QUEUE_NAME"
	EnvAubQueue          = "AUB_QUEUE_NAME"
)

// Cron Jobs --
const (
	JobEnabled = "1"

	EnvKeyRefetchConfigEnabled  = "JOB_REFETCH_CONFIG_ENABLED"
	EnvKeyRefetchConfigDuration = "JOB_REFETCH_CONFIG_DURATION"
)

// Application information/identifier
const (
	// Name application name
	Name = "svc-aub"

	// Realm application's realm
	Realm = "svc-aub"

	// Version current application version
	Version = "0.1"

	// Domain application's domain
	Domain = "http://0.0.0.0:443"

	// Username username used for testing
	Username = "tester"

	// Password password used for testing
	Password = "password"
)

// RequiredEnvironmentVars these are environment variables that is needed by the application
var RequiredEnvironmentVars = []string{
	EnvAppDomain,
	EnvAppRefDocs,
	EnvAppEnvironment,
	EnvAppAccessLog,
	EnvDbDriver,
	EnvDbOpen,
	EnvMemcachedServer,
	EnvMemcachedConfig,
	EnvAppLogFolder,
	EnvOPProxyURL,
	EnvOPProxyVersion,
}

const (
	LayoutDate         = "2006-01-02"
	LayoutDateTime8601 = "2006-01-02T15:04:05-07:00"
)
