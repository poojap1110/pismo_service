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
type Heartbeat struct {
	resource.Resource
}

// Heartbeat API resource
type Heartbeats struct {
	resource.Resource
}

// NewHeartbeat ...
func NewHeartbeat(db *sqlx.DB, cont *container.Container) (string, string, func() rest.ResourceType) {
	return resource.HeartbeatResource,
		resource.HeartbeatEndpoint,
		func() rest.ResourceType {
			return &Heartbeat{resource.New(resource.HeartbeatResource, true, db, cont)}
		}

}

//Get sends get request to /heartbeat
func (r *Heartbeat) Get() {

	var (
		timestampResponse = response.TimeStampResponse{}
	)
	timestampResponse.SystemTimeStamp = string(time.Now().Format(time.RFC3339))

	r.Status = http.StatusOK
	r.RawBody = timestampResponse
}
