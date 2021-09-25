package getTenantRole

import (
	"net/http"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/tcp/config"
	payloadConfig "bitbucket.org/matchmove/go-payloads/config"
	"github.com/go-errors/errors"
)

const (
	ResourceName = "GetUserRole"
)

// Call function ...
func Call(c *container.Container, payload string) (string, error) {
	i := config.GetInstance(c)
	(*i).SetResourceName(ResourceName)
	(*i).SetPayload(payload)
	(*i).FormatRequest()
	err := (*i).Call()

	if err != nil {
		return "", err
	}

	(*i).FormatResponse()

	return (*i).GetPayload(), nil
}

// FormatRequest function ...
func FormatRequest(req payloadConfig.TenantUserRoleRequest , headers http.Header) string {
	r := payloadConfig.TenantUserRoleRequest{
		UserHashID:   req.UserHashID,
		TenantHashID: req.TenantHashID,
		Headers:      config.PrepareHeaders(headers),
	}

	return r.Serialize()
}

// FormatResponse function ...
func FormatResponse(r string) (fr payloadConfig.UserRole, err error) {

	defer func() {
		if r := recover(); err != nil {
			err = errors.New(r)
		}
	}()
	fr.Unserialize(r)
	return fr,err
}

