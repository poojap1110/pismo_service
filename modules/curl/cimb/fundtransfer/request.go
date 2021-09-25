package fundtransfer

type PostFundTransferRequest struct {
	PaymentId              string `json:"PaymentId"`
	PaymentMode            string `json:"PaymentMode"`
	EndToEndIdentification string `json:"EndToEndIdentification"`
	Initiation             `json:"Initiation"`
}

type Initiation struct {
	CreditorAccount       `json:"CreditorAccount"`
	CreditorAgent         `json:"CreditorAgent,omitempty"`
	DebtorAccount         `json:"DebtorAccount"`
	InstructedAmount      `json:"InstructedAmount"`
	RemittanceInformation `json:"RemittanceInformation"`
}

type CreditorAccount struct {
	AcctNo string `json:"AcctNo"`
	Name   string `json:"Name"`
}

type CreditorAgent struct {
	BIC string `json:"BIC"`
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
