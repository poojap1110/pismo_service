package resource

import (
	"fmt"
)

// Constant related to api routes
const (
	// Static URI values
	PrefixURI      = "/api"
	Version        = "/v1"
	ProductCodeURI = "/{%s}"

	// Dymaic URI params
	ErrorCode = "error_code"

	AddressType = "type"

	AccountID = "account_id"
)

// Variable related to api routes
var (
	// Heartbeat Resoruce name and endpoint
	HeartbeatResource = "heartbeat"
	HeartbeatEndpoint = fmt.Sprintf(PrefixURI + "/heartbeat")

	// Error REsource name and endpoint
	ErrorResource  = "pismo-error"
	ErrorsEndpoint = fmt.Sprintf(Version + "/errors")
	ErrorEndpoint  = fmt.Sprintf(Version+"/error/{%s}", ErrorCode)

	// Account Resource name and endpoint
	AccountsResource        = "accounts"
	AccountResourceEndpoint = fmt.Sprintf(Version + "/accounts")
	GetAccountResource      = "get_account"
	GetAccountEndpoint      = fmt.Sprintf(Version+"/accounts/{%s}", AccountID)

	// Transaction Resource name and endpoint
	TransactionResource         = "transactions"
	TransactionResourceEndpoint = fmt.Sprintf(Version + "/transactions")
)
