package wallet_funds

type WalletFundsRequest struct {
	Email    string `json:"email"`
	Details  string `json:"details"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type ReqDetails struct {
	PP                string `json:"pp,omitempty"`
	PaymentRef        string `json:"payment_ref,omitempty"`
	SourceAccount     string `json:"source_account,omitempty"`
	Comments          string `json:"comments,omitempty"`
	Description       string `json:"description,omitempty"`
	PurposeOfTransfer string `json:"purpose_of_transfer"`
	ChargeFee         string `json:"charge_fee,omitempty"`
}
