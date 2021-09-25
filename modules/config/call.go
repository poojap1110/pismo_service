package config

import (
	"encoding/json"
	"fmt"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/model"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/cache"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/jmoiron/sqlx"
)

const (
	// Config service related
	ResourceApiVersion    = "x1.0.0.0"
	ResourceHeaderVersion = 3
)

// callGetConfig method - calls GetConfig endpoint of Config Microservice and returns values fetched.
func (me *config) callGetConfig(c *container.Container, filter ConfigFilter, configType string, db *sqlx.DB,
	mcache *memcache.Client) (ConfigFilter, error) {

	var (
		err          error
		mConfig, _ = model.NewConfig(db, mcache, nil)
		configs      map[string]map[string]string
		rCache      *cache.Rediscache
		key          string
		cStore     []byte
	)

	if configType == constant.TenantTypeConfig {
		key = cache.TenantConfigSuffix + filter.TenantHashID
	} else if configType == constant.TechnicalTypeConfig {
		key = cache.TechnicalConfigSuffix + filter.Owner
	}

	if rCache = cache.GetInstance(c); rCache != nil {

		if i, err := rCache.Get(key); err == nil {
			fmt.Println("load "+key+" configs from cache...")
			cStore = i.([]byte)
			// covert to interface
			err = json.Unmarshal(cStore, &configs)
		} else {
			fmt.Println("load "+key+" configs from db...")
			if configs, err = mConfig.GetActiveConfig(filter.TenantHashID, filter.Owner, filter.Name); err != nil { // map[string]map[string]string
				return filter, err
			} else {
				// store in recent redis cache key

				cStore, _ = json.Marshal(configs)
				err = rCache.Set(key, string(cStore), 0)
				if err != nil {
					return filter, err
				}
			}
		}

	} else {

		if configs, err = mConfig.GetActiveConfig(filter.TenantHashID, filter.Owner, filter.Name); err != nil { // map[string]map[string]string
			return filter, err
		}

	}

	if len(configs) > 0 {
		for cName, cValue := range configs {
			filter.Configs = append( filter.Configs,ConfigItem{
				Name: cName,
				Value: cValue["Value"],
				DefaultValue: cValue["DefaultValue"],
				FriendlyName: cValue["FriendlyName"],
			})
		}
	}

	return filter, nil
}