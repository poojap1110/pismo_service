package inhouse_transfer

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/cimb"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

// Const ...
const (
	method   = "POST"
	endpoint = "transfers/iht"
	LogName  = cimb.LogName + "-POSTInhouseTransfer"
	Retry    = 2
)

// Call ..
func Call(c *container.Container, payloadParams map[string]interface{}, api_key, client_id, signature, token string, is_retry bool) (response string, count int64, raw PostInhouseTransferResponse, err error) {
	byt, _ := json.Marshal(payloadParams)
	urlstr := os.Getenv(constant.EnvCimbHOST)
	endpoints := urlstr + endpoint
	u, _ := url.Parse(endpoints)
	q, _ := url.ParseQuery(endpoints)
	q.Add("client_id", client_id)
	u.RawQuery = q.Encode()
	cb := cimb.GetInstance(c)
	cb = (*cb).SetUrl(u.String())
	cb = (*cb).SetMethod(http.MethodPost)
	cb = (*cb).AddCustomHeader("API-Key", api_key)
	cb = (*cb).AddHeader("Signature", signature)
	cb = (*cb).SetAuthentication(curl.BearerAuth, token)
	cb = (*cb).ByteToPayload(byt)

	count = 0
	for i := 0; i < Retry; i++ {
		count++
		_, body, err2 := (*cb).Call()
		requestJson, _ := json.Marshal(payloadParams)
		if err2 != nil {
			err = err2
			go logger.GetInstance(c, LogName).WriteLog([]byte("Request ID:\n" + payloadParams["ReqId"].(string) + "\n Request :\n" + string(requestJson) + "\nResponse:\n" + err.Error()))
			if is_retry {
				continue
			}
			return
		}

		response = string(body)
		go logger.GetInstance(c, LogName).WriteLog([]byte("Request Client:\n" + payloadParams["ReqId"].(string) + "\n Request :\n" + string(requestJson) + "\nResponse:\n" + string(body)))

		err = json.Unmarshal(body, &raw)
		break
	}
	return
}
