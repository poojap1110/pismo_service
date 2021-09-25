package oauth

import (
	logs "bitbucket.org/matchmove/fmt-logs"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/model"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/kotak_oauth"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"encoding/json"
	"net/url"
	"os"
)

// Const ...
const (
	method       = "POST"
	endpoint     = "/auth/oauth/v2/token"
	LogName      = kotak_oauth.LogName + "-GETOauth_code"
	ResponseType = "code"
	GrantType    = "client_credentials"
)

// Call ..
func Call(c *container.Container, oauth *model.OauthParams) (rawdata map[string]interface{}, err error) {
	var (
		log         = logs.New()
		queryParams map[string]interface{}
	)
	oauth.GrantType = GrantType
	bt, _ := json.Marshal(oauth)
	json.Unmarshal(bt, &queryParams)
	urlstr := os.Getenv(constant.EnvKotakURL)
	endpoints := urlstr + endpoint
	u, _ := url.Parse(endpoints)
	q, _ := url.ParseQuery(endpoints)

	q.Add("grant_type", oauth.GrantType)
	q.Add("client_id", oauth.ClientID.String)
	q.Add("client_secret", oauth.ClientSecret.String)

	u.RawQuery = q.Encode()
	requestJson, _ := json.Marshal(queryParams)
	cb := kotak_oauth.GetInstance(c)
	cb = (*cb).SetUrl(u.String())
	cb = (*cb).SetMethod(method)

	log.Print("KOTAK Oauth Request :", string(requestJson))

	_, body, err2 := (*cb).Call()
	if err2 != nil {
		log.Print("KOTAK Oauth Err  Response :", err2.Error())
		log.Dump()
		err = err2
		return
	}

	go logger.GetInstance(c, LogName).WriteLog([]byte("Request auth code:\n" + "\n Request :\n" + string(requestJson) + "\nResponse:\n" + string(body)))

	log.Print("KOTAK Oauth  Response :", string(body))
	log.Dump()

	err = json.Unmarshal(body, &rawdata)

	return
}
