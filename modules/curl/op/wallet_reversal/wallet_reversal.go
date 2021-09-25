package wallet_reversal

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
	method   = "DELETE"
	Endpoint = "/oauth/consumer/wallet/transaction/%s"
	Retry    = 2
	LogName  = op.LogName + "-walletreversal"
)

//Call proper ...
func Call(c *container.Container, product_code string, ref_id string, oauthdb *sqlx.DB, is_retry bool) (response string, raw Response, err error) {
	url := fmt.Sprintf(os.Getenv(constant.EnvOPProxyURL), product_code) + os.Getenv(constant.EnvOPProxyVersion)
	endpoint := fmt.Sprintf(Endpoint, ref_id)
	opIns := op.GetInstance(c)
	opIns = (*opIns).SetUrl(url)
	opIns = (*opIns).SetMethod(method)
	opIns = (*opIns).SetEndpoint(endpoint)

	oauth, _ := model.NewOauthConsumer(oauthdb, nil)
	oauth.Products = product_code
	err = oauth.GetConsumerCredentials()
	if err != nil {
		return
	}

	opIns = (*opIns).SetAuthentication(curl.BasicAuth, oauth.Key, oauth.Secret)

	for i := 0; i < Retry; i++ {

		_, body, err2 := (*opIns).Call()

		if err2 != nil {
			go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + ref_id + "\nResponse:\n" + err2.Error()))
			err = errors.New(constant.ErrorExternalService + ": OP : " + err2.Error())
			if is_retry {
				continue
			}
			return
		}

		response = string(body)
		go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + ref_id + "\nResponse:\n" + string(body)))

		err = json.Unmarshal(body, &raw)
		break
	}
	return
}
