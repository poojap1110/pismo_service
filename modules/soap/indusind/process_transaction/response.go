package process_transaction

import "encoding/xml"

// Soap Response
type Response struct {
	XMLName xml.Name
	Body    Body
}

type Body struct {
	XMLName                 xml.Name
	ProcessTxnInXmlResponse ProcessTxnInXmlResponse `xml:"ProcessTxnInXmlResponse"`
}

type ProcessTxnInXmlResponse struct {
	XMLName               xml.Name              `xml:"ProcessTxnInXmlResponse"`
	XMLNS                 string                `xml:"xmlns,attr"`
	ProcessTxnInXmlResult ProcessTxnInXmlResult `xml:"ProcessTxnInXmlResult"`
}

type ProcessTxnInXmlResult struct {
	XMLName         xml.Name        `xml:"ProcessTxnInXmlResult"`
	PaymentResponse PaymentResponse `xml:"PaymentResponse"`
}

type PaymentResponse struct {
	XMLName     xml.Name            `xml:"PaymentResponse"`
	Transaction TransactionResponse `xml:"Transaction"`
}

type TransactionResponse struct {
	XMLName         xml.Name `xml:"Transaction"`
	IBLRefNo        string   `xml:"IBLRefNo"`
	CustomerRefNum  string   `xml:"CustomerRefNo"`
	TransactionType string   `xml:"TranType"`
	Amount          string   `xml:"Amount"`
	StatusCode      string   `xml:"StatusCode"`
	StatusDesc      string   `xml:"StatusDesc"`
}
