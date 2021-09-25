package callback

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/webhook"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"encoding/json"
	"errors"
)

const (
	method = "GET"

	LogName = webhook.LogName + "-EVENT"
)

//Call proper ...
func Call(c *container.Container, request map[string]interface{}) (raw map[string]interface{}, err error) {
	//url := fmt.Sprintf(os.Getenv(constant.EnvOPProxyURL), product_code) + os.Getenv(constant.EnvOPProxyVersion)
	param := make(map[string]interface{})
	whIns := webhook.GetInstance(c)
	whIns = (*whIns).SetUrl(request["url"].(string))
	whIns = (*whIns).SetMethod(request["method"].(string))
	whIns = (*whIns).AddHeader(curl.ContentType, request["content_type"].(string))

	regbyt, _ := json.Marshal(request["payload"])

	if request["content_type"] == curl.ApplicationJson {
		whIns = (*whIns).SetPayload(string(regbyt))
	} else {
		json.Unmarshal(regbyt, &param)
		whIns = (*whIns).StructToPayload(param)
	}
	_, body, err2 := (*whIns).Call()
	requestJson, _ := json.Marshal(request)

	if err2 != nil {
		go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + string(requestJson) + "\nResponse:\n" + err2.Error()))
		err = errors.New(constant.ErrorExternalService + ": Webhook : " + err2.Error())
		return
	}

	go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + string(requestJson) + "\nResponse:\n" + string(body)))

	err = json.Unmarshal(body, &raw)

	return

}
