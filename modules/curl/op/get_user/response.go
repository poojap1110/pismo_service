package get_user

// ErrorResponse ...
type ErrorResponse struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Link        string `json:"link"`
}
