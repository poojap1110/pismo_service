package get_txn_response

import (
	"encoding/xml"
	"log"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/soap"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/soap/indusind"
)

const (
	method     = "POST"
	soapAction = "http://tempuri.org/IDomesticPayService/GetTxnResponseInXml"

	//LogName indusind get transaction response request and response log
	LogName = indusind.LogName + "-GetTxnResponseInXml"

	// retryServiceCall
	retryServiceCall = 3
)

// GetTxnResponse
func GetTxnResponse(c *container.Container, customerId, IBLRefNo string) (paymentEnqResp PaymentEnqResp, err error) {
	var (
		body              []byte
		reqPayload        indusind.Request
		response          Response
		parsedBody        []byte
		reqPaymentEnquiry PaymentEnquiry
	)

	reqPaymentEnquiry.CustomerId = customerId
	reqPaymentEnquiry.IBLRefNo = IBLRefNo

	//Prepare CDATA for payment request xml
	reqPaymentEnquiryXML, err := xml.MarshalIndent(reqPaymentEnquiry, "", "  ")
	if err != nil {
		log.Println("Failed to do xml.MarshalIndent : GetTxnResponse PaymentEnquiry :", err)
		return
	}

	//Prepare xml request payload
	reqBody := GetTxnResponseInXml{
		InputTxn: InputTxn{
			Value: string(reqPaymentEnquiryXML),
		},
	}

	reqPayload = indusind.Request{
		XMLNsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		XMLNsTemp: "http://tempuri.org/",
		Body: indusind.RequestBody{
			Payload: reqBody,
		},
	}

	soapIns := indusind.GetInstance(c)

	soapIns = (*soapIns).SetMethod(method)
	soapIns = (*soapIns).SetPayload(reqPayload)
	soapIns = (*soapIns).AddHeader(soap.SoapAction, soapAction)

	for i := 0; i < retryServiceCall; i++ {
		// Call SOAP service
		_, body, err = (*soapIns).Call()

		if err != nil {
			log.Println("Failed to call SOAP : GetTxnResponse : ", err)
		}

		xmlReq, err2 := xml.Marshal(reqPayload)
		if err2 != nil {
			log.Println("Failed to do Marshal : GetTxnResponse Request :", err2)
			err = err2
		}

		go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + string(xmlReq) + "\nResponse:\n" + string(body)))

		if err == nil {
			break
		}
	}

	//Parse response body here convert &lt; to < for xml tag
	parsedBody = indusind.ParseIndusindResponse(body)

	//Convert SOAP response xml to go response
	err = xml.Unmarshal(parsedBody, &response)
	if err != nil {
		log.Println("Failed to do unmarshal : GetTxnResponse XML Response : ", err)
		return
	}

	paymentEnqResp = response.Body.GetTxnResponseInXmlResponse.GetTxnResponseInXmlResult.PaymentEnqResp

	return
}
