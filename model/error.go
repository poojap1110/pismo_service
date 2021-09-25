package model

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"bitbucket.org/matchmove/go-memcached-model"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/entity"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-errors/errors"
	"github.com/jmoiron/sqlx"
)

// Error struct for errors
type Error struct {
	gorm.Model           `db:"-" json:"-"`
	Id                   int            `db:"id" json:"-"`
	Code                 sql.NullString `db:"code" json:"-"`
	Description          sql.NullString `db:"description" json:"-"`
	Status               sql.NullInt64  `db:"status" json:"-"`
	FormattedCode        string         `db:"-" json:"code"`
	FormattedDescription string         `db:"-" json:"description"`
	Links                []entity.Link  `db:"-" json:"links"`
}

type Errors []Error

const (
	tableNameError = "error"

	defaultResourceNotFoundErrorCode        = constant.ResourceNotFound
	defaultResourceNotFoundErrorDescription = "Resource Not Found"
	defaultResourceNotFoundErrorStatus      = http.StatusNotFound

	defaultErrorCode        = constant.InternalServerError
	defaultErrorDescription = "Internal Server Error"
	defaultErrorStatus      = http.StatusInternalServerError
)

var (
	FillableErrorFields = []string{
		"code",
		"description",
		"status",
	}

	UpdatableErrorFields = []string{
		"code",
		"description",
		"status",
	}
)

// NewError function ...
func NewError(db *sqlx.DB, cache *memcache.Client, timezone *string) (*Error, error) {
	if db == nil {
		return nil, errors.New(gorm.SQLNoDatabaseConnection)
	}

	return &Error{Model: gorm.Model{DB: db, Cache: cache, Timezone: timezone}}, nil
}

// RetrieveDefaultError method - sets 500 internal server error.
func (me *Error) RetrieveDefaultError() {
	me.Code = sql.NullString{
		String: defaultErrorCode,
		Valid:  true,
	}
	me.Description = sql.NullString{
		String: defaultErrorDescription,
		Valid:  true,
	}
	me.Status = sql.NullInt64{
		Int64: defaultErrorStatus,
		Valid: true,
	}
	me.format()
}

// RetrieveDefaultError method - sets 404 resource not found.
func (me *Error) RetrieveDefaultResourceNotFoundError() {
	me.Code = sql.NullString{
		String: defaultResourceNotFoundErrorCode,
		Valid:  true,
	}
	me.Description = sql.NullString{
		String: defaultResourceNotFoundErrorDescription,
		Valid:  true,
	}
	me.Status = sql.NullInt64{
		Int64: defaultResourceNotFoundErrorStatus,
		Valid: true,
	}
	me.format()
}

// RetrieveErrorByCode method ...
func (me *Error) RetrieveErrorByCode(code string) (err error) {
	if err = me.Where(gorm.Condition{
		Field:    "code",
		Operator: gorm.OperatorEqual,
		Value:    code,
	}); err != nil {
		return err
	}

	if err = me.FindFirst(); err != nil {
		return err
	}

	return nil
}

// Find returns the list of records meeting the requirement filters, limit, offset, and order
func (me *Error) Find() (Errors, error) {
	query, args := me.PrepareStatement(me, "?")
	return me.gets(`SELECT * FROM %s `+query, args...)
}

// FindFirst returns first record meeting the requirement filters, limit, offset, and order
func (me *Error) FindFirst() error {
	me.Limit(1)
	query, _ := me.PrepareStatement(me, "")
	return me.get(`SELECT * FROM %s ` + query)
}

// exec method - calls format method and executes query to database
func (me *Error) exec(sql string) (result sql.Result, err error) {
	result, err = me.DB.NamedExec(me.sql(sql), me)
	me.format()
	return
}

// format method - formats object properties
func (me *Error) format() {
	uri, _ := url.Parse(os.Getenv(constant.EnvExternalAppDomain))
	path := "https://" + uri.Hostname() + "/error/"
	me.FormattedCode = me.Code.String
	me.FormattedDescription = me.Description.String
	me.Links = []entity.Link{entity.NewLink(path+me.Code.String, "error "+me.Code.String)}
}

// get method - Gets the result from database basing from sql received.
func (me *Error) get(sql string) (err error) {
	if err = me.CacheGet(me, me.sql(sql)); err == nil {
		me.format()
	}

	return
}

// gets method - Get the results from database basing from sql and arguments received.
func (me *Error) gets(sql string, args ...interface{}) (response Errors, err error) {
	var rs Errors

	if err = me.CacheGets(&rs, me.sql(sql), args...); err != nil {
		response = Errors{}
		return
	}

	for _, item := range rs {
		item.DB = me.DB
		item.Timezone = me.Timezone
		item.format()
		response = append(response, item)
	}

	return response, err
}

// sql method - converts sql query to use proper table name.
func (me *Error) sql(s string) string {
	return fmt.Sprintf(s, tableNameError)
}
