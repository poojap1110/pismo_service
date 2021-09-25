# model
--
    import "bitbucket.org/matchmove/enumeration/model"


## Usage

```go
const (
	ErrorBranchIDRequired                    = "V078"
	ErrorBankIDRequired                      = "V052"
	ErrorRemitIDRequired                     = "V051"
	ErrorAgentIDRequired                     = "V014"
	ErrorRequestEmpty                        = "V011"
	ErrorRequestProviderMismatch             = "V079"
	ErrorRemittanceConfirmStatusFailed       = "V068"
	ErrorServiceUnavailable                  = "R009"
	ErrorRemittanceAmountIsGreaterThanZero   = "V056"
	ErrorRemittanceAccountNumberEmpty        = "V057"
	ErrorRemittanceBranchCodeEmpty           = "V080"
	ErrorRemittanceBankNameEmpty             = "V081"
	ErrorRemittanceIFSCEmpty                 = "V065"
	ErrorRemittanceSWIFTEmpty                = "V064"
	ErrorRemittanceReceiverMobileNumberEmpty = "V082"
	ErrorRemittanceRoutingParamEmpty         = "V058"
)
```
error code

```go
const (
	ErrorChannelNotFound             = "R401"
	ErrorAgentNotFound               = "R402"
	ErrorRemittanceNotFound          = "R403"
	ErrorBankNotFound                = "R404"
	ErrorBranchNotFound              = "R405"
	ErrorProviderEnvironmentNotFound = "R406"
	ErrorProviderFound               = "R407"
)
```
error code

```go
const (
	// ResourceNotSupported ...
	ResourceNotSupported = "Resource is not supported"

	// TransactionRejectedErrorCode ...
	TransactionRejectedErrorCode = "REJECTED"
)
```

```go
const (
	MessageModeInvalid          = "invalid"
	MessageModeEmpty            = "empty"
	MessageModeEmptyURI         = "uri_empty"
	MessageModeEmptyHeader      = "header_empty"
	MessageModeRange            = "range"
	MessageModeLength           = "length"
	MessageModeGreaterThanEqual = "greater_than_equal"
	MessageModeGreaterThan      = "greater_than"
	MessageModeNotFound         = "not_found"
	MessageModeNotSet           = "not_set"
)
```
Message Modes

```go
const (
	// DBTableChannel specifies the name of the `channel` table in the database
	DBTableChannel = "channel"
)
```

```go
const (
	// DBTableCountry specifies the name of the `country` table in the database
	DBTableCountry = "country"
)
```

```go
const (
	// DBTableEnumeration specifies the name of the `person` table in the database
	DBTableEnumeration = "enumeration"
)
```

```go
const (
	// DBTableProvider specifies the name of the `provider` table in the database
	DBTableProvider = "provider"
)
```

```go
const (
	// DBTableSample specifies the name of the `agent` table in the database
	DBTableSample = "sample"
)
```

```go
var AccessDeniedErrors = map[string]Status{
	secure.MD5("Access Denied"): Status{
		Code: http.StatusUnauthorized,
		Error: Error{
			ID:          "S401",
			Description: "Access Denied",
		},
	},
	secure.MD5(ValidationErrorMessage("provider", MessageModeEmptyHeader)): Status{
		Code: http.StatusUnauthorized,
		Error: Error{
			ID:          "S402",
			Description: ValidationErrorMessage("provider", MessageModeEmptyHeader),
		},
	},
	secure.MD5(ValidationErrorMessage("consumer_key", MessageModeEmptyHeader)): Status{
		Code: http.StatusUnauthorized,
		Error: Error{
			ID:          "S403",
			Description: ValidationErrorMessage("consumer_key", MessageModeEmptyHeader),
		},
	},
	secure.MD5(ValidationErrorMessage("product", MessageModeEmptyHeader)): Status{
		Code: http.StatusUnauthorized,
		Error: Error{
			ID:          "S404",
			Description: ValidationErrorMessage("product", MessageModeEmptyHeader),
		},
	},
}
```
AccessDeniedErrors ....

