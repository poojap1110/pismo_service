package fundinrequest

type PostFundTransferInRequest struct {
	PaymentId              string `json:"PaymentId"`
	PaymentMode            string `json:"PaymentMode"`
	EndToEndIdentification string `json:"EndToEndIdentification"`
	MandateId              string `json:"MandateId"`
	Initiation             `json:"Initiation"`
}

type Initiation struct {
	CreditorAccount       `json:"CreditorAccount"`
	DebtorAgent           `json:"DebtorAgent,omitempty"`
	DebtorAccount         `json:"DebtorAccount"`
	InstructedAmount      `json:"InstructedAmount"`
	RemittanceInformation `json:"RemittanceInformation"`
}

type CreditorAccount struct {
	AcctNo string `json:"AcctNo"`
	Name   string `json:"Name"`
}

type DebtorAgent struct {
	BIC string `json:"BIC"`
}

type DebtorAccount struct {
	AcctNo string `json:"AcctNo"`
	Name   string `json:"Name"`
}
type InstructedAmount struct {
	Amount float64 `json:"Amount"`
}

type RemittanceInformation struct {
	CustSegmentCode  string `json:"CustSegmentCode"`
	PurposeOfPayment string `json:"PurposeOfPayment"`
	Reference        string `json:"Reference"`
}
