package inbound_pull

import (
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"reflect"

	logs "bitbucket.org/matchmove/fmt-logs"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl/kotak"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
)

const (
	Method  = "POST"
	LogName = kotak.LogName + "-inboundfunds"
	Retry   = 2
)

//Call proper ...
func Call(c *container.Container, payload InboundPullRequest, api_key, token string, is_retry bool, endpoint string) (response string, res InboundPullResponse, is_details bool, err error) {

	var (
		log = logs.New()
	)

	is_details = false

	urlstr := os.Getenv(constant.EnvKotakURL)
	endpoints := urlstr + endpoint
	u, _ := url.Parse(endpoints)
	q, _ := url.ParseQuery(endpoints)

	q.Add("apikey", api_key)

	u.RawQuery = q.Encode()

	reqbyt, _ := json.Marshal(payload)
	cb := kotak.GetInstance(c)
	cb = (*cb).SetUrl(u.String())
	cb = (*cb).SetMethod(Method)
	//cb = (*cb).SetEndpoint("")
	cb = (*cb).ByteToPayload(reqbyt)
	cb = (*cb).SetAuthentication(curl.BearerAuth, token)

	log.Print("KOTAK URL: ", u.String())
	log.Print("KOTAK Request: ", string(reqbyt))

	for i := 0; i < Retry; i++ {
		_, body, err2 := (*cb).Call()
		if err2 != nil {
			log.Print("KOTAK Response Err: ", err2.Error())
			log.Dump()
			go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + string(reqbyt) + "\nResponse:\n" + err2.Error()))
			err = errors.New(constant.ErrorExternalService + ": KOTAK : " + err2.Error())
			if is_retry {
				continue
			}
			return
		}
		//body = FormatMockResponse()
		log.Print("KOTAK Response: ", string(body), api_key)
		log.Dump()
		go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + string(reqbyt) + "\nResponse:\n" + string(body)))
		response = string(body)
		var raw map[string]interface{}
		json.Unmarshal(body, &raw)
		if _, ok := raw["CMSGenericInboundResponse"]; ok {
			details := raw["CMSGenericInboundResponse"].(map[string]interface{})["CMSGenericInboundRes"].(map[string]interface{})["CollectionDetails"]
			if reflect.TypeOf(details).String() != "string" {
				if len(details.(map[string]interface{})["CollectionDetail"].([]interface{})) > 0 {
					json.Unmarshal(body, &res)
					is_details = true
				}
			}
			return
		}

		errResponse := InboundPullErrResponse{}
		json.Unmarshal(body, &errResponse)
		err = errors.New(constant.ErrorExternalService + ": KOTAK: " + errResponse.ErrorDesc)
		break
	}
	return

}

func FormatMockResponse() (byt []byte) {
	var str = "{\n    \"CMSGenericInboundResponse\": {\n        \"Header\": {\n            \"Srcappcd\": \"ECOLLECTION\",\n            \"RequestID\": \"1434567890\"\n        },\n        \"CMSGenericInboundRes\": {\n            \"CollectionDetails\": {\n                \"CollectionDetail\": [\n                    {\n                        \"Master_Acc_No\": \"09582650000173\",\n                        \"Remitt_Info\": \"RTGS\",\n                        \"Remit_Name\": \"XYZ12\",\n                        \"Remit_Ifsc\": \"UTIB0001030\",\n                        \"REF3\": \"\",\n                        \"Amount\": \"833.75\",\n                        \"Txn_Ref_No\": \"MATCH20011500012\",\n                        \"Utr_No\": \"MATCH20011500012\",\n                        \"Pay_Mode\": \"RTGS\",\n                        \"E_Coll_Acc_No\": \"BAXA500682566421071965\",\n                        \"Remit_Ac_Nmbr\": \"915010022313757\",\n                        \"Creditdateandtime\": \"2020-07-13 00:00:00\",\n                        \"REF1\": \"\",\n                        \"REF2\": \"\",\n                        \"Txn_Date\": \"2021-10-02 00:00:00\",\n                        \"Bene_Cust_Acname\": \"BHARATI AXA LTD\"\n                    },\n\t\t     {\n                        \"Master_Acc_No\": \"09582650000173\",\n                        \"Remitt_Info\": \"RTGS\",\n                        \"Remit_Name\": \"XYZ13\",\n                        \"Remit_Ifsc\": \"UTIB0001030\",\n                        \"REF3\": \"\",\n                        \"Amount\": \"6135.09\",\n                        \"Txn_Ref_No\": \"MATCH20011500013\",\n                        \"Utr_No\": \"MATCH20011500013\",\n                        \"Pay_Mode\": \"RTGS\",\n                        \"E_Coll_Acc_No\": \"BAXA500682566421071979\",\n                        \"Remit_Ac_Nmbr\": \"915010022313757\",\n                        \"Creditdateandtime\": \"2020-07-13 00:00:00\",\n                        \"REF1\": \"\",\n                        \"REF2\": \"\",\n                        \"Txn_Date\": \"2021-10-02 00:00:00\",\n                        \"Bene_Cust_Acname\": \"BHARATI AXA LTD\"\n                    }\n                ]\n            }\n        }\n    }\n}"

	byt = []byte(str)
	return
}