```go
var BadRequestErrors = map[string]Status{
	secure.MD5("limit must be an unsigned integer"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V001",
			Description: ValidationErrorMessage("limit", MessageModeInvalid),
		},
	},
	secure.MD5("offset must be an unsigned integer"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V002",
			Description: ValidationErrorMessage("offset", MessageModeInvalid),
		},
	},
	secure.MD5("limit must be a value 1 to 50"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V003",
			Description: ValidationErrorMessage("limit", MessageModeRange, "1", "50"),
		},
	},
	secure.MD5("offset must be >= 0"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V005",
			Description: ValidationErrorMessage("offset", MessageModeGreaterThanEqual, "0"),
		},
	},
	secure.MD5("Country is required"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V006",
			Description: ValidationErrorMessage("country", MessageModeEmpty),
		},
	},
	secure.MD5("invalid date_range filter parameter"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V007",
			Description: ValidationErrorMessage("date_range", MessageModeInvalid),
		},
	},
	secure.MD5("Cannot accept `order` request"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V008",
			Description: ValidationErrorMessage("order", MessageModeInvalid),
		},
	},
	secure.MD5("ProviderCode: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V009",
			Description: ValidationErrorMessage("provider_code", MessageModeEmpty),
		},
	},
	secure.MD5("ProviderCode: invalid pattern=provider_code"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V010",
			Description: ValidationErrorMessage("provider_code", MessageModeInvalid),
		},
	},
	secure.MD5("Request body is empty"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V011",
			Description: "Request body is empty",
		},
	},
	secure.MD5("Request type invalid"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V012",
			Description: "Request type invalid",
		},
	},
	secure.MD5("ProviderCode: invalid pattern=provider_code"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V013",
			Description: ValidationErrorMessage("provider_code", MessageModeInvalid),
		},
	},
	secure.MD5(ValidationErrorMessage("agentId", MessageModeEmptyURI)): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V014",
			Description: ValidationErrorMessage("agentId", MessageModeEmptyURI),
		},
	},
	secure.MD5("ChannelCode: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V015",
			Description: ValidationErrorMessage("channel_code", MessageModeEmpty),
		},
	},
	secure.MD5("ChannelCode: invalid pattern=channel_code"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V016",
			Description: ValidationErrorMessage("channel_code", MessageModeInvalid),
		},
	},

	secure.MD5("SendCurrency: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V017",
			Description: ValidationErrorMessage("send_currency", MessageModeEmpty),
		},
	},
	secure.MD5("SendCurrency: invalid pattern=currency"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V018",
			Description: ValidationErrorMessage("send_currency", MessageModeInvalid),
		},
	},

	secure.MD5("ReceiveCurrency: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V019",
			Description: ValidationErrorMessage("receive_currency", MessageModeEmpty),
		},
	},
	secure.MD5("ReceiveCurrency: invalid pattern=currency"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V020",
			Description: ValidationErrorMessage("receive_currency", MessageModeInvalid),
		},
	},
	secure.MD5("ReceiveCountry: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V021",
			Description: ValidationErrorMessage("receive_country", MessageModeEmpty),
		},
	},
	secure.MD5("ReceiveCountry: invalid pattern=country"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V022",
			Description: ValidationErrorMessage("receive_country", MessageModeInvalid),
		},
	},
	secure.MD5("ExchangeRate: invalid pattern=decimal"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V023",
			Description: ValidationErrorMessage("exchange_rate", MessageModeInvalid),
		},
	},
	secure.MD5("Receive.Country.Code: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V024",
			Description: ValidationErrorMessage("receive.country.code", MessageModeEmpty),
		},
	},
	secure.MD5("Receive.Country.Code: invalid pattern=country"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V025",
			Description: ValidationErrorMessage("receive.country.code", MessageModeInvalid),
		},
	},
	secure.MD5("Receive.Currency: invalid pattern=currency"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V026",
			Description: ValidationErrorMessage("receive.currency", MessageModeInvalid),
		},
	},
	secure.MD5("Receive.Currency: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V027",
			Description: ValidationErrorMessage("receive.currency", MessageModeEmpty),
		},
	},

	secure.MD5("Send.Country.Code: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V028",
			Description: ValidationErrorMessage("send.country.code", MessageModeEmpty),
		},
	},
	secure.MD5("Send.Country.Code: invalid pattern=country"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V029",
			Description: ValidationErrorMessage("send.country.code", MessageModeInvalid),
		},
	},
	secure.MD5("Send.Currency: invalid pattern=currency"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V030",
			Description: ValidationErrorMessage("send.currency", MessageModeInvalid),
		},
	},
	secure.MD5("Send.Currency: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V031",
			Description: ValidationErrorMessage("send.currency", MessageModeEmpty),
		},
	},

	secure.MD5("Sender.FirstName: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V032",
			Description: ValidationErrorMessage("sender.first_name", MessageModeEmpty),
		},
	},
	secure.MD5("Sender.MiddleName: greater than max"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V033",
			Description: ValidationErrorMessage("sender.middle_name", MessageModeLength, "20"),
		},
	},
	secure.MD5("Sender.LastName: greater than max"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V034",
			Description: ValidationErrorMessage("sender.last_name", MessageModeLength, "20"),
		},
	},
	secure.MD5("Sender.Email: invalid pattern=email"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V035",
			Description: ValidationErrorMessage("sender.last_name", MessageModeInvalid),
		},
	},
	secure.MD5("Sender.Email: greater than max"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V036",
			Description: ValidationErrorMessage("sender.email", MessageModeLength, "60"),
		},
	},
	secure.MD5("Sender.Gender: Parameter must be male or female"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V037",
			Description: ValidationErrorMessage("sender.gender", MessageModeInvalid),
			Pattern:     "",
		},
	},
	secure.MD5("Sender.MobileNumber: greater than max"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V038",
			Description: ValidationErrorMessage("sender.mobile_number", MessageModeLength, "20"),
			Pattern:     "",
		},
	},
	secure.MD5("Sender.BirthDate: invalid pattern=birthday"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V039",
			Description: ValidationErrorMessage("sender.brith_date", MessageModeInvalid),
			Pattern:     "",
		},
	},
	secure.MD5("Sender.Occupation: greater than max"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V040",
			Description: ValidationErrorMessage("sender.occupation", MessageModeLength, "100"),
			Pattern:     "",
		},
	},
	secure.MD5("Receiver.FirstName: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V041",
			Description: ValidationErrorMessage("receiver.first_name", MessageModeEmpty),
		},
	},
	secure.MD5("Receiver.MiddleName: greater than max"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V042",
			Description: ValidationErrorMessage("receiver.middle_name", MessageModeLength, "20"),
		},
	},
	secure.MD5("Receiver.LastName: greater than max"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V043",
			Description: ValidationErrorMessage("receiver.last_name", MessageModeLength, "20"),
		},
	},
	secure.MD5("Receiver.Email: invalid pattern=email"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V044",
			Description: ValidationErrorMessage("receiver.email", MessageModeInvalid),
		},
	},
	secure.MD5("Receiver.Email: greater than max"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V045",
			Description: ValidationErrorMessage("receiver.email", MessageModeLength, "60"),
		},
	},
	secure.MD5("Receiver.Gender: Parameter must be male or female"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V046",
			Description: ValidationErrorMessage("receiver.gender", MessageModeInvalid),
			Pattern:     "",
		},
	},
	secure.MD5("Receiver.MobileNumber: greater than max"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V047",
			Description: ValidationErrorMessage("receiver.mobile_number", MessageModeLength, "20"),
			Pattern:     "",
		},
	},
	secure.MD5("Receiver.BirthDate: invalid pattern=birthday"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V048",
			Description: ValidationErrorMessage("receiver.birth_date", MessageModeInvalid),
			Pattern:     "",
		},
	},
	secure.MD5("Receiver.Occupation: greater than max"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V049",
			Description: ValidationErrorMessage("receiver.occupation", MessageModeLength, "100"),
			Pattern:     "",
		},
	},
	secure.MD5(ValidationErrorMessage("providerCode", MessageModeEmptyURI)): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V050",
			Description: ValidationErrorMessage("providerCode", MessageModeEmptyURI),
		},
	},
	secure.MD5(ValidationErrorMessage("remitId", MessageModeEmptyURI)): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V051",
			Description: ValidationErrorMessage("remitId", MessageModeEmptyURI),
		},
	},
	secure.MD5(ValidationErrorMessage("bankId", MessageModeEmptyURI)): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V052",
			Description: ValidationErrorMessage("bankId", MessageModeEmptyURI),
		},
	},
	secure.MD5(ValidationErrorMessage("exchange_rate", "greater_than", "0")): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V053",
			Description: ValidationErrorMessage("exchange_rate", "greater_than", "0"),
		},
	},
	secure.MD5("PayoutAgent: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V054",
			Description: ValidationErrorMessage("send.payout_agent", MessageModeEmpty),
		},
	},
	secure.MD5("Type: invalid pattern=annex_type"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V055",
			Description: ValidationErrorMessage("send.annex_type", MessageModeInvalid),
		},
	},
	secure.MD5(ValidationErrorMessage("send.amount", "greater_than", "0")): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V056",
			Description: ValidationErrorMessage("send.amount", "greater_than", "0"),
		},
	},
	secure.MD5("Bank.Account.Number: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V057",
			Description: ValidationErrorMessage("send.bank.account.number", MessageModeEmpty),
		},
	},
	secure.MD5("RoutingParam: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V058",
			Description: ValidationErrorMessage("receiver.routing_param", MessageModeEmpty),
		},
	},
	secure.MD5("Reason: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V059",
			Description: ValidationErrorMessage("reason", MessageModeEmpty),
		},
	},
	secure.MD5("Name: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V060",
			Description: ValidationErrorMessage("name", MessageModeEmpty),
		},
	},

	secure.MD5("Code: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V061",
			Description: ValidationErrorMessage("code", MessageModeEmpty),
		},
	},

	secure.MD5("Country: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V062",
			Description: ValidationErrorMessage("country", MessageModeEmpty),
		},
	},
	secure.MD5("Bank.Branch.Code: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V063",
			Description: ValidationErrorMessage("bank.branch.code", MessageModeEmpty),
		},
	},
	secure.MD5("Bank.SWIFT: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V064",
			Description: ValidationErrorMessage("bank.swift", MessageModeEmpty),
		},
	},
	secure.MD5("Bank.IFSC: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V065",
			Description: ValidationErrorMessage("bank.ifsc", MessageModeEmpty),
		},
	},
	secure.MD5("Sender Address: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V066",
			Description: ValidationErrorMessage("sender.address", MessageModeEmpty),
		},
	},
	secure.MD5("Requested provider is not supported"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V066",
			Description: "Requested provider is not supported",
		},
	},
	secure.MD5(ValidationErrorMessage("provider", MessageModeNotSet)): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V067",
			Description: ValidationErrorMessage("provider", MessageModeNotSet),
		},
	},

	secure.MD5("Remittance record current status is preventing it from being cancelled"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V069",
			Description: "Remittance record current status is preventing it from being cancelled",
			Pattern:     ``,
		},
	},

	secure.MD5("Receiver nationality is required"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V070",
			Description: ValidationErrorMessage("receiver.nationality", MessageModeEmpty),
			Pattern:     ``,
		},
	},

	secure.MD5("Receiver address is required"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V071",
			Description: ValidationErrorMessage("receiver.address", MessageModeEmpty),
			Pattern:     ``,
		},
	},

	secure.MD5("Receiver address is required"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V072",
			Description: ValidationErrorMessage("receiver.address", MessageModeEmpty),
			Pattern:     ``,
		},
	},

	secure.MD5("Sender address is required"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V073",
			Description: ValidationErrorMessage("sender.address", MessageModeEmpty),
			Pattern:     ``,
		},
	},

	secure.MD5("Sender address is required"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V074",
			Description: ValidationErrorMessage("sender.address", MessageModeEmpty),
			Pattern:     ``,
		},
	},

	secure.MD5("Sender.MobileNumber: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V075",
			Description: ValidationErrorMessage("sender.mobile_number", MessageModeEmpty),
			Pattern:     ``,
		},
	},

	secure.MD5("Sender.BirthDate: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V076",
			Description: ValidationErrorMessage("sender.birth_date", MessageModeEmpty),
			Pattern:     ``,
		},
	},

	secure.MD5("Receiver.BirthDate: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V077",
			Description: ValidationErrorMessage("receiver.birth_date", MessageModeEmpty),
			Pattern:     ``,
		},
	},

	secure.MD5(ValidationErrorMessage("branchId", MessageModeEmptyURI)): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V078",
			Description: ValidationErrorMessage("branchId", MessageModeEmptyURI),
		},
	},

	secure.MD5("Requested provider does not match current enabled provider"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V079",
			Description: "Requested provider does not match current enabled provider",
		},
	},

	secure.MD5("Bank.Branch.Code: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V080",
			Description: ValidationErrorMessage("send.bank.branch.code", MessageModeEmpty),
		},
	},

	secure.MD5("Bank.Name: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V081",
			Description: ValidationErrorMessage("send.bank.name", MessageModeEmpty),
		},
	},

	secure.MD5("Receiver.MobileNumber: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V082",
			Description: ValidationErrorMessage("receiver.mobile_number", MessageModeEmpty),
			Pattern:     ``,
		},
	},

	secure.MD5("Sender nationality is required"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V083",
			Description: ValidationErrorMessage("sender.nationality", MessageModeEmpty),
			Pattern:     ``,
		},
	},

	secure.MD5(ValidationErrorMessage("provider", MessageModeNotSet)): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V084",
			Description: ValidationErrorMessage("provider", MessageModeNotSet),
		},
	},

	secure.MD5("CalculationMode: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V085",
			Description: ValidationErrorMessage("calculation_mode", MessageModeEmpty),
		},
	},

	secure.MD5("CalculationMode: invalid pattern=calculation_mode"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V086",
			Description: ValidationErrorMessage("calculation_mode", MessageModeInvalid),
		},
	},

	secure.MD5(ValidationErrorMessage("payout_agent", MessageModeInvalid)): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V087",
			Description: ValidationErrorMessage("payout_agent", MessageModeInvalid),
		},
	},

	secure.MD5("ReferenceNumber: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V088",
			Description: ValidationErrorMessage("reference_number", MessageModeEmpty),
		},
	},

	secure.MD5(ValidationErrorMessage("send.payout_agent", MessageModeInvalid)): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V089",
			Description: ValidationErrorMessage("send.payout_agent", MessageModeInvalid),
		},
	},

	secure.MD5("Send currency must match the configured settlement currency"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V100",
			Description: "Parameter `send.currency` does not matched the configured value",
		},
	},

	secure.MD5("Unsupported media type"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V090",
			Description: "Unsupported media type",
		},
	},

	secure.MD5(ValidationErrorMessage("receive.amount", "greater_than", "0")): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V101",
			Description: ValidationErrorMessage("receive.amount", "greater_than", "0"),
		},
	},

	secure.MD5("PayoutAgent: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V102",
			Description: ValidationErrorMessage("payout_agent", MessageModeEmpty),
		},
	},

	secure.MD5(ValidationErrorMessage("receive.amount", MessageModeInvalid)): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V103",
			Description: ValidationErrorMessage("receive.amount", MessageModeInvalid),
		},
	},

	secure.MD5("Request body is invalid"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V104",
			Description: "Request body is invalid",
		},
	},

	secure.MD5("Request body `amount` is invalid"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V105",
			Description: "Request body `amount` is invalid",
		},
	},

	secure.MD5("Send country code must match the configured settlement country"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V106",
			Description: "Parameter `send.country.code` does not matched the configured value",
		},
	},

	secure.MD5("Receive country code does not match the country code in file"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V107",
			Description: "Parameter `receive.country.code` does not matched the configured value",
		},
	},

	secure.MD5("Receive currency code does not match the currency code in file"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V108",
			Description: "Parameter `receive.currency` does not matched the configured value",
		},
	},

	secure.MD5("Destination currency not offered by payer"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V108",
			Description: "Parameter `receive.currency` does not matched the configured value",
		},
	},

	secure.MD5("SendCountry: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V110",
			Description: ValidationErrorMessage("send_country", MessageModeEmpty),
		},
	},

	secure.MD5("ERN: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V120",
			Description: ValidationErrorMessage("sender.ern", MessageModeEmpty),
		},
	},

	secure.MD5("invalid pattern=currency"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V121",
			Description: ValidationErrorMessage("currency", MessageModeInvalid),
		},
	},

	secure.MD5("invalid pattern=country"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V122",
			Description: ValidationErrorMessage("country", MessageModeInvalid),
		},
	},

	secure.MD5("Receiver.ChineseName: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V123",
			Description: ValidationErrorMessage("receiver.chinese_name", MessageModeEmpty),
		},
	},

	secure.MD5("Receiver.Identification.Number: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V124",
			Description: ValidationErrorMessage("receiver.identification.number", MessageModeEmpty),
		},
	},

	secure.MD5("Receiver.Identification.Type: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V125",
			Description: ValidationErrorMessage("receiver.identification.type", MessageModeEmpty),
		},
	},

	secure.MD5("Receiver.MobileNumber: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V126",
			Description: ValidationErrorMessage("receiver.mobile_number", MessageModeEmpty),
		},
	},

	secure.MD5("Receiver.Occupation: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V127",
			Description: ValidationErrorMessage("receiver.occupation", MessageModeEmpty),
		},
	},

	secure.MD5("Bank.Code: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V128",
			Description: ValidationErrorMessage("bank.code", MessageModeEmpty),
		},
	},

	secure.MD5("PurposeOfTransfer: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V129",
			Description: ValidationErrorMessage("purpose_of_transfer", MessageModeEmpty),
		},
	},

	secure.MD5("SourceOfIncome: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V130",
			Description: ValidationErrorMessage("source_of_income", MessageModeEmpty),
		},
	},

	secure.MD5(ValidationErrorMessage("Nationality", MessageModeInvalid)): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V131",
			Description: ValidationErrorMessage("nationality", MessageModeInvalid),
		},
	},

	secure.MD5("HashID: Parameter is a required field"): Status{
		Code: http.StatusBadRequest,
		Error: Error{
			ID:          "V132",
			Description: ValidationErrorMessage("hash_id", MessageModeEmpty),
		},
	},
}
```
BadRequestErrors Validation errors

