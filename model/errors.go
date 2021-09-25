package model

import "net/http"

const (
	validationErrorPrefix          = "V"
	unauthorizedErrorPrefix        = "U"
	forbiddenErrorPrefix           = "F"
	notFoundErrorPrefix            = "NF"
	unprocessableEntityErrorPrefix = "UE"
	internalServerErrorPrefix      = "IS"
	serviceUnavailableErrorPrefix  = "SU"
	gatewayTimeoutErrorPrefix      = "GT"
)

type errorStatus struct {
	rawKey           string
	errorCode        int
	errorPrefix      string
	errorDescription string
}

type errorStatuses []errorStatus

// 400
var BadRequestErrors = errorStatuses{
	{
		"invalid_body",
		http.StatusBadRequest,
		validationErrorPrefix,
		"Request body contains an invalid value",
	},
	{
		"Invalid Headers value",
		http.StatusBadRequest,
		validationErrorPrefix,
		"Invalid Headers value",
	},
	{
		"RoleName is already in use",
		http.StatusBadRequest,
		validationErrorPrefix,
		"RoleName is already in use",
	},
	{
		"TenantRole record not found",
		http.StatusBadRequest,
		validationErrorPrefix,
		"TenantRole record not found",
	},
}

// 401
var UnauthorizedErrors = errorStatuses{
	{
		"access_denied",
		http.StatusUnauthorized,
		unauthorizedErrorPrefix,
		"Access denied",
	},
}

// 403
var ForbiddenErrors = errorStatuses{

}

// 404
var ResourceNotFoundErrors = errorStatuses{
	{
		"resource_not_found",
		http.StatusNotFound,
		notFoundErrorPrefix,
		"Resource not found",
	},
	{
		"record_not_found",
		http.StatusNotFound,
		notFoundErrorPrefix,
		"Record not found",
	},
}

// 422
var UnprocessableEntityErrors = errorStatuses{
	{
		"cannot_parse_empty_body",
		http.StatusUnprocessableEntity,
		unprocessableEntityErrorPrefix,
		"Cannot parse empty Body content",
	},
}

// 500
var InternalServerErrors = errorStatuses{
	{
		"internal_server_error",
		http.StatusInternalServerError,
		internalServerErrorPrefix,
		"Internal Server Error",
	},
}

// 503
var ServiceUnavailableErrors = errorStatuses{
	{
		"destination_unavailable",
		http.StatusUnprocessableEntity,
		serviceUnavailableErrorPrefix,
		"Destination unavailable",
	},
}

// 504
var GatewayTimeoutErrors = errorStatuses{
	{
		"gateway_timeout",
		http.StatusUnprocessableEntity,
		gatewayTimeoutErrorPrefix,
		"Gateway Timeout",
	},
}