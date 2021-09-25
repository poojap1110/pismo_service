package soap

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
)

type Soap struct {
	m       sync.Mutex
	headers map[string]string
	method  string
	url     string
	payload interface{}
	client  *http.Client
}

type ICall interface {
	AddHeader(string, string) *ICall
	Call() (*http.Response, []byte, error)
	SetClient(*http.Client) *ICall
	SetMethod(string) *ICall
	SetUrl(string) *ICall
	SetPayload(interface{}) *ICall
	SetAuthentication(string, ...interface{}) *ICall
	StructToPayload(params map[string]interface{}) *ICall
	ByteToPayload([]byte) *ICall
}

const (
	InstanceKey           = "SoapIns"
	RequestTimeoutSeconds = 30

	// Headers
	Accept        = "Accept"
	Authorization = "Authorization"
	ContentType   = "Content-Type"
	SoapAction    = "SOAPAction"

	// Content Types
	ApplicationAll            = "*/*"
	TextXML                   = "text/xml; charset=utf-8"
	ApplicationJson           = "application/json"
	ApplicationFormUrlEncoded = "application/x-www-form-urlencoded"

	// Authentication Types
	BasicAuth  = "Basic"
	BearerAuth = "Bearer"
)

// GetRegistry method ...
func GetRegistry() container.Registries {
	return container.Registries{
		container.Registry{
			Key:   InstanceKey,
			Value: nil,
		},
	}
}

// New method - prepares new curl object
func New(c *container.Container) *ICall {
	if ins := c.Get(InstanceKey); ins != nil {
		return ins.(*ICall)
	}

	me := &Soap{
		m: sync.Mutex{},
	}

	i := ICall(me)
	return &i
}

// AddHeader method - add new header key value pair
func (me *Soap) AddHeader(key string, value string) *ICall {
	if me.headers == nil {
		me.headers = make(map[string]string)
	}

	me.m.Lock()
	me.headers[key] = value
	me.m.Unlock()
	i := ICall(me)

	return &i
}

// Call method - triggers the curl call
func (me *Soap) Call() (response *http.Response, rawBody []byte, err error) {

	if me.method == "" || me.url == "" {
		panic("method and url are required values")
	}

	payload, err := xml.MarshalIndent(me.payload, "", "  ")
	if err != nil {
		return
	}

	url := strings.TrimRight(me.url, "/")

	request, err := http.NewRequest(me.method, url, bytes.NewBuffer(payload))
	if err != nil {
		return
	}

	for k, v := range me.headers {
		request.Header.Add(k, v)
	}

	if me.client == nil {
		me.setDefaultClient()
	}

	response, err = me.client.Do(request)

	if err != nil {
		return
	}

	defer response.Body.Close()

	rawBody, err = ioutil.ReadAll(response.Body)

	return
}

// SetClient method ...
func (me *Soap) SetClient(client *http.Client) *ICall {
	me.client = client
	i := ICall(me)
	return &i
}

// SetMethod method ...
func (me *Soap) SetMethod(method string) *ICall {
	me.method = method
	i := ICall(me)
	return &i
}

// SetUrl method ...
func (me *Soap) SetUrl(url string) *ICall {
	me.url = url
	i := ICall(me)
	return &i
}

// SetPayload method ...
func (me *Soap) SetPayload(payload interface{}) *ICall {
	me.payload = payload
	i := ICall(me)
	return &i
}

// setDefaultClient method ...
func (me *Soap) setDefaultClient() {
	me.SetClient(&http.Client{
		Timeout: time.Duration(RequestTimeoutSeconds) * time.Second,
	})
}

// StructToPayload convert struct to paylaod
func (me *Soap) StructToPayload(params map[string]interface{}) *ICall {

	values := url.Values{}
	for k, val := range params {
		if val.(string) != "" {
			values.Set(k, val.(string))
		}
	}

	me.SetPayload(values.Encode())

	i := ICall(me)
	return &i
}

// GetStatusCode function - handles timeout and returns dynamic status code basing from http.Response.
func GetStatusCode(hr *http.Response) int {
	statusCode := http.StatusGatewayTimeout

	if hr != nil {
		statusCode = hr.StatusCode
	}

	return statusCode
}

// ByteToPayload convert byte to paylaod
func (me *Soap) ByteToPayload(s []byte) *ICall {

	me.SetPayload(string(s))
	i := ICall(me)
	return &i
}

// SetAuthentication ...
func (me *Soap) SetAuthentication(authType string, args ...interface{}) *ICall {
	switch authType {
	case BasicAuth:
		if len(args) < 2 {
			panic("Pass `Username` and `Password` as string on args")
		}

		username := args[0].(string)
		password := args[1].(string)

		me.AddHeader(Authorization, BasicAuth+" "+base64.StdEncoding.EncodeToString([]byte(username+":"+password)))

	case BearerAuth:
		if len(args) < 1 {
			panic("Pass `access_token` as string on args")
		}

		accessToken := args[0].(string)

		me.AddHeader(Authorization, BearerAuth+" "+accessToken)
	}

	i := ICall(me)
	return &i
}
