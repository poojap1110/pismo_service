package process_transaction

import "encoding/xml"

type ProcessTxnInXml struct {
	XMLName  xml.Name `xml:"tem:ProcessTxnInXml"`
	CustId   string   `xml:"tem:strCustId"`
	InputTxn InputTxn `xml:"tem:strInputTxn"`
}

type InputTxn struct {
	XMLName xml.Name `xml:"tem:strInputTxn"`
	Value   string   `xml:",cdata"`
}

type PaymentRequest struct {
	XMLName     xml.Name           `xml:"PaymentRequest"`
	Transaction TransactionRequest `xml:"Transaction"`
}

type TransactionRequest struct {
	XMLName          xml.Name `xml:"Transaction"`
	CustomerRefNum   string   `xml:"CustomerRefNum"`
	TransactionType  string   `xml:"TranType"`
	DebitAccount     string   `xml:"DebitAccount"`
	Amount           string   `xml:"Amount"`
	ValueDate        string   `xml:"ValueDate"`
	BenName          string   `xml:"BenName"`
	BenAccountNumber string   `xml:"BENE_ACNO"`
	BenIFSCCode      string   `xml:"BENE_IFSC_CODE"`
	BenBranch        string   `xml:"BENE_BRANCH"`
	BenBank          string   `xml:"BENE_BANK"`
	BenMobileNumber  string   `xml:"Bene_MobileNo"`
	BenEmailId       string   `xml:"Bene_EmailId"`
	BenMMId          string   `xml:"Bene_MMId"`
	MakerId          string   `xml:"MakerId"`
	CheckerId        string   `xml:"CheckerId"`
	Reserve1         string   `xml:"Reserve1"`
	Reserve2         string   `xml:"Reserve2"`
	Reserve3         string   `xml:"Reserve3"`
}
