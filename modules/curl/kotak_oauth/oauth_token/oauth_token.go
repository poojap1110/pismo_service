package oauth_token

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	logs "bitbucket.org/matchmove/fmt-logs"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/model"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/kotak_oauth"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"github.com/jmoiron/sqlx"
)

// Const ...
const (
	method   = "POST"
	endpoint = "/auth/oauth/v2/token"
	LogName  = kotak_oauth.LogName + "-POSTAccessToken"
)

// Call ..
func Call(c *container.Container, db *sqlx.DB, programcode, endpoint_name string) (raw map[string]interface{}, token string, err error) {

	var (
		oauthparams = make(map[string]interface{})
		oauth, _    = model.NewOauthParams(db, nil, nil)
		log         = logs.New()
		ok          bool
	)

	oauth.BankProvider.SetValid(strings.ToLower(constant.BankKOTAK))
	oauth.Status.SetValid("active")
	oauth.EndpointName.SetValid(endpoint_name)
	oauth.ProgramCode.SetValid(programcode)
	err = oauth.FindFirst()
	if err != nil {
		log.Print("oauth param err :", err.Error())
		log.Dump()
		return
	}

	oauthparams["grant_type"] = "client_credentials"
	oauthparams["client_id"] = oauth.ClientID.String
	oauthparams["client_secret"] = oauth.ClientSecret.String

	url := os.Getenv(constant.EnvKotakURL)
	kt := kotak_oauth.GetInstance(c)
	kt = (*kt).SetUrl(url)
	kt = (*kt).SetMethod(http.MethodPost)
	kt = (*kt).SetEndpoint(oauth.EndpointUri.String)
	kt = (*kt).StructToPayload(oauthparams)

	requestJson, _ := json.Marshal(oauthparams)
	_, body, err2 := (*kt).Call()
	log.Print(" KOTAK Access Token Request curl : ", string(requestJson))
	log.Print(" KOTAK Access Token Response : ", string(body))
	if err2 != nil {
		go logger.GetInstance(c, LogName).WriteLog([]byte("Request Client:\n" + oauthparams["client_id"].(string) + "\n Request :\n" + string(requestJson) + "\nResponse:\n" + err2.Error()))
		log.Print(" KOATK Access Token Response  err: ", err2.Error())
		log.Dump()
		err = err2
		return
	}

	log.Dump()

	go logger.GetInstance(c, LogName).WriteLog([]byte("Request Client:\n" + oauthparams["client_id"].(string) + "\n Request :\n" + string(requestJson) + "\nResponse:\n" + string(body)))

	err = json.Unmarshal(body, &raw)
	if err != nil {
		log.Print(" KOATK Access Token marshalling err: ", err.Error())
		log.Dump()
		return
	}

	if token, ok = raw["access_token"].(string); !ok {
		log.Print("KOTAK Oauth token err: ", raw["error_description"].(string))
		log.Dump()
		err = errors.New(raw["error_description"].(string))
		return
	}

	return
}
