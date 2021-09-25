package get_user

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/model"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/op"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

const (
	method   = "GET"
	Endpoint = "users"

	LogName = op.LogName + "-GetUser"
)

//Call proper ...
func Call(c *container.Container, product_code string, user_hash_id string, oauthdb *sqlx.DB) (raw map[string]interface{}, err error) {
	url := fmt.Sprintf(os.Getenv(constant.EnvOPProxyURL), product_code) + os.Getenv(constant.EnvOPProxyVersion)

	opIns := op.GetInstance(c)
	opIns = (*opIns).SetUrl(url)
	opIns = (*opIns).SetMethod(method)
	opIns = (*opIns).SetEndpoint(Endpoint)
	opIns = (*opIns).AddHeader("X-Auth-User-Id", user_hash_id)

	oauth, _ := model.NewOauthConsumer(oauthdb, nil)
	oauth.Products = product_code
	err = oauth.GetConsumerCredentials()
	if err != nil {
		return
	}

	opIns = (*opIns).SetAuthentication(curl.BasicAuth, oauth.Key, oauth.Secret)

	_, body, err2 := (*opIns).Call()
	if err2 != nil {
		go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + user_hash_id + "\nResponse:\n" + err2.Error()))
		err = errors.New(constant.ErrorExternalService + ": OP : " + err2.Error())
		return
	}

	go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + user_hash_id + "\nResponse:\n" + string(body)))

	json.Unmarshal(body, &raw)
	if _, ok := raw["email"].(string); !ok {
		errResponse := op.ErrorResponse{}
		json.Unmarshal(body, &errResponse)
		err = errors.New(constant.ErrorExternalService + ": OP : " + errResponse.Description)
		return
	}
	return
}
