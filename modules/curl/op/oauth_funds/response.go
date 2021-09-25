package oauth_funds

type Transaction struct {
	ID             string `json:"id"`
	RefID          string `json:"ref_id"`
	Status         string `json:"status"`
	Description    string `json:"description"`
	SorTransaction string `json:"sor_transaction"`
}

type Transactions []Transaction