```go
var FillableChannelFields = []string{
	"hash_id",
	"name",
	"code",
}
```
FillableChannelFields ...

```go
var FillableEnumerationFields = []string{
	"name",
}
```
FillableEnumerationFields ...

```go
var FillableProviderFields = []string{
	"prefix",
	"code",
	"name",
	"description",
}
```
FillableProviderFields ...

```go
var FillableSampleFields = []string{
	"name",
}
```
FillableSampleFields ...

```go
var ForbiddenRequestErrors = map[string]Status{
	secure.MD5("Remittance record current status is preventing it from being confirmed"): Status{
		Code: http.StatusForbidden,
		Error: Error{
			ID:          "V068",
			Description: "Remittance record current status is preventing it from being confirmed",
			Pattern:     ``,
		},
	},
	secure.MD5("Destination service provider not allowed"): Status{
		Code: http.StatusForbidden,
		Error: Error{
			ID:          "V091",
			Description: "The receiver E-Wallet system does not allow this calculation mode",
			Pattern:     ``,
		},
	},
	secure.MD5("We only support calculation mode `source` at the moment"): Status{
		Code: http.StatusForbidden,
		Error: Error{
			ID:          "V106",
			Description: "The calculation mode `receive` is unvailable at the moment",
			Pattern:     ``,
		},
	},
}
```
ForbiddenRequestErrors Validation errors

