package webhook

import (
	"net/http"
	"time"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
)

const (
	LogName = "WEBHOOK"
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
