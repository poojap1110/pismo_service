package api

import (
	"errors"
	"net/http"
	"$GOPATH/pismo-service/resource"
	"time"

	"bitbucket.org/matchmove/integration-svc-aub/app/errs"

	"bitbucket.org/matchmove/integration-svc-aub/app/resource/api/response"
	"bitbucket.org/matchmove/integration-svc-aub/model"
	"bitbucket.org/matchmove/integration-svc-aub/modules/container"
	"bitbucket.org/matchmove/integration-svc-aub/modules/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/matchmove/rest.v2"
)

// Account API resource
type Account struct {
	resource.Resource
}

// Account API resource
type Accounts struct {
	resource.Resource
}

// NewAccountsResource ...
func NewAccountsResource(db *sqlx.DB, cont *container.Container) (string, string, func() rest.ResourceType) {
	return resource.AccountsResource,
		resource.AccountResourceEndpoint,
		func() rest.ResourceType {
			return &Account{resource.New(resource.AccountsResource, true, db, cont)}
		}

}

//Post sends get request to /Account
func (r *Account) Post() {
	var (
		req             = account.Request{}
		res             = account.Response{}
		accountModel, _ = model.NewAccount(r.DB)
	)

	if err := r.GetParams(&req); err != nil {
		r.Log.Print("[PostPartners] Invalid body/parameters")
		r.FormatException(r, errors.New(errs.ErrRequestBodyInvalid))
		return
	}

	// =========== Validating fields =========== //
	if err := helper.Validate(req, r.DB, ""); err != nil {
		r.Log.Print("[PostPartners] Error Validating Fields : ", err.Error())
		r.FormatException(r, err)
		return
	}

	_, err := accountModel.Create()
	if err != nil {
		r.Log.Print("[PostPartners] Error Creating Partner : ", err.Error())
		r.FormatException(r, err)
		return
	}

	// response ...
	res.Status = "Success"

	r.Status = http.StatusOK
	r.RawBody = res
}

//Get sends get request to /Account
func (r *Account) Get() {

	var (
		timestampResponse = response.TimeStampResponse{}
	)
	timestampResponse.SystemTimeStamp = string(time.Now().Format(time.RFC3339))

	r.Status = http.StatusOK
	r.RawBody = timestampResponse
}
