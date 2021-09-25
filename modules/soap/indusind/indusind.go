package indusind

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
	"time"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/config"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/soap"
)

const (
	LogName          = "Indusind"
	XIBMClientId     = "X-IBM-Client-Id"
	XIBMClientSecret = "X-IBM-Client-Secret"
)

//Indusind SOAP Request
type Request struct {
	XMLName   xml.Name `xml:"soapenv:Envelope"`
	XMLNsSoap string   `xml:"xmlns:soapenv,attr"`
	XMLNsTemp string   `xml:"xmlns:tem,attr"`
	Header    string   `xml:"soapenv:Header"`
	Body      RequestBody
}

type RequestBody struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	Payload interface{}
}

// GetInstance function ...
func GetInstance(c *container.Container) *soap.ICall {
	var (
		url            string
		clientId       string
		clientSecret   string
		indusindConfig interface{}
	)

	indusindConfig = config.GetInstance(c).GetTechnicalConfigs(constant.TechnicalIndusind)

	url = indusindConfig.(map[string]string)["url"]
	clientId = indusindConfig.(map[string]string)["client_id"]
	clientSecret = indusindConfig.(map[string]string)["client_secret"]

	fmt.Println("-------------------------")
	fmt.Println(url)
	fmt.Println(clientId)
	fmt.Println(clientSecret)
	fmt.Println("-------------------------")

	if clientId == "" || clientSecret == "" || url == "" {
		panic("indusind config values required but not declared")
	}

	SoapIns := soap.New(c)
	SoapIns = (*SoapIns).AddHeader(soap.ContentType, soap.TextXML)
	SoapIns = (*SoapIns).AddHeader(XIBMClientId, clientId)
	SoapIns = (*SoapIns).AddHeader(XIBMClientSecret, clientSecret)
	SoapIns = (*SoapIns).SetUrl(url)

	SoapIns = (*SoapIns).SetClient(&http.Client{
		Timeout: time.Duration(soap.RequestTimeoutSeconds) * time.Second,
	})

	return SoapIns
}

func ParseIndusindResponse(body []byte) (response []byte) {
	response = []byte(strings.Replace(string(body), "&lt;", "<", -1))
	return
}
