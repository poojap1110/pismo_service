package getTenantPermission

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/logger"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/tcp/config"
	payloadConfig "bitbucket.org/matchmove/go-payloads/config"
	"github.com/go-errors/errors"
)

const (
	ResourceName = "GetTenantPermission"
	LogName      = config.LogName + "-" + ResourceName
)

// Call function ...
func Call(c *container.Container, payload string) (string, error) {
	i := config.GetInstance(c)
	(*i).SetResourceName(ResourceName)
	(*i).SetPayload(payload)
	(*i).FormatRequest()

	if err := (*i).Call(); err != nil {
		return "", err
	} else if err := (*i).FormatResponse(); err != nil {
		return "", err
	}

	go logger.GetInstance(c, LogName).WriteLog([]byte("Request:\n" + (*i).GetRawRequest() + "\nResponse:\n" + (*i).GetRawResponse()))

	return (*i).GetPayload(), nil
}

// FormatRequestPayload function ...
func FormatRequestPayload(tenantHashId string) string {
	r := payloadConfig.GetTenantPermissionRequest{
		ID:           "",
		TenantHashID: tenantHashId,
		AccessName:   "",
	}

	return r.Serialize()
}

// FormatResponsePayload function ...
func FormatResponsePayload(r string) (fr payloadConfig.TenantPermissions, err error) {
	defer func() {
		if r := recover(); err != nil {
			err = errors.New(r)
		}
	}()

	fr.Unserialize(r)

	return fr, err
}
