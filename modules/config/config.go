package config

import (
	"os"
	"sync"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/jmoiron/sqlx"
)

type config struct {
	sync.Mutex
	technicalConfigs       map[string]interface{}
	globalTechnicalConfigs map[string]interface{}
	tenantConfigs          map[string]interface{}
}

type ConfigItem struct {
	Name         string
	Value        string
	DefaultValue string
	FriendlyName string
}

type ConfigFilter struct {
	Owner        string
	Name         string
	TenantHashID string
	Configs      []ConfigItem
}

// Constants for config service name value key pair
const (
	// InstanceKey ...
	InstanceKey = "Configs" // Global Container's config instance

	// LogName ...
	LogName = "CONFIG"

	// MockConfigEnable ...
	MockConfigEnable = "1"
)

// GetRegistry function ...
func GetRegistry() container.Registries {
	configObject := &config{
		technicalConfigs:       make(map[string]interface{}),
		globalTechnicalConfigs: make(map[string]interface{}),
		tenantConfigs:          make(map[string]interface{}),
	}

	return container.Registries{
		container.Registry{
			Key:   InstanceKey,
			Value: configObject,
		},
	}
}

// ContainerGetRegistry function - creates new instance for tenant_config & tenant_info for new container.
func ContainerGetRegistry(c *container.Container) container.Registries {
	return GetInstance(c).ContainerGetRegistry()
}

// ContainerGetRegistry method - creates new instance for tenant_config & tenant_info for new container.
func (me *config) ContainerGetRegistry() container.Registries {
	configObject := config{
		technicalConfigs:       me.technicalConfigs,
		globalTechnicalConfigs: me.globalTechnicalConfigs,
		tenantConfigs:          make(map[string]interface{}),
	}

	// inherit tenant configs and place to a new memory location
	for k, v := range me.tenantConfigs {
		configObject.tenantConfigs[k] = v
	}

	return container.Registries{
		container.Registry{
			Key:   InstanceKey,
			Value: &configObject,
		},
	}
}

// GetInstance function ...
func GetInstance(c *container.Container) *config {
	return c.Get(InstanceKey).(*config)
}

// GetTechnicalConfigs method - get technical configs
func (me *config) GetTechnicalConfigs(key string) (value interface{}) {
	if me.technicalConfigs[key] == nil {
		panic(key + " configs are required but not found!")
	}

	return me.technicalConfigs[key]
}

// SetTechnicalConfigs method - set technical config
func (me *config) SetTechnicalConfigs(key string, val interface{}) {
	defer func() {
		me.Unlock()
	}()
	me.Lock()
	me.technicalConfigs[key] = val
}

// GetTenantConfig method - get technical config
func (me *config) GetTenantConfig(key string) (value interface{}) {
	if me.tenantConfigs[key] == nil {
		panic(key + " configs are required but not found!")
	}
	return me.tenantConfigs[key]
}

// MockConfigs method - populates mock data to config map.
func (me *config) MockConfigs() {
	me.technicalConfigs = map[string]interface{}{
		constant.TechnicalAWS: map[string]string{
			"key":    os.Getenv("AWS_KEY"),
			"secret": os.Getenv("AWS_SECRET"),
			"region": os.Getenv("AWS_REGION"),
			"bucket": os.Getenv("AWS_BUCKET"),
			"qurl":   os.Getenv("AWS_SQS_URL"),
		},
		constant.TechnicalMailer: map[string]interface{}{
			"server":   os.Getenv("MAILER_SERVER"),
			"port":     os.Getenv("MAILER_PORT"),
			"username": os.Getenv("MAILER_USERNAME"),
			"password": os.Getenv("MAILER_PASSWORD"),
		},
		constant.TechnicalFinanceMail: map[string]string{
			"recipients":    os.Getenv("FINANCE_MAIL_RECIPIENTS"),
			"ccRecipients":  os.Getenv("FINANCE_MAIL_CC_RECIPIENTS"),
			"bccRecipients": os.Getenv("FINANCE_MAIL_BCC_RECIPIENTS"),
			"subject":       os.Getenv("FINANCE_MAIL_SUBJECT"),
			"htmlBody":      os.Getenv("FINANCE_MAIL_BODY"),
		},
		constant.TechnicalIndusind: map[string]string{
			"client_id":      os.Getenv("INDUSIND_CLIENT_ID"),
			"client_secret":  os.Getenv("INDUSIND_CLIENT_SECRET"),
			"url":            os.Getenv("INDUSIND_URL"),
			"mm_account_no":  os.Getenv("MM_ACCOUNT_NO"),
			"mm_customer_id": os.Getenv("MM_CUSTOMER_ID"),
		},
	}

	me.tenantConfigs = map[string]interface{}{
		constant.TenantTimezone:          os.Getenv(constant.TenantTimezone),
		constant.TenantADUrl:             os.Getenv(constant.TenantADUrl),
		constant.TenantADUsername:        os.Getenv(constant.TenantADUsername),
		constant.TenantADPassword:        os.Getenv(constant.TenantADPassword),
		constant.TenantADClientSecret:    os.Getenv(constant.TenantADClientSecret),
		constant.TenantADClientID:        os.Getenv(constant.TenantADClientID),
		constant.TenantADConsumerID:      os.Getenv(constant.TenantADConsumerID),
		constant.TenantADGrantType:       os.Getenv(constant.TenantADGrantType),
		constant.TenantADApplicationType: os.Getenv(constant.TenantADApplicationType),
		constant.TenantVAUsername:        os.Getenv(constant.TenantVAUsername),
		constant.TenantVAPassword:        os.Getenv(constant.TenantVAPassword),
		constant.TenantInstitutionCode:   os.Getenv(constant.TenantInstitutionCode),
		constant.TenantInfoName:          os.Getenv(constant.TenantInfoName),
		constant.TenantProductCode:       os.Getenv(constant.TenantProductCode),
		constant.TenantCallbackEnabled:   os.Getenv(constant.TenantCallbackEnabled),
		constant.TenantCallbackDetails:   os.Getenv(constant.TenantCallbackDetails),
		constant.TenantADCode:            os.Getenv(constant.TenantADCode),
	}
}

