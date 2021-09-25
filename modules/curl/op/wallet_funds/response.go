package wallet_funds

type WalletFundsResponse struct {
	ID       string                   `json:"id"`
	RefID    string                   `json:"ref_id"`
	UserID   string                   `json:"user_id"`
	WalletId string                   `json:"wallet_id"`
	Confirm  string                   `json:"confirm"`
	Links    []map[string]interface{} `json:"links"`
}
