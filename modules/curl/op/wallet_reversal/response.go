package wallet_reversal

type Response struct {
	UserID      string `json:"user_id"`
	WalletID    string `json:"wallet_id"`
	Transaction `json:"transaction"`
}

type Transaction struct {
	Reversal `json:"reversal"`
}

type Reversal struct {
	ID    string `json:"id"`
	RefID string `json:"ref_id"`
}
