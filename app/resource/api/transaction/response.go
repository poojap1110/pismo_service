package transaction

// Response ...
type Response struct {
	TransactionRefNo  string `json:"transaction_ref_no"`
	EventDate         string `json:"event_date"`
	TransactionAmount string `json:"transaction_amount"`
	Balance           string `json:"balance"`
}
