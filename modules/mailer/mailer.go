package mailer

import (
	"strconv"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/config"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"gopkg.in/gomail.v2"
)

const (
	HeaderFrom       = "From"
	HeaderTo         = "To"
	HeaderCc         = "Cc"
	HeaderBcc        = "Bcc"
	HeaderSubject    = "Subject"
	ContentTypeHtml  = "text/html"
	ContentTypePlain = "text/plain"
	InstanceKey      = "mailerIns"
)

type Mailer struct {
	gm       *gomail.Message
	server   string
	port     int
	username string
	password string
}

type ISend interface {
	AddHeader(field string, value ...string) *ISend
	AddAddressHeader(field, address, name string) *ISend
	AddRecipient(emailAddresses ...string) *ISend
	AddCC(address, name string) *ISend
	AddBCC(address, name string) *ISend
	SetSubject(subj string) *ISend
	SetHTMLBody(body string) *ISend
	SetPlainBody(body string) *ISend
	AttachFile(filePath string) *ISend
	Send() error
}

// New ...
func New(c *container.Container) *ISend {
	if ins := c.Get(InstanceKey); ins != nil { // Just to enable mocking.
		return ins.(*ISend)
	}

	mailerConfs := config.GetInstance(c).GetTechnicalConfigs(constant.TechnicalMailer).(map[string]interface{})
	msg := gomail.NewMessage()
	msg.SetHeader(HeaderFrom, "no-reply <"+mailerConfs["username"].(string)+">")
	port, _ := strconv.Atoi(mailerConfs["port"].(string))
	i := ISend(&Mailer{
		server:   mailerConfs["server"].(string),
		port:     port,
		username: mailerConfs["username"].(string),
		password: mailerConfs["password"].(string),
		gm:       msg,
	})

	return &i
}

// AddHeader method ...
func (me *Mailer) AddHeader(field string, value ...string) *ISend {
	me.gm.SetHeader(field, value...)
	i := ISend(me)
	return &i
}

// AddAddressHeader method ...
func (me *Mailer) AddAddressHeader(field, address, name string) *ISend {
	me.gm.SetAddressHeader(field, address, name)
	i := ISend(me)
	return &i
}

// AddRecipient method ...
func (me *Mailer) AddRecipient(emailAddresses ...string) *ISend {
	me.AddHeader(HeaderTo, emailAddresses...)
	i := ISend(me)
	return &i
}

// AddCC method ...
func (me *Mailer) AddCC(address, name string) *ISend {
	me.AddAddressHeader(HeaderCc, address, name)
	i := ISend(me)
	return &i
}

// AddBCC method ...
func (me *Mailer) AddBCC(address, name string) *ISend {
	me.AddAddressHeader(HeaderBcc, address, name)
	i := ISend(me)
	return &i
}

// SetSubject method ...
func (me *Mailer) SetSubject(subj string) *ISend {
	me.AddHeader(HeaderSubject, subj)
	i := ISend(me)
	return &i
}

// SetHTMLBody method ...
func (me *Mailer) SetHTMLBody(body string) *ISend {
	me.gm.SetBody(ContentTypeHtml, body)
	i := ISend(me)
	return &i
}

// SetPlainBody method ...
func (me *Mailer) SetPlainBody(body string) *ISend {
	me.gm.SetBody(ContentTypePlain, body)
	i := ISend(me)
	return &i
}

// AttachFile method ...
func (me *Mailer) AttachFile(filePath string) *ISend {
	me.gm.Attach(filePath)
	i := ISend(me)
	return &i
}

// Send method ...
func (me *Mailer) Send() error {
	return gomail.NewPlainDialer(me.server, me.port, me.username, me.password).DialAndSend(me.gm)
}
