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

// OperationType ...
type OperationType struct {
	gorm.Model  `db:"-" json:"-"`
	ID          int64  `db:"id" json:"id"`
	OperationID string `db:"operation_id" json:"operation_id"`
	Description string `db:"description" json:"description"`
	DateAdded   []byte `db:"date_added"`
}

// OperationTypes ...
type OperationTypes []OperationType

// FillableOperationTypeFields ...
var FillableOperationTypeFields = []string{
	"id",
	"operation_id",
	"description",
	"date_added",
}

// NewOperationType ...
func NewOperationType(db *sqlx.DB) (*OperationType, error) {
	if db == nil {
		return nil, errors.New(gorm.SQLNoDatabaseConnection)
	}

	return &OperationType{Model: gorm.Model{DB: db}}, nil
}

// Create ....
func (me *OperationType) Create() (result sql.Result, err error) {
	if string(me.DateAdded) == "" {
		me.DateAdded = []byte(time.Now().UTC().Format(gorm.SQLDatetime))
	}

	query := me.GenericInsertSQL(FillableOperationTypeFields...)
	result, err = me.exec(query)

	if err != nil {
		err = errors.New(errs.ErrInternalDBError + ": " + err.Error())
		return nil, err
	}

	return
}

// GetAllOperationType method - Get all record
func (me *OperationType) GetAllOperationType() (OperationTypes OperationTypes, err error) {
	me.OrderBy([]string{"date_added DESC"})
	OperationTypes, err = me.Find()

	return
}

// GetOperationType ...
func (me *OperationType) GetOperationType() error {
	var conditions []gorm.Condition
	conditions = append(conditions,
		gorm.Condition{
			Field:    "operation_id",
			Operator: "=",
			Value:    me.OperationID,
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
func (me *OperationType) Find() (OperationTypes, error) {
	query, args := me.PrepareStatement(me, "?")
	return me.gets(`SELECT * FROM %s `+query, args...)
}

// FindFirst returns first record meeting the requirement filters, limit, offset, and order
func (me *OperationType) FindFirst() error {
	me.Limit(1)
	query, _ := me.PrepareStatement(me, "")
	return me.get(`SELECT * FROM %s ` + query)
}

// exec method - calls format method and executes query to database
func (me *OperationType) exec(sql string) (result sql.Result, err error) {
	result, err = me.DB.NamedExec(me.sql(sql), me)
	return
}

// get method - Gets the result from database basing from sql received.
func (me *OperationType) get(sql string) (err error) {
	return me.CacheGet(me, me.sql(sql))
}

// gets method - Get the results from database basing from sql and arguments received.
func (me *OperationType) gets(sql string, args ...interface{}) (response OperationTypes, err error) {
	var rs OperationTypes

	if err = me.CacheGets(&rs, me.sql(sql), args...); err != nil {
		response = OperationTypes{}
		return
	}

	for _, item := range rs {
		item.DB = me.DB
		response = append(response, item)
	}

	return response, err
}

// sql method - converts sql query to use proper table name.
func (me *OperationType) sql(s string) string {
	return fmt.Sprintf(s, "OperationTypes")
}
