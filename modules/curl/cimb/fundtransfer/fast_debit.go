package fundtransfer

import (
	logs "bitbucket.org/matchmove/logs"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/cimb"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/helper"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
)

// Const ...
const (
	method   = "POST"
	endpoint = "payments/sgps"
	LogName  = cimb.LogName + "-fasttransfer"
	Retry    = 2
)

// Call ..
func Call(c *container.Container, payloadParams map[string]interface{}, api_key, client_id, signature, token string, is_retry bool) (response string, count int64, raw PostFundTransferResponse, err error) {

	log := logs.New()

	var (
		body   []byte
		fndreq PostFundTransferRequest
		err2   error
	)

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

	requestJson, _ := json.Marshal(payloadParams)
	log.Print("CIMB FastCredit Request :", string(requestJson))
	json.Unmarshal(requestJson, &fndreq)

	for i := 0; i < Retry; i++ {
		count++
		if fndreq.Reference == "MOCK" {
			body = mock(fndreq)
		} else {
			_, body, err2 = (*cb).Call()
		}
		if err2 != nil {
			err = err2
			go logger.GetInstance(c, LogName).WriteLog([]byte("Request ID:\n" + payloadParams["PaymentId"].(string) + "\n Request :\n" + string(requestJson) + "\nResponse:\n" + err.Error()))
			if is_retry {
				continue
			}
			return
		}
		var rawres map[string]interface{}
		json.Unmarshal(body, &rawres)
		byt, _ := json.Marshal(rawres)
		response = string(byt)
		go logger.GetInstance(c, LogName).WriteLog([]byte("Request ID:\n" + payloadParams["PaymentId"].(string) + "\n Request :\n" + string(requestJson) + "\nResponse:\n" + string(body)))
		json.Unmarshal(body, &raw)
		if raw.Status != "ACTC" {
			err = errors.New(raw.RejectDescription)
			log.Print("CIMB Reject :", raw.RejectDescription)
			log.Dump()
			return
		}
		log.Print("CIMB Response :", response)
		log.Dump()
		break
	}
	return
}

func mock(payloadParams PostFundTransferRequest) (body []byte) {
	var res PostFundTransferResponse

	res.Status = "ACTC"
	res.PaymentId = payloadParams.PaymentId
	res.PaymentInfId = "API" + payloadParams.PaymentId
	res.PaymentSubmissionId = helper.GetRandomID()

	body, _ = json.Marshal(res)
	return body
}
