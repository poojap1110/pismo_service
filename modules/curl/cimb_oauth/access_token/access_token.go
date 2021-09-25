package access_token

import (
	logs "bitbucket.org/matchmove/fmt-logs"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/cimb_oauth"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"encoding/json"
	"net/http"
	"os"
)

// Const ...
const (
	method   = "POST"
	endpoint = "token"
	LogName  = cimb_oauth.LogName + "-POSTAccessToken"
)

// Call ..
func Call(c *container.Container, payloadParams map[string]interface{}) (raw map[string]interface{}, err error) {

	var (
		log = logs.New()
	)

	url := os.Getenv(constant.EnvCimbOauthHOST)
	cb := cimb_oauth.GetInstance(c)
	cb = (*cb).SetUrl(url)
	cb = (*cb).SetMethod(http.MethodPost)
	cb = (*cb).SetEndpoint(endpoint)
	cb = (*cb).StructToPayload(payloadParams)

	requestJson, _ := json.Marshal(payloadParams)

	log.Print("CIMB Auth Access token Request :", string(requestJson))
	_, body, err2 := (*cb).Call()
	if err2 != nil {
		log.Print("CIMB Auth Access token Error :", err2.Error())
		log.Dump()
		err = err2
		return
	}

	log.Print("CIMB Auth Access token Response :", string(body))
	log.Dump()
	go logger.GetInstance(c, LogName).WriteLog([]byte("Request Client:\n" + payloadParams["client_id"].(string) + "\n Request :\n" + string(requestJson) + "\nResponse:\n" + string(body)))

	err = json.Unmarshal(body, &raw)
	return
}
