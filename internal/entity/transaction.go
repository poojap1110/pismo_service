package entity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

// GenerateID generates a unique ID that can be used as an identifier for an entity.
func GenerateID() string {
	return uuid.New().String()
}

// TransactionRequest ...
type TransactionRequest struct {
	AccountID       string `json:"account_id" validation_key:"account_id" add_validation:"required"`
	OperationTypeID string `json:"operation_type_id" validation_key:"operation_type_id" add_validation:"required"`
	Amount          string `json:"amount" validation_key:"amount" add_validation:"required"`
}

// TransactionResponse ...
type TransactionResponse struct {
	TransactionRefNo  string `json:"transaction_ref_no"`
	EventDate         string `json:"event_date"`
	TransactionAmount string `json:"transaction_amount"`
	Balance           string `json:"balance"`
}

// Validate validates the CreateAccountRequest fields.
func (m TransactionRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.AccountID, validation.Required),
		validation.Field(&m.OperationTypeID, validation.Required),
		validation.Field(&m.Amount, validation.Required),
	)
}
