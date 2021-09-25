package tcp

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/go-commons/message"
	"bitbucket.org/matchmove/go-commons/service"
	"github.com/go-errors/errors"
)

type TCP struct {
	host            string
	port            string
	headerVersion   int
	resourceName    string
	resourceVersion string
	flag 	int
	payload         string
	rawRequest      string
	rawResponse     string
}

type (
	ITCP interface {
		IHostGS
		IPortGS
		IHeaderVersionGS
		IResourceNameGS
		IResourceVersionGS
		IFlagGS
		IPayloadGS
		IRawRequestGS
		IRawResponseGS
		IFormatter
		ICaller
	}

	IHostGS interface {
		GetHost() string
		SetHost(string)
	}

	IPortGS interface {
		GetPort() string
		SetPort(string)
	}

	IHeaderVersionGS interface {
		GetHeaderVersion() int
		SetHeaderVersion(int)
	}

	IResourceNameGS interface {
		GetResourceName() string
		SetResourceName(string)
	}

	IResourceVersionGS interface {
		GetResourceVersion() string
		SetResourceVersion(string)
	}

	IFlagGS interface {
		GetFlag() int
		SetFlag(int)
	}

	IPayloadGS interface {
		GetPayload() string
		SetPayload(string)
	}

	IRawRequestGS interface {
		GetRawRequest() string
		SetRawRequest(string)
	}

	IRawResponseGS interface {
		GetRawResponse() string
		SetRawResponse(string)
	}

	IFormatter interface {
		FormatRequest() error
		FormatResponse() error
	}

	ICaller interface {
		Call() error
	}
)

const (
	InstanceKey = "TCPIns"
)

// New function - prepares new tcp object
func New(c *container.Container) *ITCP {
	if ins := c.Get(InstanceKey); ins != nil {
		return ins.(*ITCP)
	}

	me := &TCP{}
	i := ITCP(me)

	return &i
}

// GetHost method ...
func (me *TCP) GetHost() string {
	return me.host
}

// SetHost method ...
func (me *TCP) SetHost(h string) {
	me.host = h
}

// GetPort method ...
func (me *TCP) GetPort() string {
	return me.port
}

// SetPort method ...
func (me *TCP) SetPort(p string) {
	me.port = p
}

// GetHeaderVersion method ...
func (me *TCP) GetHeaderVersion() int {
	return me.headerVersion
}

// SetHeaderVersion method ...
func (me *TCP) SetHeaderVersion(hv int) {
	me.headerVersion = hv
}

// GetResourceName method ...
func (me *TCP) GetResourceName() string {
	return me.resourceName
}

// SetResourceName method ...
func (me *TCP) SetResourceName(rn string) {
	me.resourceName = rn
}

// GetResourceVersion method ...
func (me *TCP) GetResourceVersion() string {
	return me.resourceVersion
}

// SetResourceVersion method ...
func (me *TCP) SetResourceVersion(rv string) {
	me.resourceVersion = rv
}

// GetFlag method ...
func (me *TCP) GetFlag() int {
	return me.flag
}

// SetFlag method ...
func (me *TCP) SetFlag(f int) {
	me.flag = f
}

// GetPayload method ...
func (me *TCP) GetPayload() string {
	return me.payload
}

// SetPayload method ...
func (me *TCP) SetPayload(p string) {
	me.payload = p
}

// GetRawRequest method ...
func (me *TCP) GetRawRequest() string {
	return me.rawRequest
}

// SetRawRequest method ...
func (me *TCP) SetRawRequest(rr string) {
	me.rawRequest = rr
}

// GetRawResponse method ...
func (me *TCP) GetRawResponse() string {
	return me.rawResponse
}

// SetRawResponse method ...
func (me *TCP) SetRawResponse(rr string) {
	me.rawResponse = rr
}

// FormatRequest method ...
func (me *TCP) FormatRequest() (err error) {
	// bypass formatting if rawRequest value already provided.
	if me.rawRequest != "" {
		return nil
	}

	var m = message.New()

	defer func() {
		if r := recover(); err != nil {
			err = errors.New(r)
		}
	}()

	me.rawRequest = string(m.Serialize(me.headerVersion, me.resourceName, me.resourceVersion, me.payload))

	return nil
}

// FormatResponse method ...
func (me *TCP) FormatResponse() (err error) {
	var m = message.New()

	defer func() {
		if r := recover(); err != nil {
			err = errors.New(r)
		}
	}()

	m.Unserialize([]byte(me.rawResponse))
	me.headerVersion = m.HeaderVersion
	me.resourceName = m.Action
	me.resourceVersion = m.Version
	me.flag = m.Flags
	me.payload = m.Payload

	return nil
}

// Call method ...
func (me *TCP) Call() (err error) {
	var serv *service.Service

	me.rawResponse, err = serv.Consume(me.host, me.port, me.resourceName, []byte(me.rawRequest))

	return err
}

