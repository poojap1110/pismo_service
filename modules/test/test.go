package test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"bitbucket.org/matchmove/go-memcached-database"
	"bitbucket.org/matchmove/logs"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/app/middleware"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/config"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/errorcache"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/jmoiron/sqlx"
)

var (
	conn = make(map[string]*sqlx.DB)
	log   = logs.New()
	cache *memcache.Client
	cont  *container.Container
)

// GetEnvironmentDB ...
func GetEnvironmentDB() (mysqlConn map[string]*sqlx.DB, cache *memcache.Client, cont *container.Container) {
	var (
		userDB *db.DB
		err    error
	)
	_ = os.Setenv(constant.EnvMockConfigs, "1")
	if conn == nil || len(conn) == 0 {
		userDB, err = db.New(os.Getenv(constant.EnvDbDriver), os.Getenv(constant.EnvDbOpen))
		if err != nil {
			panic(err)
		}
		conn[constant.DbBankAccount] = userDB.Connection
	}

	mysqlConn = conn
	cache = middleware.GetMemcache()
	mockedGlobalContainer := container.New().
		Register(config.GetRegistry()).
		Register(errorcache.GetRegistry())
	cont = mockedGlobalContainer.Duplicate()
	config.Populate(cont)
	errorcache.PopulateErrorCodes(cont, mysqlConn[constant.DbBankAccount], cache)
	return
}

// ConvertToRequestPayload ...
func ConvertToRequestPayload(m interface{}) string {
	data, _ := json.Marshal(m)
	return string(data)
}

// ConvertToRequestQuery ...
func ConvertToRequestQuery(m interface{}) string {
	req := make(map[string]interface{})
	data, _ := json.Marshal(m)
	json.Unmarshal(data, &req)

	var slc []string
	for k, j := range req {
		switch v := j.(type) {
		case string:
			slc = append(slc, k+"="+v)
		case fmt.Stringer:
			slc = append(slc, k+"="+v.String())
		case float64:
			slc = append(slc, k+"="+fmt.Sprintf("%.f", v))
		default:
			slc = append(slc, k+"="+fmt.Sprintf("%v", v))
		}
	}

	return strings.Join(slc, "&")
}
