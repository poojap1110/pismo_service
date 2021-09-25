package get_txn_response

import "encoding/xml"

// Soap Response
type Response struct {
	XMLName xml.Name
	Body    Body
}

type Body struct {
	XMLName                     xml.Name
	GetTxnResponseInXmlResponse GetTxnResponseInXmlResponse `xml:"GetTxnResponseInXmlResponse"`
}

type GetTxnResponseInXmlResponse struct {
	XMLName                   xml.Name                  `xml:"GetTxnResponseInXmlResponse"`
	XMLNS                     string                    `xml:"xmlns,attr"`
	GetTxnResponseInXmlResult GetTxnResponseInXmlResult `xml:"GetTxnResponseInXmlResult"`
}

type GetTxnResponseInXmlResult struct {
	XMLName        xml.Name       `xml:"GetTxnResponseInXmlResult"`
	PaymentEnqResp PaymentEnqResp `xml:"PaymentEnqResp"`
}

type PaymentEnqResp struct {
	XMLName     xml.Name            `xml:"PaymentEnqResp"`
	Transaction TransactionResponse `xml:"Transaction"`
}

type TransactionResponse struct {
	XMLName       xml.Name `xml:"Transaction"`
	IBLRefNo      string   `xml:"IBLRefNo"`
	CustomerRefNo string   `xml:"CustomerRefNo"`
	TranType      string   `xml:"TranType"`
	Amount        string   `xml:"Amount"`
	StatusCode    string   `xml:"StatusCode"`
	StatusDesc    string   `xml:"StatusDesc"`
	UTR           string   `xml:"UTR"`
	PaymentDate   string   `xml:"PaymentDate"`
	ImpsBeneName  string   `xml:"ImpsBeneName"`
}
