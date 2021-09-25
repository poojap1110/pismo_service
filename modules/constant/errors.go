package constant

const (
	DynamicErrorSeparator = ": "
)

const (
	// ErrorCodeInternalServer ....
	ErrorCodeInternalServer = 500

	// ErrorCodeBadRequest ....
	ErrorCodeBadRequest = 400

	// ErrorCodeForbidden ....
	ErrorCodeForbidden = 401

	// ErrorCodeResourceNotFound ....
	ErrorCodeResourceNotFound = 404

	// SuccessCodeOk ...
	SuccessCodeOk = 200
)

const (

	// ErrorParameterRequired ....
	ErrorParameterRequired = "EUP0400101"

	// ErrorNotAcceptableValues ...
	ErrorNotAcceptableValues = "EUP0400102"

	// ErrorInvalidPattern ...
	ErrorInvalidPattern = "EUP0400103"

	// ErrorParameterNotEqual ...
	ErrorParameterNotEqual = "EUP0400104"

	// ErrorParameterRequiredLength ...
	ErrorParameterRequiredLength = "EUP0400105"

	// ErrorParameterMinValue ...
	ErrorParameterMinValue = "EUP0400106"

	// ErrorParameterMaxValue ...
	ErrorParameterMaxValue = "EUP0400107"

	// ErrorParameterMinLengthValue ...
	ErrorParameterMinLengthValue = "EUP0400108"

	// ErrorParameterMaxLengthValue ...
	ErrorParameterMaxLengthValue = "EUP0400109"

	// ErrorInvalidPassword ...
	ErrorInvalidPassword = "EUP0400110"

	// ErrorInvalidParamValue ...
	ErrorInvalidParamValue = "EUP0400111"

	// ResourceNotFound ....
	ResourceNotFound = "EUP0404001"

	// InvalidBody
	InternalServerError = "EUP0500000"

	// InvalidBody ...
	ErrorInvalidBody = "EUP0400112"

	// GatewayTimeout ...
	GatewayTimeout = "EUP0504001"

	// Forbidden ...
	Forbidden = "EUP0403000"

	// ErrorCodeInternalServerProductNotFound ...
	ErrorCodeInternalServerProductNotFound = "EUP0500001"

	// ErrorCodeInternalServerProductNotDefined ...
	ErrorCodeInternalServerProductNotDefined = "EUP0500002"

	// ErrorNotAcceptableValues ...
	ErrorExternalService = "EUP0500004"

	// UserNotFound ...
	AccountNotFound = "EUP0400113"

	// UserPasswordFirstLastName ...
	UserPasswordFirstLastName = "EUP0400114"

	// Unauthorized ...
	Unauthorized = "EPS0401000"

	// InvalidAuthToken ...
	InvalidAuthToken = "EUP0401200"

	// ErrorCodeInternalServerProductNotDefined ...
	AuthDefinitionMissing = "EUP0500003"

	// InvalidConsumerCredentials
	InvalidConsumerCredentials = "EUP0401300"

	// AccessDeniedForThisConsumer
	AccessDenied = "EUP0401311"
)

// Application Specific codes
const (
	// EmptyRequest
	EmptyRequest = "EUP0004000"

	// BadRequest
	BadRequest = "EUP0004001"
)

const (
	// ErrorCodeNotFoundMsg ....
	ErrorCodeNotFoundMsg = "Error code not found"

	PartnerDeleteSuccessMsg = "Partner delete"
	// EmptyVirtualAccountNumber
	EmptyVirtualAccountNumber = "EUP0400200"

	// InvalidVirtualAccountNumberLength
	InvalidVirtualAccountNumberLength = "EUP0400201"

	// InvalidClientCode
	InvalidClientCode = "EUP0400202"

	// InvalidPurposeCode
	InvalidPurposeCode = "EUP0400203"

	// DuplicateVirtualAccountNumber
	DuplicateVirtualAccountNumber = "EUP0400204"

	// InvalidProvider
	InvalidProvider = "EUP0400205"

	// AlreadyAssigned
	AlreadyAssigned = "EUP0400206"

	// CallbackDecodeFailed
	CallbackDecodeFailed = "EUP0400207"

	// Record already exist ...
	ErrorRecordExist = "EUP0400208"

	// InvalidProgramCode
	InvalidProgramCode = "EUP0400209"

	// InvalidAccount
	InvalidAccount = "EUP0400210"

	InvalidAmount = "EUP0400211"

	// InvalidBankAccount
	InvalidBankAccount = "EUP0400212"

	// DatabaseError ....
	DatabaseError = "ERS0500001"

	RecordNotFound = "EUP0404002"

	InvalidInstruction = "EUP0400213"

	InvalidConsent = "EUP0400214"

	VirtualAccountOutofBound = "EUP0400215"

	SettlementNotFound = "EUP0400216"

	//FailedToGenerateCSV
	FailedToGenerateCSV = "EUP0400217"

	SettlementFailed = "EUP0400218"

	InvalidDate = "EUP0400219"

	SingleSettlementAccount = "EUP0400220"

	InstructionsNotFound = "EUP0400221"

	SettlementAccountExist = "EUP0400222"

	ConsumerAccountExist = "EUP0400223"

	InvalidExpiryDate = "EUP0400224"

	InvalidFromDate = "EUP0400225"

	InvalidToDate = "EUP0400226"

	InvalidOrderFilter = "EUP0400227"

	FeeConfigExist = "EUP0400228"

	FeeConfigNotExist = "EUP0400229"

	// InvalidTimezone
	InvalidTimezone = "EUP0400230"

	VaConfigNotExist = "EUP0400231"
)