```go
var GatewayTimeoutErrors = map[string]Status{
	secure.MD5("Gateway Timeout"): Status{
		Code: http.StatusGatewayTimeout,
		Error: Error{
			ID:          "S504",
			Description: "Gateway Timeout",
			Pattern:     "",
		},
	},
}
```
GatewayTimeoutErrors ...

```go
var InternalServerErrors = map[string]Status{
	secure.MD5("Internal Server Error"): Status{
		Code: http.StatusInternalServerError,
		Error: Error{
			ID:          "S000",
			Description: "Internal Server Error",
		},
	},

	secure.MD5("Failed to fetch config"): Status{
		Code: http.StatusInternalServerError,
		Error: Error{
			ID:          "S001",
			Description: "Failed to fetch config",
		},
	},
	secure.MD5("One of the Amazon Web Services environment variable was not set"): Status{
		Code: http.StatusInternalServerError,
		Error: Error{
			ID:          "S002",
			Description: "One of the Amazon Web Services environment variable was not set",
		},
	},
	secure.MD5("Unable to download config file"): Status{
		Code: http.StatusInternalServerError,
		Error: Error{
			ID:          "S003",
			Description: "Unable to download config file",
		},
	},
	secure.MD5("Unable to upload config file"): Status{
		Code: http.StatusInternalServerError,
		Error: Error{
			ID:          "S004",
			Description: "Unable to upload config file",
		},
	},
}
```
InternalServerErrors Internal Server Error

