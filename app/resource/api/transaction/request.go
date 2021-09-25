package transaction

// Request ...
type Request struct {
	AccountID       string `json:"account_id" validation_key:"account_id" add_validation:"required"`
	OperationTypeID string `json:"operation_type_id" validation_key:"operation_type_id" add_validation:"required"`
	Amount          string `json:"amount" validation_key:"amount" add_validation:"required"`
}
