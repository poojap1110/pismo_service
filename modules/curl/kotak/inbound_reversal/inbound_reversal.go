package inbound_reversal

import (
	"encoding/json"
	"errors"
	"net/url"
	"os"

	logs "bitbucket.org/matchmove/fmt-logs"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/kotak"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
)

const (
	Method  = "POST"
	LogName = kotak.LogName + "-inboundreversal"
	Retry   = 2
)

//Call proper ...
func Call(c *container.Container, payload InboundRevesalRequest, api_key, token string, is_retry bool, endpoint string) (response string, res InboundRevesalResponse, err error) {

	var (
		log = logs.New()
	)

	urlstr := os.Getenv(constant.EnvKotakURL)
	endpoints := urlstr + endpoint
	u, _ := url.Parse(endpoints)
	q, _ := url.ParseQuery(endpoints)

	q.Add("apikey", api_key)

	u.RawQuery = q.Encode()

	reqbyt, _ := json.Marshal(payload)
	cb := kotak.GetInstance(c)
	cb = (*cb).SetUrl(u.String())
	cb = (*cb).SetMethod(Method)
	//cb = (*cb).SetEndpoint("")
	cb = (*cb).SetAuthentication(curl.BearerAuth, token)
	cb = (*cb).ByteToPayload(reqbyt)

	log.Print("KOTAK Inbound reversal URL: ", urlstr, u.String())
	log.Print("KOTAK Inbound reversal Request: ", string(reqbyt))

	for i := 0; i < Retry; i++ {
		_, body, err2 := (*cb).Call()
		if err2 != nil {
			log.Print("KOTAK Inbound reversal Response Err: ", err2.Error())
			log.Dump()
			go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + string(reqbyt) + "\nResponse:\n" + err2.Error()))
			err = errors.New(constant.ErrorExternalService + ": KOTAK: " + err2.Error())
			if is_retry {
				continue
			}
			return
		}
		log.Print("KOTAK Inbound reversal Response: ", string(body), api_key)
		log.Dump()
		go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + string(reqbyt) + "\nResponse:\n" + string(body)))
		response = string(body)
		var raw map[string]interface{}
		json.Unmarshal(body, &raw)
		if _, ok := raw["CMSGenericInboundResponse"]; ok {
			json.Unmarshal(body, &res)
			return
		}
		break
	}
	return

}