```go
var ResourceNotFoundErrors = map[string]Status{
	secure.MD5(ValidationErrorMessage("Channel", MessageModeNotFound)): Status{
		Code: http.StatusNotFound,
		Error: Error{
			ID:          "R401",
			Description: ValidationErrorMessage("Channel", MessageModeNotFound),
		},
	},
	secure.MD5(ValidationErrorMessage("Agent", MessageModeNotFound)): Status{
		Code: http.StatusNotFound,
		Error: Error{
			ID:          "R402",
			Description: ValidationErrorMessage("Agent", MessageModeNotFound),
		},
	},
	secure.MD5(ValidationErrorMessage("Remittance", MessageModeNotFound)): Status{
		Code: http.StatusNotFound,
		Error: Error{
			ID:          "R403",
			Description: ValidationErrorMessage("Remittance", MessageModeNotFound),
		},
	},
	secure.MD5(ValidationErrorMessage("Bank", MessageModeNotFound)): Status{
		Code: http.StatusNotFound,
		Error: Error{
			ID:          "R404",
			Description: ValidationErrorMessage("Bank", MessageModeNotFound),
		},
	},
	secure.MD5(ValidationErrorMessage("Branch", MessageModeNotFound)): Status{
		Code: http.StatusNotFound,
		Error: Error{
			ID:          "R405",
			Description: ValidationErrorMessage("Branch", MessageModeNotFound),
		},
	},
	secure.MD5(ValidationErrorMessage("Provider", MessageModeNotFound)): Status{
		Code: http.StatusNotFound,
		Error: Error{
			ID:          "R407",
			Description: ValidationErrorMessage("Provider", MessageModeNotFound),
		},
	},
}
```
ResourceNotFoundErrors ...

