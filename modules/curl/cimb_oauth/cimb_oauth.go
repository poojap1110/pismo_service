package cimb_oauth

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"net/http"
	"time"
)

const (
	LogName = "CIMB_Auth"
)

// GetInstance function ...
func GetInstance(c *container.Container) *curl.ICall {
	curlIns := curl.New(c)
	curlIns = (*curlIns).AddHeader(curl.ContentType, curl.ApplicationFormUrlEncoded)
	curlIns = (*curlIns).SetClient(&http.Client{
		Timeout: time.Duration(curl.RequestTimeoutSeconds) * time.Second,
	})

	return curlIns
}
