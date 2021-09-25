package fundtransfer

type PostFundTransferResponse struct {
	Data `json:"Data"`
}

type Data struct {
	PaymentSetupResponse `json:"PaymentSetupResponse"`
}

type PaymentSetupResponse struct {
	PaymentId           string `json:"PaymentId,omitempty"`
	PaymentSubmissionId string `json:"PaymentSubmissionId,omitempty"`
	PaymentInfId        string `json:"PaymentInfId,omitempty"`
	Status              string `json:"Status,omitempty"`
	RejectCode          string `json:"RejectCode,omitempty"`
	RejectDescription   string `json:"RejectDescription,omitempty"`
}
