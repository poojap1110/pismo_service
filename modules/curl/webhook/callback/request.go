package callback

type TransactionRequest struct {
	URL         string             `json:"url"`
	Method      string             `json:"method"`
	ContentType string             `json:"content_type"`
	Payload     TransactionPayload `json:"payload"`
}

type TransactionPayload struct {
	RefHash         string `json:"event_id"`
	Event           string `json:"event"`
	EventTime       string `json:"event_time"`
	PaymentRefID    string `json:"payment_ref_id"`
	ProgramCode     string `json:"program_code"`
	AccountType     string `json:"account_type"`
	UserHashID      string `json:"user_hash_id,omitempty"`
	WalletHashID    string `json:"wallet_hash_id,omitempty"`
	SourceAccount   string `json:"source_account"`
	VirtualAccount  string `json:"virtual_account,omitempty"`
	PP              string `json:"pp"`
	Amount          string `json:"amount"`
	ChargeFee       string `json:"charge_fee"`
	Currency        string `json:"currency"`
	TransactionType string `json:"transaction_type"`
	Message         string `json:"message"`
	Comments        string `json:"comments"`
	Status          string `json:"status"`
}
