package post_encrypt

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/encryption"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"encoding/json"
	"errors"
	"os"
)

const (
	Endpoint = "/encrypt"
	Method   = "POST"
	LogName  = encryption.LogName + "-encrypt"
	Retry    = 2
)

//Call proper ...
func Call(c *container.Container, payload PostEncryptRequest, is_retry bool) (response []byte, err error) {
	url := os.Getenv(constant.EnvEncryptURL)
	byt, _ := json.Marshal(payload)
	opIns := encryption.GetInstance(c)
	opIns = (*opIns).SetUrl(url)
	opIns = (*opIns).SetMethod(Method)
	opIns = (*opIns).SetEndpoint(Endpoint)
	opIns = (*opIns).ByteToPayload(byt)

	for i := 0; i < Retry; i++ {
		_, body, err2 := (*opIns).Call()

		requestJson, _ := json.Marshal(payload)
		if err2 != nil {
			go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + string(requestJson) + "\nResponse:\n" + err2.Error()))
			err = errors.New(constant.ErrorExternalService + ": Encrypt : " + err2.Error())
			if is_retry {
				continue
			}
			return
		}

		go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + string(requestJson) + "\nResponse:\n" + string(body)))
		response = body
		break
	}
	return

}
