package prefund

type PrefundRequest struct {
	Amount           string `json:"amount"`
	TransactionRefID string `json:"transaction_ref_id"`
	Comments         string `json:"comments"`
	AccountRefID     string `json:"account_ref_id"`
	SourceRefID      string `json:"source_ref_id"`
	Narration        string `json:"narration"`
}