```go
var ServiceUnavailableErrors = map[string]Status{
	secure.MD5("Destination temporarily unavailable"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R000",
			Description: "The receiver E-Wallet system is currently unavailable",
			Pattern:     "",
		},
	},
	secure.MD5("Payer is inactive in your account"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R000",
			Description: "The receiver E-Wallet system is currently unavailable",
			Pattern:     "",
		},
	},
	secure.MD5("\"Security failure\""): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R000",
			Description: "The receiver E-Wallet system is currently unavailable",
			Pattern:     "",
		},
	},

	secure.MD5("Temporarily unavailable"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R000",
			Description: "The receiver E-Wallet system is currently unavailable",
			Pattern:     "",
		},
	},

	secure.MD5("Maximum number of transactions exceeded by receiver"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R001",
			Description: "Maximum number of transactions exceeded by receiver",
			Pattern:     "",
		},
	},
	secure.MD5("Maximum cumulative amount of transactions exceeded by receiver"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R001",
			Description: "Maximum number of transactions exceeded by receiver",
			Pattern:     "",
		},
	},
	secure.MD5("Maximum number of transactions exceeded by receiver"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R001",
			Description: "Maximum number of transactions exceeded by receiver",
			Pattern:     "",
		},
	},
	secure.MD5("Maximum cumulative amount of transactions exceeded by receiving network"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R001",
			Description: "Maximum number of transactions exceeded by receiver",
			Pattern:     "",
		},
	},
	secure.MD5("Maximum number of transactions exceeded by receiving network"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R001",
			Description: "Maximum number of transactions exceeded by receiver",
			Pattern:     "",
		},
	},

	secure.MD5("Maximum number of transactions exceeded by sender"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R002",
			Description: "Maximum number of transactions exceeded by sender",
			Pattern:     "",
		},
	},
	secure.MD5("Maximum number of transactions exceeded by sending network"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R002",
			Description: "Maximum number of transactions exceeded by sender",
			Pattern:     "",
		},
	},
	secure.MD5("Maximum cumulative amount of transactions exceeded by sending network"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R002",
			Description: "Maximum number of transactions exceeded by sender",
			Pattern:     "",
		},
	},
	secure.MD5("Maximum cumulative amount of transactions exceeded by sender"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R002",
			Description: "Maximum number of transactions exceeded by sender",
			Pattern:     "",
		},
	},
	secure.MD5("Maximum number of transactions exceeded by sender"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R002",
			Description: "Maximum number of transactions exceeded by sender",
			Pattern:     "",
		},
	},

	secure.MD5("Unknown receiver"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R003",
			Description: "The receiver E-Wallet system was unable to find the recipient",
			Pattern:     "",
		},
	},
	secure.MD5("Destination system route error"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R003",
			Description: "The receiver E-Wallet system was unable to find the recipient",
			Pattern:     "",
		},
	},
	secure.MD5("Invalid input value"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R003",
			Description: "The receiver E-Wallet system was unable to find the recipient",
			Pattern:     "",
		},
	},
	secure.MD5("Remittance receiver does not exist in receiving rail"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R003",
			Description: "The receiver E-Wallet system was unable to find the recipient",
			Pattern:     "",
		},
	},
	secure.MD5("Remittance destination temporarily unavailable"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R004",
			Description: "The receiver E-Wallet system temporarily unavailable",
			Pattern:     "",
		},
	},
	secure.MD5("Destination temporarily unavailable"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R004",
			Description: "The receiver E-Wallet system temporarily unavailable",
			Pattern:     "",
		},
	},

	secure.MD5("Temporarily unavailable"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R004",
			Description: "The receiver E-Wallet system temporarily unavailable",
			Pattern:     "",
		},
	},
	secure.MD5("Failed to credit remittance transaction"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R005",
			Description: "The receiver E-Wallet system was unable to credit the recipient",
			Pattern:     "",
		},
	},
	secure.MD5("The receiver E-Wallet system failed cancel the transaction"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R006",
			Description: "The receiver E-Wallet system failed cancel the transaction",
			Pattern:     "",
		},
	},
	secure.MD5("The receiver E-Wallet system rejected the transaction"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R007",
			Description: "The receiver E-Wallet system rejected the transaction",
			Pattern:     "",
		},
	},
	secure.MD5("The receiver E-Wallet system rejected the transaction"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R008",
			Description: "The receiver E-Wallet system rejected the transaction",
			Pattern:     "",
		},
	},
	secure.MD5("Service unavailable"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R009",
			Description: "The receiver E-Wallet service is currently unavailable",
			Pattern:     "",
		},
	},
	secure.MD5("Receiver crediting error"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R010",
			Description: "The receiver E-Wallet system was unable to credit the user",
		},
	},
	secure.MD5("Request already in progress"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R011",
			Description: "Possible duplicate request detected",
		},
	},

	secure.MD5("Missing required fields"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R012",
			Description: "Remittance request is missing some required field",
			Pattern:     "",
		},
	},
	secure.MD5("REJECTED"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R013",
			Description: "The receiver E-Wallet system rejected the transaction",
			Pattern:     "",
		},
	},
	secure.MD5("Cannot cancel remittance available for refund"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R014",
			Description: "Cannot cancel remittance available for refund",
			Pattern:     "",
		},
	},
	secure.MD5("Cannot cancel already cancelled remittance"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R015",
			Description: "Cannot cancel already cancelled remittance",
			Pattern:     "",
		},
	},
	secure.MD5("Cannot cancel paid remittance"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R016",
			Description: "Cannot cancel paid remittance",
			Pattern:     "",
		},
	},
	secure.MD5("Cannot cancel pending cancellation remittance"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R017",
			Description: "Cannot cancel pending cancellation remittance",
			Pattern:     "",
		},
	},
	secure.MD5("Cannot cancel refunded remittance"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R018",
			Description: "Cannot cancel refunded remittance",
			Pattern:     "",
		},
	},
	secure.MD5("Cannot cancel rejected remittance"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R019",
			Description: "Cannot cancel rejected remittance",
			Pattern:     "",
		},
	},
	secure.MD5("Remittance cancellation failed"): Status{
		Code: http.StatusServiceUnavailable,
		Error: Error{
			ID:          "R020",
			Description: "Remittance cancellation failed",
			Pattern:     "",
		},
	},
}
```
ServiceUnavailableErrors The server is currently unable to handle the request
due to a temporary overloading or maintenance of the server

```go
var SortChannelFields = map[string]string{}
```
SortChannelFields sortable fields

```go
var SortEnumerationFields = map[string]string{}
```
SortEnumerationFields sortable fields

```go
var SortProviderFields = map[string]string{}
```
SortProviderFields sortable fields

```go
var SortSampleFields = map[string]string{}
```
SortSampleFields sortable fields

```go
var StatusUnprocessableEntityErrors = map[string]Status{
	secure.MD5("Cannot parse empty Body content"): Status{
		Code: http.StatusUnprocessableEntity,
		Error: Error{
			ID:          "S422",
			Description: "Cannot parse empty Body content",
		},
	},
}
```
StatusUnprocessableEntityErrors ...

```go
var UpdatableChannelFields = []string{
	"name",
	"code",
}
```
UpdatableChannelFields ...

```go
var UpdatableEnumerationFields = []string{
	"name",
}
```
UpdatableEnumerationFields ...

