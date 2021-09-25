package resource_test

import (
	"os"
	"testing"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/app/resource"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	db "bitbucket.org/matchmove/go-memcached-database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestResource_New(t *testing.T) {
	r := resource.New("Resource", false, nil, nil, nil)
	assert.IsType(t, resource.Resource{}, r, "")
}

func ResourceTestDB() map[string]*db.DB {
	return db.Databases{
		"sgmock": &db.DB{
			Driver: os.Getenv(constant.EnvDbDriver),
			Open:   os.Getenv(constant.EnvDbOpen),
		},
		"default": &db.DB{
			Driver: os.Getenv(constant.EnvDbDriver),
			Open:   os.Getenv(constant.EnvDbOpen),
		},
		"mock": &db.DB{
			Driver: "testdb",
			Open:   "",
		},
	}
}
