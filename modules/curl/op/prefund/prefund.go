package prefund

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
	Endpoint = "oauth/consumer/%s/prefund/credit"

	LogName = op.LogName + "-Prefund"
	Retry   = 2
)

//Call proper ...
func Call(c *container.Container, payload map[string]interface{}, product_code string, method string, oauthdb *sqlx.DB, is_retry bool) (response string, res PrefundResponse, err error) {
	url := fmt.Sprintf(os.Getenv(constant.EnvOPProxyURL), product_code) + os.Getenv(constant.EnvOPProxyVersion)

	opIns := op.GetInstance(c)
	opIns = (*opIns).SetUrl(url)
	opIns = (*opIns).SetMethod(method)

	opIns = (*opIns).StructToPayload(payload)

	oauth, _ := model.NewOauthConsumer(oauthdb, nil)
	oauth.Products = product_code
	err = oauth.GetConsumerCredentials()
	if err != nil {
		return
	}

	opIns = (*opIns).SetAuthentication(curl.BasicAuth, oauth.Key, oauth.Secret)
	endpoint := fmt.Sprintf(Endpoint, oauth.Key)
	opIns = (*opIns).SetEndpoint(endpoint)

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

		response = string(body)
		go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + string(requestJson) + "\nResponse:\n" + string(body)))
		var raw map[string]interface{}
		json.Unmarshal(body, &raw)
		if _, ok := raw["amount"].(string); ok {
			json.Unmarshal(body, &res)
			return
		}

		errResponse := op.ErrorResponse{}
		json.Unmarshal(body, &errResponse)
		err = errors.New(constant.ErrorExternalService + ": OP : " + errResponse.Description)
		break
	}
	return
}
