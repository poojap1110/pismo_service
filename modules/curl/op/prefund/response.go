package prefund

type PrefundResponse struct {
	Amount     string `json:"amount"`
	OldBalance string `json:"old_balance"`
	NewBalance string `json:"new_balance"`
}
