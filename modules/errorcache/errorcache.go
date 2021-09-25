package errorcache

import (
	"sync"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/model"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/jmoiron/sqlx"
)

type errorsCache struct {
	sync.Mutex
	bag map[string]model.Error
}

const (
	InstanceKey = "Errors"
)

// GetRegistry function ...
func GetRegistry() container.Registries {
	obj := &errorsCache{
		bag: make(map[string]model.Error),
	}

	return container.Registries{
		container.Registry{
			Key:   InstanceKey,
			Value: obj,
		},
	}
}

// GetInstance function ...
func GetInstance(c *container.Container) *errorsCache {
	return c.Get(InstanceKey).(*errorsCache)
}

// PopulateErrorCodes function ...
func PopulateErrorCodes(c *container.Container, db *sqlx.DB, cache *memcache.Client) {
	ins := GetInstance(c)
	ins.RetrieveErrors(c, db, cache)
	c.StoreToGlobal(InstanceKey, ins)
}

// GetError method - gets the stored error object on Container
func (me *errorsCache) GetError(errorCode string) *model.Error {
	var mError = me.bag[errorCode]

	if mError.Id == 0 {
		return nil
	}

	return &mError
}

// RetrieveErrors method ...
func (me *errorsCache) RetrieveErrors(c *container.Container, db *sqlx.DB, cache *memcache.Client) {
	mError, _ := model.NewError(db, cache, nil)
	errs, _ := mError.Find()
	newBag := make(map[string]model.Error)

	for _, v := range errs {
		newBag[v.Code.String] = v
	}

	me.bag = newBag
}