package errs

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

// Error constants for 400 bad request
const (
	ErrParameterRequired            = "EAUB400001"
	ErrParameterNumeric             = "EAUB400002"
	ErrParameterInvalid             = "EAUB400003"
	ErrRequestBodyInvalid           = "EAUB400004"
	ErrURIParameterMissing          = "EAUB400005"
	ErrUnsupportedRequest           = "EAUB400006"
	ErrValidation                   = "EAUB400007"
	ErrUserNotFound                 = "EAUB400008"
	ErrParameterRequiredLength      = "EAUB400009"
	ErrParameterMinValue            = "EAUB400010"
	ErrParameterMaxValue            = "EAUB400011"
	ErrParameterMinLengthValue      = "EAUB400012"
	ErrParameterMaxLengthValue      = "EAUB400013"
	ErrParameterInvalidPattern      = "EAUB400014"
	ErrParameterInvalidFormat       = "EAUB400015"
	ErrParameterShouldBeEqual       = "EAUB400016"
	ErrParameterNotAcceptableValues = "EAUB400017"
	ErrUserAlreadyExist             = "EAUB400018"
	ErrInvalidConsumerKey           = "EAUB400019"
	ErrUnsupportedValue             = "EAUB400020"
	ErrProfileIncompleteForKyc      = "EAUB400021"
)

// Error constants for 401 Unauthorized
const (
	ErrAccessDenied = "EAUB401001"
)

// Error constants for 404 bad request
const (
	ErrResourceNotFound = "EAUB404001"
	ErrCodeNotFound     = "EAUB404002"
)

// Error constants for 500 bad request
const (
	ErrInternalServerError  = "EAUB500001"
	ErrInternalAppError     = "EAUB500002"
	ErrInternalDBError      = "EAUB500003"
	ErrExtenalAppError      = "EAUB500004"
	ErrRegistrationRejected = "EAUB500005"
)

// Error constants for 504 Gateway timeout
const (
	ErrGatewayTimeout = "EAUB504001"
)

// Error constants for 422 Unprocessable Entity
const (
	ErrEmptyBodyContent = "ESOR422001"
)

// Errors - maps of error with the error code
type Errors map[int]map[string]string

// Error - return error response
type Error struct {
	Code     string `json:"code"`
	HTTPCode int    `json:"http_code,string"`
	Message  string `json:"message"`
}

// ErrorLink ...
type ErrorLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

// ErrorResponse - Use to trow the errors to users
type ErrorResponse struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Link    ErrorLink `json:"link"`
}

var errs *Errors
var allErrs map[string]Error

// Init function for errs package
func init() {
	errs = &Errors{
		http.StatusBadRequest: {
			ErrParameterRequired:            "Parameter `%s` is a required field",
			ErrParameterNumeric:             "Parameter `%s`  only allows numeric values",
			ErrParameterInvalid:             "Parameter `%s` contains an invalid value",
			ErrRequestBodyInvalid:           "Invalid Content-Type or one or more request parameters is an invalid type",
			ErrURIParameterMissing:          "URI parameter `%s` is missing",
			ErrUnsupportedRequest:           "Unsupported input request",
			ErrValidation:                   "Validation Error, `%s`",
			ErrUserNotFound:                 "User record not found",
			ErrParameterRequiredLength:      "Parameter `%s` does not match the required length of `%s`",
			ErrParameterMinValue:            "Minimum value allowed for parameter `%s` is `%s`",
			ErrParameterMaxValue:            "Maximum value allowed for parameter `%s` is `%s`",
			ErrParameterMinLengthValue:      "Minimum length allowed for parameter `%s` is `%s` characters",
			ErrParameterMaxLengthValue:      "Maximum length allowed for parameter `%s` is `%s` characters",
			ErrParameterInvalidPattern:      "Parameter `%s` value does not match the pattern provided",
			ErrParameterInvalidFormat:       "Parameter `%s` contains an invalid format : %s",
			ErrParameterShouldBeEqual:       "Parameter `%s` should be equal with value : %s",
			ErrParameterNotAcceptableValues: "Parameter `%s` only accepts the following values : [%s]",
			ErrUserAlreadyExist:             "`%s` is already in use",
			ErrInvalidConsumerKey:           "Invalid consumer",
			ErrUnsupportedValue:             "Unsupported %s",
			ErrProfileIncompleteForKyc:      "User profile is incomplete, missing following values: [%s]",
		},
		http.StatusUnauthorized: {
			ErrAccessDenied: "Access denied",
		},
		http.StatusNotFound: {
			ErrResourceNotFound: "Resource Not found",
			ErrCodeNotFound:     "Error code not found",
		},
		http.StatusInternalServerError: {
			ErrInternalServerError:  "Internal Server Error",
			ErrInternalAppError:     "Internal Application Error, `%s`",
			ErrInternalDBError:      "Database Error, `%s`",
			ErrExtenalAppError:      "External library/application Error from %s : %s",
			ErrRegistrationRejected: "Cannot process registration successfully",
		},
		http.StatusGatewayTimeout: {
			ErrGatewayTimeout: "Gateway Timeout",
		},
		http.StatusUnprocessableEntity: {
			ErrEmptyBodyContent: "Cannot parse empty body",
		},
	}
	allErrs = make(map[string]Error)

	for httpcode, err := range *errs {
		for code, msg := range err {
			tmp := &Error{Code: code, HTTPCode: httpcode, Message: msg}
			allErrs[code] = *tmp
		}
	}
}

// GetErrorByCode ...
func GetErrorByCode(code string) (res Error, err error) {
	var ok bool
	if res, ok = allErrs[code]; !ok {
		err = errors.New(ErrCodeNotFound)
		return
	}
	return
}

// GetErrors ...
func GetErrors() (res map[string]Error) {
	res = allErrs
	return
}

// FormateErrorResponse ...
func FormateErrorResponse(mErr Error, val ...interface{}) (res ErrorResponse) {
	path := os.Getenv("APP_EXTERNAL_DOMAIN") + "/v1/error/"

	if len(val) > 0 {
		mErr.Message = fmt.Sprintf(mErr.Message, val...)
	}

	errRes := &ErrorResponse{
		Code:    mErr.Code,
		Message: mErr.Message,
		Link: ErrorLink{
			Rel:  mErr.Code,
			Href: path + mErr.Code,
		},
	}
	return *errRes
}