```go
var UpdatableProviderFields = []string{
	"prefix",
	"code",
	"name",
	"description",
}
```
UpdatableProviderFields ...

```go
var UpdatableSampleFields = []string{
	"name",
}
```
UpdatableSampleFields ...

#### func  InitErrors

```go
func InitErrors()
```
InitErrors initializes the properties of Errors

#### func  Messages

```go
func Messages() map[string]Status
```
Messages all errors

#### func  ValidationErrorMessage

```go
func ValidationErrorMessage(field string, mode string, ranges ...string) string
```
ValidationErrorMessage generate generic validation Errors

#### type Channel

```go
type Channel struct {
	gorm.Model `json:"-"`
	ID         string      `db:"id" json:"-"`
	HashID     string      `db:"hash_id" json:"id"`
	Code       string      `db:"code" json:"code" validate:"required"`
	Name       string      `db:"name" json:"name" validate:"required"`
	Date       entity.Date `json:"date,omitempty"`
	CreatedAt  string      `db:"created_at" json:"-"`
	UpdatedAt  string      `db:"updated_at" json:"-"`
}
```

Channel model

#### func  NewChannel

```go
func NewChannel(db *sqlx.DB) (*Channel, error)
```
NewChannel ...

#### func (*Channel) CountAll

```go
func (me *Channel) CountAll() (count int64, err error)
```
CountAll count all the records

#### func (*Channel) Create

```go
func (me *Channel) Create() (result sql.Result, err error)
```
Create ....

#### func (*Channel) Delete

```go
func (me *Channel) Delete() (result sql.Result, err error)
```
Delete delete record

#### func (*Channel) Filter

```go
func (me *Channel) Filter(query string, args []interface{}) (Channels, error)
```
Filter ...

#### func (*Channel) Find

```go
func (me *Channel) Find() (Channels, error)
```
Find returns the list of records meeting the filter requirement filters, limit,
offset, order

#### func (*Channel) FindFirst

```go
func (me *Channel) FindFirst() error
```
FindFirst find first

#### func (*Channel) FirstOrCreate

```go
func (me *Channel) FirstOrCreate() (err error)
```
FirstOrCreate ... find first or create

#### func (*Channel) GenerateHashID

```go
func (me *Channel) GenerateHashID() string
```
GenerateHashID generate hash id

#### func (*Channel) GetByCode

```go
func (me *Channel) GetByCode(code string) error
```
GetByCode ...

#### func (*Channel) Patch

```go
func (me *Channel) Patch(model *Channel)
```
Patch ...

#### func (*Channel) Update

```go
func (me *Channel) Update() (sql.Result, error)
```
Update ....

#### type Channels

```go
type Channels []Channel
```

Channels list of Channel

#### type Countries

```go
type Countries []Country
```

Countries list of Country

#### type Country

```go
type Country struct {
	gorm.Model    `json:"-"`
	ID            string      `db:"id" json:"-"`
	ISO2          string      `db:"iso2" json:"iso2"`
	ISO3166Alpha3 string      `db:"iso3166_alpha3" json:"iso3166_alpha3"`
	ISO4217       string      `db:"iso4217" json:"iso4217"`
	Name          string      `db:"name" json:"name"`
	Nationality   string      `db:"nationality" json:"nationality"`
	Date          entity.Date `json:"date,omitempty"`
	CreatedAt     string      `db:"created_at" json:"-"`
}
```

Country model

#### func  NewCountry

```go
func NewCountry(db *sqlx.DB) (*Country, error)
```
NewCountry ...

#### func (*Country) CountAll

```go
func (me *Country) CountAll() (count int64, err error)
```
CountAll count all the records

#### func (*Country) Find

```go
func (me *Country) Find() (Countries, error)
```
Find returns the list of records meeting the filter requirement filters, limit,
offset, order

#### func (*Country) FindFirst

```go
func (me *Country) FindFirst() error
```
FindFirst find first

#### type Enumeration

```go
type Enumeration struct {
	gorm.Model     `json:"-"`
	ID             string           `db:"id" json:"-"`
	Name           string           `db:"name" json:"-"`
	ConfigAsObject *json.RawMessage `json:"configuration,omitempty"`
	S3             aws.AWS          `db:"-" json:"-"`
	Date           entity.Date      `json:"date,omitempty"`
	CreatedAt      string           `db:"created_at" json:"-"`
}
```

Enumeration model

#### func  NewEnumeration

```go
func NewEnumeration(db *sqlx.DB) (*Enumeration, error)
```
NewEnumeration ...

#### func (*Enumeration) CompleteFilename

```go
func (me *Enumeration) CompleteFilename(filename string) string
```
CompleteFilename get complete file name

#### func (*Enumeration) CountAll

```go
func (me *Enumeration) CountAll() (count int64, err error)
```
CountAll count all the records

#### func (*Enumeration) Create

```go
func (me *Enumeration) Create() (result sql.Result, err error)
```
Create ....

#### func (*Enumeration) Delete

```go
func (me *Enumeration) Delete() (result sql.Result, err error)
```
Delete delete record

#### func (*Enumeration) Download

```go
func (me *Enumeration) Download(filename string) (bool, error)
```
Download file from s3

#### func (*Enumeration) Find

```go
func (me *Enumeration) Find() (Enumerations, error)
```
Find returns the list of records meeting the filter requirement filters, limit,
offset, order

#### func (*Enumeration) FindFirst

```go
func (me *Enumeration) FindFirst() error
```
FindFirst find first

#### func (*Enumeration) FirstOrCreate

```go
func (me *Enumeration) FirstOrCreate() (err error)
```
FirstOrCreate ... find first or create

#### func (*Enumeration) FormatFilename

```go
func (me *Enumeration) FormatFilename(value string) string
```
FormatFilename ...

#### func (*Enumeration) GetEnumeration

