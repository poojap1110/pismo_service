package api

import (
	"net/http"

	"bitbucket.org/matchmove/integration-svc-aub/app/resource"
	"bitbucket.org/matchmove/integration-svc-aub/modules/container"
	"github.com/jmoiron/sqlx"
	"github.com/pooja1110/pismo_test/app/errs"
	"gopkg.in/matchmove/rest.v2"
)

// PismoErrors API resource
type PismoError struct {
	resource.Resource
}

// PismoError API resource
type PismoErrors struct {
	resource.Resource
}

// NewPismoErrors ...
func NewPismoErrors(db *sqlx.DB, cont *container.Container) (string, string, func() rest.ResourceType) {
	return resource.ErrorResource, resource.ErrorsEndpoint, func() rest.ResourceType {
		return &PismoErrors{resource.New(resource.ErrorResource, true, db, cont)}
	}
}

// NewPismoError ...
func NewPismoError(db *sqlx.DB, cont *container.Container) (string, string, func() rest.ResourceType) {
	return resource.ErrorResource, resource.ErrorEndpoint, func() rest.ResourceType {
		return &PismoError{resource.New(resource.ErrorResource, true, db, cont)}
	}
}

// Get - sends get request to /heartbeat
func (r *PismoErrors) Get() {
	r.Response.Header().Set("Accept", rest.ContentTypeJSON)
	r.Status = http.StatusOK
	r.RawBody = errs.GetErrors()
	return
}

// Get - sends get request to /heartbeat
func (r *PismoError) Get() {
	r.Response.Header().Set("Accept", rest.ContentTypeJSON)
	r.Status = http.StatusOK
	mErr, err := errs.GetErrorByCode(r.Vars[resource.ErrorCode])
	if err != nil {
		r.FormatException(r, err)
		return
	}
	r.RawBody = mErr
	return
}