// Populate method ...
func Populate(c *container.Container) {
	ins := GetInstance(c)

	if os.Getenv(constant.EnvMockConfigs) == MockConfigEnable {
		ins.MockConfigs()
	} else {
		ins.RetrieveTechnicalConfigs(c)
	}

	c.StoreToGlobal(InstanceKey, ins)
}

// RetrieveTechnicalConfigs method - populates config map with technical config values.
func (me *config) RetrieveTechnicalConfigs(c *container.Container) {
	me.technicalConfigs = map[string]interface{}{
		constant.TechnicalAWS: map[string]string{
			"key":    os.Getenv("AWS_KEY"),
			"secret": os.Getenv("AWS_SECRET"),
			"region": os.Getenv("AWS_REGION"),
			"bucket": os.Getenv("AWS_BUCKET"),
			"qurl":   os.Getenv("AWS_SQS_URL"),
		},
		constant.TechnicalMailer: map[string]interface{}{
			"server":   os.Getenv("MAILER_SERVER"),
			"port":     os.Getenv("MAILER_PORT"),
			"username": os.Getenv("MAILER_USERNAME"),
			"password": os.Getenv("MAILER_PASSWORD"),
		},
		constant.TechnicalFinanceMail: map[string]string{
			"recipients":    os.Getenv("FINANCE_MAIL_RECIPIENTS"),
			"ccRecipients":  os.Getenv("FINANCE_MAIL_CC_RECIPIENTS"),
			"bccRecipients": os.Getenv("FINANCE_MAIL_BCC_RECIPIENTS"),
			"subject":       os.Getenv("FINANCE_MAIL_SUBJECT"),
			"htmlBody":      os.Getenv("FINANCE_MAIL_BODY"),
		},
		constant.TechnicalIndusind: map[string]string{
			"client_id":     os.Getenv("INDUSIND_CLIENT_ID"),
			"client_secret": os.Getenv("INDUSIND_CLIENT_SECRET"),
			"url":           os.Getenv("INDUSIND_URL"),
			"mm_account_no":  os.Getenv("MM_ACCOUNT_NO"),
			"mm_customer_id": os.Getenv("MM_CUSTOMER_ID"),
		},
	}
}

// RetrieveTenantConfigs method - populates config map with tenant-specific values.
func (me *config) RetrieveTenantConfigs(c *container.Container, tenantHashID string, db *sqlx.DB, cache *memcache.Client) {
	if os.Getenv("MOCK_CONFIGS") == MockConfigEnable {
		return
	}

	var (
		err    error
		filter = ConfigFilter{
			TenantHashID: tenantHashID,
		}
	)

	if filter, err = me.callGetConfig(c, filter, constant.TenantTypeConfig, db, cache); err != nil {
		panic("Failed to fetch config values: " + err.Error())
	}

	me.ResetTenantConfigs()

	for _, v := range filter.Configs {
		me.SetTenantConfig(v.Name, getValue(v))
	}
}

// SetTenantConfig method ...
func (me *config) SetTenantConfig(key string, value interface{}) {
	me.Lock()
	me.tenantConfigs[key] = value
	me.Unlock()
}

//ResetTenantConfigs ...
func (me *config) ResetTenantConfigs() {
	me.Lock()
	me.tenantConfigs = nil
	me.tenantConfigs = make(map[string]interface{})
	me.Unlock()
}

// getValue function ...
func getValue(mC ConfigItem) (val string) {
	if mC.Value != "" {
		val = mC.Value
	} else {
		val = mC.DefaultValue
	}

	return
}
