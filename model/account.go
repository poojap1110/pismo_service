package model

import (
	"database/sql"
	"fmt"
	"time"

	gorm "bitbucket.org/matchmove/go-memcached-model"
	"bitbucket.org/matchmove/integration-svc-aub/app/errs"
	"github.com/go-errors/errors"
	"github.com/jmoiron/sqlx"
)

// Account ...
type Account struct {
	gorm.Model     `db:"-" json:"-"`
	ID             int64  `db:"id" json:"id"`
	AccountID      string `db:"account_id" json:"account_id"`
	DocumentNumber string `db:"name" json:"document_number"`
	DateAdded      []byte `db:"date_added"`
}

// Accounts ...
type Accounts []Account

// FillableAccountFields ...
var FillableAccountFields = []string{
	"id",
	"document_number",
	"date_added",
}

// NewAccount ...
func NewAccount(db *sqlx.DB) (*Account, error) {
	if db == nil {
		return nil, errors.New(gorm.SQLNoDatabaseConnection)
	}

	return &Account{Model: gorm.Model{DB: db}}, nil
}

// Create ....
func (me *Account) Create() (result sql.Result, err error) {
	if string(me.DateAdded) == "" {
		me.DateAdded = []byte(time.Now().UTC().Format(gorm.SQLDatetime))
	}

	query := me.GenericInsertSQL(FillableAccountFields...)
	result, err = me.exec(query)

	if err != nil {
		err = errors.New(errs.ErrInternalDBError + ": " + err.Error())
		return nil, err
	}

	return
}

// GetAllAccount method - Get all record
func (me *Account) GetAllAccount() (Accounts Accounts, err error) {
	me.OrderBy([]string{"date_added DESC"})
	Accounts, err = me.Find()

	return
}

// GetAccount ...
func (me *Account) GetAccount() error {
	var conditions []gorm.Condition
	conditions = append(conditions,
		gorm.Condition{
			Field:    "account_id",
			Operator: "=",
			Value:    me.AccountID,
		},
	)

	me.Where(conditions...)

	if err := me.FindFirst(); err != nil {
		if err.Error() == gorm.SQLNoRowsErrorCode {
			return nil
		}

		return errors.New(errs.ErrInternalDBError + ": " + err.Error())
	}

	return nil
}

// Find returns the list of records meeting the requirement filters, limit, offset, and order
func (me *Account) Find() (Accounts, error) {
	query, args := me.PrepareStatement(me, "?")
	return me.gets(`SELECT * FROM %s `+query, args...)
}

// FindFirst returns first record meeting the requirement filters, limit, offset, and order
func (me *Account) FindFirst() error {
	me.Limit(1)
	query, _ := me.PrepareStatement(me, "")
	return me.get(`SELECT * FROM %s ` + query)
}

// exec method - calls format method and executes query to database
func (me *Account) exec(sql string) (result sql.Result, err error) {
	result, err = me.DB.NamedExec(me.sql(sql), me)
	return
}

// get method - Gets the result from database basing from sql received.
func (me *Account) get(sql string) (err error) {
	return me.CacheGet(me, me.sql(sql))
}

// gets method - Get the results from database basing from sql and arguments received.
func (me *Account) gets(sql string, args ...interface{}) (response Accounts, err error) {
	var rs Accounts

	if err = me.CacheGets(&rs, me.sql(sql), args...); err != nil {
		response = Accounts{}
		return
	}

	for _, item := range rs {
		item.DB = me.DB
		response = append(response, item)
	}

	return response, err
}

// sql method - converts sql query to use proper table name.
func (me *Account) sql(s string) string {
	return fmt.Sprintf(s, "Accounts")
}
