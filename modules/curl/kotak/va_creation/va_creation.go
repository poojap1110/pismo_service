package va_creation

import (
	logs "bitbucket.org/matchmove/fmt-logs"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/model"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"encoding/xml"
	"errors"
	"net/url"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/kotak"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"net/http"

	"os"
)

// Const ...
const (
	method   = "POST"
	endpoint = "/VirtualAccountCreation"
	LogName  = kotak.LogName + "-va_creation"
	Retry    = 2
)

// Call ..
func Call(c *container.Container, payloadParams []byte, token string, oauth *model.OauthParams, is_retry bool) (response string, res VirtualAccountMastersResponse, count int64, err error) {

	log := logs.New()

	var (
		body []byte
		err2 error
	)

	urlstr := os.Getenv(constant.EnvKotakURL)
	endpoints := urlstr + endpoint
	u, _ := url.Parse(endpoints)
	q, _ := url.ParseQuery(endpoints)

	q.Add("grant_type", oauth.GrantType)
	q.Add("client_id", oauth.ClientID.String)
	q.Add("client_secret", oauth.ClientSecret.String)

	u.RawQuery = q.Encode()
	cb := kotak.GetInstance(c)
	cb = (*cb).SetUrl(urlstr)
	cb = (*cb).SetMethod(http.MethodPost)
	cb = (*cb).SetEndpoint(endpoint)
	cb = (*cb).SetAuthentication(curl.BearerAuth, token)
	cb = (*cb).ByteToPayload(payloadParams)
	count = 0

	//requestJson, _ := json.Marshal(payloadParams)
	log.Print("KOTAK VaCreation Request :", string(payloadParams))
	log.Print("KOTAK VaCreation url :", u.String())

	for i := 0; i < Retry; i++ {
		count++

		_, body, err2 = (*cb).Call()

		if err2 != nil {
			err = err2
			go logger.GetInstance(c, LogName).WriteLog([]byte("Request ID:\n" + "\n Request :\n" + string(payloadParams) + "\nResponse:\n" + err.Error()))
			if is_retry {
				continue
			}
			return
		}

		response = string(body)

		err = xml.Unmarshal(body, &res)
		if err != nil {
			return
		}

		go logger.GetInstance(c, LogName).WriteLog([]byte("Request ID:\n" + "\n Request :\n" + string(payloadParams) + "\nResponse:\n" + string(body)))

		if res.ResponseParameters.ERRORFLAG != "N" {
			err = errors.New(res.ResponseParameters.RESDESC)
		}

		log.Print("KOTAK VA  Response :", response)
		log.Dump()
		break
	}
	return
}
