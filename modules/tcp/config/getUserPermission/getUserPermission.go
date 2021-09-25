package getUserPermission

import (
	"strings"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/tcp/config"
	payloadConfig "bitbucket.org/matchmove/go-payloads/config"
)

const (
	ResourceName = "GetUserPermissions"
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
func FormatRequest(tenantHashId string, userHashId string, adminEmail string, tenantCode string, deviceDetails string, route string) string {
	var (
		header      = [20]payloadConfig.Header{}
	)

	r := payloadConfig.TenantUserRoleRequest{
		TenantHashID: tenantHashId,
		UserHashID:   userHashId,
	}

	header[0] = payloadConfig.Header{
		Key:   "X-MM-Admin-Email",
		Value: adminEmail,
	}
	header[1] = payloadConfig.Header{
		Key:   "X-MM-Tenant-Code",
		Value: tenantCode,
	}
	header[2] = payloadConfig.Header{
		Key:   "X-MM-Tenant-HashId",
		Value: tenantHashId,
	}
	header[3] = payloadConfig.Header{
		Key:   "X-MM-Device-Details",
		Value: deviceDetails,
	}
	r.Headers = header
	return r.Serialize()
}

// FormatResponse function ...
func FormatResponse(r string) (fr payloadConfig.UserPermissions) {
	fr.Unserialize(r)
	return fr
}

// CheckUserPemrission...
func CheckUserPemrission(roles payloadConfig.UserPermissions , url string) (authStatus int, err error){
	var validAccess =  make(map[string]string)
	validAccess["heartbeat"] = "GET /heartbeat"
	validAccess["authorize"] = "POST /authorize"
	authStatus = 0

	/**
	 *  check ep like login, forgot-password, heartbeat etc...
	 */
	for _, routeName := range validAccess {
		if strings.ToLower(routeName) == strings.ToLower(url) {
			authStatus = 1
			return
		}
	}

	/**
	 *	check ep like suspend-card, close-card, account-closure etc..
	 */
	for _, v := range roles.TenantRole.Permissions {
		if authStatus == 1 || strings.ToLower(url) == strings.ToLower(v.Value) {
			authStatus = 1
			break
		}
	}

	return
}