package model

import (
	"database/sql"
	"fmt"

	gorm "bitbucket.org/matchmove/go-memcached-model"
	"bitbucket.org/matchmove/integration-svc-aub/app/errs"
	"github.com/go-errors/errors"
	"github.com/jmoiron/sqlx"
)

// Transaction ...
type Transaction struct {
	gorm.Model      `db:"-" json:"-"`
	ID              int64  `db:"id" json:"id"`
	AccountID       string `db:"account_id" json:"account_id"`
	TransactionID   string `db:"transaction_id" json:"transaction_id"`
	OperationTypeID string `db:"operation_type_id" json:"operation_type_id"`
	amount          string `db:"amount" json:"amount"`
	EventDate       string `db:"event_date" json:"event_date"`
}

// Transactions ...
type Transactions []Transaction

// FillableTransactionFields ...
var FillableTransactionFields = []string{
	"id",
	"account_id",
	"operation_type_id",
	"amount",
	"event_date",
}

// NewTransaction ...
func NewTransaction(db *sqlx.DB) (*Transaction, error) {
	if db == nil {
		return nil, errors.New(gorm.SQLNoDatabaseConnection)
	}

	return &Transaction{Model: gorm.Model{DB: db}}, nil
}

// Create ....
func (me *Transaction) Create() (result sql.Result, err error) {

	query := me.GenericInsertSQL(FillableTransactionFields...)
	result, err = me.exec(query)

	if err != nil {
		err = errors.New(errs.ErrInternalDBError + ": " + err.Error())
		return nil, err
	}

	return
}

// GetAllTransaction method - Get all record
func (me *Transaction) GetAllTransaction() (Transactions Transactions, err error) {
	me.OrderBy([]string{"date_added DESC"})
	Transactions, err = me.Find()

	return
}

// GetTransaction ...
func (me *Transaction) GetTransaction() error {
	var conditions []gorm.Condition
	conditions = append(conditions,
		gorm.Condition{
			Field:    "transaction_id",
			Operator: "=",
			Value:    me.TransactionID,
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
func (me *Transaction) Find() (Transactions, error) {
	query, args := me.PrepareStatement(me, "?")
	return me.gets(`SELECT * FROM %s `+query, args...)
}

// FindFirst returns first record meeting the requirement filters, limit, offset, and order
func (me *Transaction) FindFirst() error {
	me.Limit(1)
	query, _ := me.PrepareStatement(me, "")
	return me.get(`SELECT * FROM %s ` + query)
}

// exec method - calls format method and executes query to database
func (me *Transaction) exec(sql string) (result sql.Result, err error) {
	result, err = me.DB.NamedExec(me.sql(sql), me)
	return
}

// get method - Gets the result from database basing from sql received.
func (me *Transaction) get(sql string) (err error) {
	return me.CacheGet(me, me.sql(sql))
}

// gets method - Get the results from database basing from sql and arguments received.
func (me *Transaction) gets(sql string, args ...interface{}) (response Transactions, err error) {
	var rs Transactions

	if err = me.CacheGets(&rs, me.sql(sql), args...); err != nil {
		response = Transactions{}
		return
	}

	for _, item := range rs {
		item.DB = me.DB
		response = append(response, item)
	}

	return response, err
}

// sql method - converts sql query to use proper table name.
func (me *Transaction) sql(s string) string {
	return fmt.Sprintf(s, "Transactions")
}
