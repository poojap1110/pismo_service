package bank_enumeration

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
	endpoint = "/IbftWsExternal/receiving-banks"
	LogName  = aub.LogName + "-GETBankEnumeration"
)

// Call ..
func Call(c *container.Container, payloadParams map[string]interface{}) (raw map[string]interface{}, err error) {

	url := os.Getenv(constant.EnvAUBHOST)
	endpoints := fmt.Sprintf(endpoint)

	byt, _ := json.Marshal(payloadParams)
	cb := cimb.GetInstance(c)
	cb = (*cb).SetUrl(url)
	cb = (*cb).SetMethod(method)
	cb = (*cb).SetEndpoint(endpoints)
	cb = (*cb).AddCustomHeader("AUB_NONCE:", api_key)
	cb = (*cb).ByteToPayload(byt)

	requestJson, _ := json.Marshal(payloadParams)
	log.Print("AUB GET Bank Enumeration Request :", string(requestJson))

	//cb = (*cb).StructToPayload(payloadParams)
	_, body, err2 := (*cb).Call()

	if err2 != nil {
		err = err2
		return
	}

	go logger.GetInstance(c, LogName).WriteLog([]byte("Request Client:\n"  + payloadParams["userName"] + "\nResponse:\n" + string(body)))

	err = json.Unmarshal(body, &raw)
	return
}
