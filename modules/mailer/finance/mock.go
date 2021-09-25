package finance

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/mailer"
)

type MockMailer struct {
	mailer.Mailer
}

// GetRegistry method ...
func GetMockRegistry() container.Registries {
	i := mailer.ISend(&MockMailer{})
	return container.Registries {
		container.Registry{
			Key: mailer.InstanceKey,
			Value: &i,
		},
	}
}

// AddHeader method ...
func (me *MockMailer) AddHeader(field string, value ...string) *mailer.ISend {
	i := mailer.ISend(me)
	return &i
}

// AddAddressHeader method ...
func (me *MockMailer) AddAddressHeader(field, address, name string) *mailer.ISend {
	i := mailer.ISend(me)
	return &i
}

// AddRecipient method ...
func (me *MockMailer) AddRecipient(emailAddresses ...string) *mailer.ISend {
	i := mailer.ISend(me)
	return &i
}

// AddCC method ...
func (me *MockMailer) AddCC(address, name string) *mailer.ISend {
	i := mailer.ISend(me)
	return &i
}

// AddBCC method ...
func (me *MockMailer) AddBCC(address, name string) *mailer.ISend {
	i := mailer.ISend(me)
	return &i
}

// SetSubject method ...
func (me *MockMailer) SetSubject(subj string) *mailer.ISend {
	i := mailer.ISend(me)
	return &i
}

// SetHTMLBody method ...
func (me *MockMailer) SetHTMLBody(body string) *mailer.ISend {
	i := mailer.ISend(me)
	return &i
}

// SetPlainBody method ...
func (me *MockMailer) SetPlainBody(body string) *mailer.ISend {
	i := mailer.ISend(me)
	return &i
}

// AttachFile method ...
func (me *MockMailer) AttachFile(filePath string) *mailer.ISend {
	i := mailer.ISend(me)
	return &i
}

// Send method ...
func (me *MockMailer) Send() error {
	return nil
}
