package entity

import validation "github.com/go-ozzo/ozzo-validation/v4"

// AccountRequest ...
type AccountRequest struct {
	ID             string `json:"id" validation_key:"id" add_validation:"required"`
	DocumentNumber string `json:"document_number" validation_key:"document_number" add_validation:"required"`
}

// AccountResponse ...
type AccountResponse struct {
	Status string `json:"status"`
}

// Validate validates the CreateAccountRequest fields.
func (m AccountRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.DocumentNumber, validation.Required, validation.Length(0, 128)),
	)
}
