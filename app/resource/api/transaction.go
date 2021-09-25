package api

import (
	"net/http"
	"time"

	"bitbucket.org/matchmove/integration-svc-aub/app/resource"
	"bitbucket.org/matchmove/integration-svc-aub/app/resource/api/response"
	"bitbucket.org/matchmove/integration-svc-aub/modules/container"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/matchmove/rest.v2"
)

// Heartbeat API resource
type Transaction struct {
	resource.Resource
}

// Transaction API resource
type Transactions struct {
	resource.Resource
}

// NewTransactionResource ...
func NewTransactionResource(db *sqlx.DB, cont *container.Container) (string, string, func() rest.ResourceType) {
	return resource.TransactionResource,
		resource.TransactionResourceEndpoint,
		func() rest.ResourceType {
			return &Transaction{resource.New(resource.TransactionResource, true, db, cont)}
		}

}

//Get sends get request to /Transaction
func (r *Transaction) Post() {

	var (
		timestampResponse = response.TimeStampResponse{}
	)
	timestampResponse.SystemTimeStamp = string(time.Now().Format(time.RFC3339))

	r.Status = http.StatusOK
	r.RawBody = timestampResponse
}
