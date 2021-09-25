package config

import (
	"net/http"
	"os"
	"strings"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/tcp"
	payloadConfig "bitbucket.org/matchmove/go-payloads/config"
)

const (
	LogName               = "ConfigService"
	ResourceAPIVersion    = "x1.0.0.0"
	ResourceHeaderVersion = 3

	// SuccessCode
	SuccessCode = 200

	// HeaderAdminEmail
	HeaderAdminEmail = `X-MM-Admin-Email`

	// HeaderDeviceDetails
	HeaderDeviceDetails = `X-MM-Device-Details`

	// HeaderTenantCode
	HeaderTenantCode = `X-MM-Tenant-Code`

	// HeaderTenantHashId
	HeaderTenantHashId = `X-MM-Tenant-HashId`
)

var (
	HeaderMap = map[string]string{
		"X-Mm-Admin-Email":    HeaderAdminEmail,
		"X-Mm-Device-Details": HeaderDeviceDetails,
		"X-Mm-Tenant-Code":    HeaderTenantCode,
		"X-Mm-Tenant-Hashid":  HeaderTenantHashId,
	}
)

// GetInstance function ...
func GetInstance(c *container.Container) *tcp.ITCP {
	var h, p string

	if h, p = os.Getenv(constant.EnvConfigAppHost), os.Getenv(constant.EnvConfigAppPort); h == "" || p == "" {
		panic("Config host and/or port values are required but not declared")
	}

	i := tcp.New(c)
	(*i).SetHost(h)
	(*i).SetPort(p)
	(*i).SetHeaderVersion(ResourceHeaderVersion)
	(*i).SetResourceVersion(ResourceAPIVersion)

	return i
}

// PrepareHeaders ..  to correct the case to X-MM
func PrepareHeaders(headers http.Header) (pHeaders [20]payloadConfig.Header) {

	// prepare headers
	var (
		i               = 0
		supportedHeader string
		ok              bool
	)

	for h, v := range headers {

		// skip unsupported headers
		if supportedHeader, ok = HeaderMap[h]; !ok {
			continue
		}

		pHeaders[i] = payloadConfig.Header{
			Key:   supportedHeader,
			Value: strings.Join(v, ","), // header values are in string array
		}
		i++
	}

	return
}