```go
func (me *Enumeration) GetEnumeration() (conf *json.RawMessage, err error)
```
GetEnumeration retrive config for the enumeration

#### func (*Enumeration) GetFilename

```go
func (me *Enumeration) GetFilename() string
```
GetFilename get config filename

#### func (*Enumeration) InitS3

```go
func (me *Enumeration) InitS3() error
```
InitS3 init S3

#### func (*Enumeration) Patch

```go
func (me *Enumeration) Patch(model *Enumeration)
```
Patch ...

#### func (*Enumeration) SetEnumeration

```go
func (me *Enumeration) SetEnumeration(data *json.RawMessage) (conf *json.RawMessage, err error)
```
SetEnumeration sets the config for enumeration

#### func (*Enumeration) Update

```go
func (me *Enumeration) Update() (sql.Result, error)
```
Update ....

#### func (*Enumeration) Upload

```go
func (me *Enumeration) Upload(filePath string, filename string) (bool, error)
```
Upload file to s3

#### type Enumerations

```go
type Enumerations []Enumeration
```

Enumerations list of Enumeration

#### type Error

```go
type Error struct {
	ID          string        `json:"id"`
	Description string        `json:"message"`
	Pattern     string        `json:"pattern,omitempty"`
	Links       []entity.Link `json:"links"`
}
```

Error struct for errors

#### func  Errors

```go
func Errors(code string) (Error, error)
```
Errors get error code

#### func  NewError

```go
func NewError() *Error
```
NewError creates a new blank error model

#### func (*Error) GetAllByFilter

```go
func (e *Error) GetAllByFilter(desc string, limit int, offset int, order []string, filter gorm.Conditions) ([]Error, error)
```
GetAllByFilter filters Errors based on condition

#### type Provider

```go
type Provider struct {
	gorm.Model  `json:"-"`
	ConsumerKey string           `db:"-" json:"-"`
	ID          string           `db:"id" json:"-"`
	Product     string           `db:"-" json:"-"`
	Prefix      string           `db:"prefix" json:"prefix" validate:"required"`
	Code        string           `db:"code" json:"code" validate:"required"`
	Name        string           `db:"name" json:"name" validate:"required"`
	Description string           `db:"description" json:"description"`
	Date        modelEntity.Date `json:"date,omitempty"`
	CreatedAt   string           `db:"created_at" json:"-"`
	UpdatedAt   string           `db:"updated_at" json:"-"`
}
```

Provider model

#### func  NewProvider

```go
func NewProvider(db *sqlx.DB) (*Provider, error)
```
NewProvider ...

#### func (*Provider) CountAll

```go
func (me *Provider) CountAll() (count int64, err error)
```
CountAll count all the records

#### func (*Provider) Create

```go
func (me *Provider) Create() (result sql.Result, err error)
```
Create ....

#### func (*Provider) Delete

```go
func (me *Provider) Delete() (result sql.Result, err error)
```
Delete delete record

#### func (*Provider) Find

```go
func (me *Provider) Find() (Providers, error)
```
Find returns the list of records meeting the filter requirement filters, limit,
offset, order

#### func (*Provider) FindFirst

```go
func (me *Provider) FindFirst() error
```
FindFirst find first

#### func (*Provider) FirstOrCreate

```go
func (me *Provider) FirstOrCreate() (err error)
```
FirstOrCreate ... find first or create

#### func (*Provider) GetByCode

```go
func (me *Provider) GetByCode(code string) error
```
GetByCode ...

#### func (*Provider) Patch

```go
func (me *Provider) Patch(model *Provider)
```
Patch ...

#### func (*Provider) Update

```go
func (me *Provider) Update() (sql.Result, error)
```
Update ....

#### type Providers

```go
type Providers []Provider
```

Providers list of Provider

#### type Sample

```go
type Sample struct {
	gorm.Model `json:"-"`
	ID         string      `db:"id" json:"-"`
	Name       string      `db:"name" json:"-"`
	Date       entity.Date `json:"date,omitempty"`
	CreatedAt  string      `db:"created_at" json:"-"`
	UpdatedAt  string      `db:"updated_at" json:"-"`
}
```

Sample model

#### func  NewSample

```go
func NewSample(db *sqlx.DB) (*Sample, error)
```
NewSample ...

#### func (*Sample) CountAll

```go
func (me *Sample) CountAll() (count int64, err error)
```
CountAll count all the records

#### func (*Sample) Create

```go
func (me *Sample) Create() (result sql.Result, err error)
```
Create ....

#### func (*Sample) Delete

```go
func (me *Sample) Delete() (result sql.Result, err error)
```
Delete delete record

#### func (*Sample) Find

```go
func (me *Sample) Find() (Samples, error)
```
Find returns the list of records meeting the filter requirement filters, limit,
offset, order

#### func (*Sample) FindFirst

```go
func (me *Sample) FindFirst() error
```
FindFirst find first

#### func (*Sample) FirstOrCreate

```go
func (me *Sample) FirstOrCreate() (err error)
```
FirstOrCreate ... find first or create

#### func (*Sample) GenerateHashID

```go
func (me *Sample) GenerateHashID() string
```
GenerateHashID generate hash id

#### func (*Sample) GetByHashID

```go
func (me *Sample) GetByHashID(id string) error
```
GetByHashID ...

#### func (*Sample) Patch

```go
func (me *Sample) Patch(model *Sample)
```
Patch ...

#### func (*Sample) Update

```go
func (me *Sample) Update() (result sql.Result, err error)
```
Update modify the existing address record

#### type Samples

```go
type Samples []Sample
```

Samples list of Sample

#### type Status

```go
type Status struct {
	Code  int
	Error Error
}
```

Status ...

#### func  Message

```go
func Message(code string) (Status, error)
```
Message get status error
