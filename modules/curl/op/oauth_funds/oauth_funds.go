package oauth_funds

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/op"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

const (
	Endpoint = "/oauth/consumer/funds"
	LogName  = op.LogName + "-consumerfunds"
)

//Call proper ...
func Call(c *container.Container, payload map[string]interface{}, product_code string, key, secret string) (res Transactions, err error) {
	url := fmt.Sprintf(os.Getenv(constant.EnvOPProxyURL), product_code) + os.Getenv(constant.EnvOPProxyVersion)
	//byt, _ := json.Marshal(payload)
	opIns := op.GetInstance(c)
	opIns = (*opIns).SetUrl(url)
	opIns = (*opIns).SetMethod(http.MethodDelete)
	opIns = (*opIns).SetEndpoint(Endpoint)
	opIns = (*opIns).StructToPayload(payload)

	opIns = (*opIns).SetAuthentication(curl.BasicAuth, key, secret)

	_, body, err2 := (*opIns).Call()

	requestJson, _ := json.Marshal(payload)

	if err2 != nil {
		go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + string(requestJson) + "\nResponse:\n" + err2.Error()))
		err = errors.New(constant.ErrorExternalService + ": OP : " + err2.Error())
		return
	}

	go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + string(requestJson) + "\nResponse:\n" + string(body)))

	var raw map[string]interface{}
	json.Unmarshal(body, &raw)
	if _, ok := raw["id"].(string); ok {
		json.Unmarshal(body, &res)
		return
	}

	errResponse := op.ErrorResponse{}
	json.Unmarshal(body, &errResponse)
	err = errors.New(constant.ErrorExternalService + ": OP : " + errResponse.Description)

	return

}
