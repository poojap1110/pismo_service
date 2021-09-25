package balance_enquiry

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/cimb"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"encoding/json"
	"fmt"
	"os"
)

// Const ...
const (
	method   = "GET"
	endpoint = "accounts/CASA/%s"
	LogName  = cimb.LogName + "-GETBalance"
)

// Call ..
func Call(c *container.Container, account_id string, token string) (raw map[string]interface{}, err error) {

	url := os.Getenv(constant.EnvCimbHOST)
	endpoints := fmt.Sprintf(endpoint, account_id)

	cb := cimb.GetInstance(c)
	cb = (*cb).SetUrl(url)
	cb = (*cb).SetMethod(method)
	cb = (*cb).SetEndpoint(endpoints)
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
