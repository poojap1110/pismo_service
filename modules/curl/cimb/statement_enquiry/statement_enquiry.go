package statement_enquiry

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/cimb"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

// Const ...
const (
	method   = "GET"
	endpoint = "accounts/CASA/%s/statements"
	LogName  = cimb.LogName + "-GETStatement"
)

// Call ..
func Call(c *container.Container, queryParams map[string]interface{}, account_id string, token string) (raw map[string]interface{}, err error) {

	urlstr := os.Getenv(constant.EnvCimbHOST)
	endpoints := fmt.Sprintf(endpoint, account_id)
	endpoints = urlstr + endpoints

	u, _ := url.Parse(endpoints)
	q, _ := url.ParseQuery(endpoints)

	for key, value := range queryParams {
		q.Add(key, value.(string))
	}
	u.RawQuery = q.Encode()

	cb := cimb.GetInstance(c)
	cb = (*cb).SetUrl(u.String())
	cb = (*cb).SetMethod(method)
	//cb = (*cb).SetEndpoint(endpoints)
	cb = (*cb).AddHeader("API-Key", os.Getenv(constant.EnvCimbAPIKey))
	cb = (*cb).SetAuthentication(curl.BearerAuth, token)

	//cb = (*cb).StructToPayload(payloadParams)
	_, body, err2 := (*cb).Call()
	if err2 != nil {
		err = err2
		return
	}

	go logger.GetInstance(c, LogName).WriteLog([]byte("Request Client:\n" + account_id + "\nResponse:\n" + string(body)))

	err = json.Unmarshal(body, &raw)
	return
}
