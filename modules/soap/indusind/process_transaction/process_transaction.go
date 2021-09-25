package process_transaction

import (
	"encoding/xml"
	"errors"
	"log"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/soap"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/soap/indusind"
)

const (
	method     = "POST"
	soapAction = "http://tempuri.org/IDomesticPayService/ProcessTxnInXml"

	//LogName indusind process transaction request and response log
	LogName = indusind.LogName + "-ProcessTxnInXml"

	// SuccessCode Indusind success status code
	SuccessCode = "R000"

	// InternalServerErrorCode internal server error code
	InternalServerErrorCode = "R005"

	// retryServiceCall
	retryServiceCall = 3
)

func ProcessTransaction(c *container.Container, custId string, paymentReq PaymentRequest) (paymentResponse PaymentResponse, err error) {

	var (
		body       []byte
		reqPayload indusind.Request
		response   Response
		parsedBody []byte
	)

	//Prepare CDATA for payment request xml
	paymentRequestXML, err := xml.MarshalIndent(paymentReq, "", "  ")
	if err != nil {
		log.Println("Failed to do xml.MarshalIndent : ProcessTransaction PaymentRequest :", err)
		return
	}

	//Prepare xml request payload
	reqBody := ProcessTxnInXml{
		CustId: custId,
		InputTxn: InputTxn{
			Value: string(paymentRequestXML),
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
			log.Println("Failed to call SOAP : ProcessTransaction : ", err)
		}

		xmlReq, err2 := xml.Marshal(reqPayload)
		if err2 != nil {
			log.Println("Failed to do Marshal : ProcessTransaction Request :", err2)
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
		log.Println("Failed to do unmarshal : ProcessTransaction XML Response : ", err)
		return
	}

	paymentResponse = response.Body.ProcessTxnInXmlResponse.ProcessTxnInXmlResult.PaymentResponse

	if paymentResponse.Transaction.StatusCode == SuccessCode {
		return
	} else if paymentResponse.Transaction.StatusCode == InternalServerErrorCode {
		err = errors.New("Indusind Error :" + paymentResponse.Transaction.StatusDesc)
		return
	}

	return
}
