package encryption

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"net/http"
	"time"
)

const (
	LogName         = "Encryption"
	Timeoutduration = 10
)

// GetInstance function ...
func GetInstance(c *container.Container) *curl.ICall {
	curlIns := curl.New(c)
	curlIns = (*curlIns).AddHeader(curl.ContentType, curl.ApplicationJson)
	curlIns = (*curlIns).SetClient(&http.Client{
		Timeout: time.Duration(Timeoutduration) * time.Second,
	})

	return curlIns
}
