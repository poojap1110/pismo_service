package account

// Request ...
type Request struct {
	DocumentNumber string `json:"document_number" validation_key:"document_number" add_validation:"required"`
}
