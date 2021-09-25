package auth_code

import (
	logs "bitbucket.org/matchmove/fmt-logs"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/model"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/cimb_oauth"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/cimb_oauth/access_token"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"encoding/json"
	"net/url"
	"os"
	"strconv"
)

// Const ...
const (
	method       = "GET"
	endpoint     = "authorization"
	LogName      = cimb_oauth.LogName + "-GETAuth_code"
	ResponseType = "code"
	GrantType    = "authorization_code"
)

// Call ..
func Call(c *container.Container, oauth *model.OauthParams) (raw map[string]interface{}, err error) {
	var (
		log         = logs.New()
		queryParams map[string]interface{}
	)
	oauth.RedirectUri = os.Getenv(constant.EnvCimbRedirectUri)
	oauth.ResponseType = ResponseType
	oauth.GrantType = GrantType
	bt, _ := json.Marshal(oauth)
	json.Unmarshal(bt, &queryParams)
	urlstr := os.Getenv(constant.EnvCimbOauthHOST)
	endpoints := urlstr + endpoint
	u, _ := url.Parse(endpoints)
	q, _ := url.ParseQuery(endpoints)

	for key, value := range queryParams {
		switch value.(type) {
		case string:
			q.Add(key, value.(string))
			break
		case float64:
			q.Add(key, strconv.FormatFloat(value.(float64), 'E', 2, 32))
		}
	}

	u.RawQuery = q.Encode()

	cb := cimb_oauth.GetInstance(c)
	cb = (*cb).SetUrl(u.String())
	cb = (*cb).SetMethod(method)
	//cb = (*cb).SetEndpoint("")
	//cb = (*cb).StructToPayload(payloadParams)
	log.Print("CIMB Auth code Request :", u.String())
	_, body, err2 := (*cb).Call()
	if err2 != nil {
		log.Print("CIMB Auth code Error :", err2.Error())
		log.Dump()
		err = err2
		return
	}

	log.Print("CIMB Auth code Response :", string(body))
	log.Dump()
	requestJson, _ := json.Marshal(queryParams)
	go logger.GetInstance(c, LogName).WriteLog([]byte("Request auth code:\n" + "\n Request :\n" + string(requestJson) + "\nResponse:\n" + string(body)))
	var rawdata map[string]interface{}
	err = json.Unmarshal(body, &rawdata)
	if err != nil {
		return
	}
	var ok bool
	if queryParams["code"], ok = rawdata["code"].(string); !ok {
		return
	}
	raw, err = access_token.Call(c, queryParams)
	if err != nil {
		return
	}

	return
}
