package wallet_funds

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/model"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/op"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/op/oauth_funds"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
)

const (
	Endpoint = "users/wallets/funds"

	LogName = op.LogName + "-walletfunds"
	Retry   = 2
)

//Call proper ...
func Call(c *container.Container, payload map[string]interface{}, product_code string, method string, oauthdb *sqlx.DB, is_retry bool) (response string, res WalletFundsResponse, err error) {
	url := fmt.Sprintf(os.Getenv(constant.EnvOPProxyURL), product_code) + os.Getenv(constant.EnvOPProxyVersion)
	//byt, _ := json.Marshal(payload)
	opIns := op.GetInstance(c)
	opIns = (*opIns).SetUrl(url)
	opIns = (*opIns).SetMethod(method)
	opIns = (*opIns).SetEndpoint(Endpoint)
	opIns = (*opIns).StructToPayload(payload)

	oauth, _ := model.NewOauthConsumer(oauthdb, nil)
	oauth.Products = product_code
	err = oauth.GetConsumerCredentials()
	if err != nil {
		return
	}

	opIns = (*opIns).SetAuthentication(curl.BasicAuth, oauth.Key, oauth.Secret)

	for i := 0; i < Retry; i++ {
		_, body, err2 := (*opIns).Call()

		requestJson, _ := json.Marshal(payload)
		if err2 != nil {
			go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + string(requestJson) + "\nResponse:\n" + err2.Error()))
			err = errors.New(constant.ErrorExternalService + ": OP : " + err2.Error())
			if is_retry {
				continue
			}
			return
		}

		go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + string(requestJson) + "\nResponse:\n" + string(body)))
		response = string(body)
		var raw map[string]interface{}
		json.Unmarshal(body, &raw)
		if _, ok := raw["id"].(string); ok {
			json.Unmarshal(body, &res)
			if (method == http.MethodDelete || method == http.MethodPost) && res.Confirm == "require" {
				params := map[string]interface{}{
					"ids": raw["id"],
				}
				_, err2 := oauth_funds.Call(c, params, product_code, oauth.Key, oauth.Secret)
				if err2 != nil {
					err2 = errors.New(constant.ErrorExternalService + ": OP")
					return
				}
			}
			return
		}

		errResponse := op.ErrorResponse{}
		json.Unmarshal(body, &errResponse)
		err = errors.New(constant.ErrorExternalService + ": OP : " + errResponse.Description)
		break
	}
	return

}
