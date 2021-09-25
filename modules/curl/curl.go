package curl

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
)

type Curl struct {
	m            sync.Mutex
	headers      map[string]string
	customheader map[string]string
	method       string
	url          string
	endpoint     string
	payload      string
	client       *http.Client
}

type ICall interface {
	AddHeader(string, string) *ICall
	AddCustomHeader(string, string) *ICall
	Call() (*http.Response, []byte, error)
	SetClient(*http.Client) *ICall
	SetMethod(string) *ICall
	SetUrl(string) *ICall
	SetEndpoint(string) *ICall
	SetAuthentication(string, ...interface{}) *ICall
	SetPayload(string) *ICall
	StructToPayload(params map[string]interface{}) *ICall
	ByteToPayload([]byte) *ICall
}

type Executor func(req *http.Request) (res *http.Response, err error)

const (
	InstanceKey           = "CurlIns"
	RequestTimeoutSeconds = 30

	// Headers
	Accept        = "Accept"
	Authorization = "Authorization"
	ContentType   = "Content-Type"

	// Content Types
	ApplicationAll            = "*/*"
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

	me := &Curl{
		m: sync.Mutex{},
	}

	i := ICall(me)
	return &i
}

// AddHeader method - add new header key value pair
func (me *Curl) AddHeader(key string, value string) *ICall {
	if me.headers == nil {
		me.headers = make(map[string]string)
	}

	me.m.Lock()
	me.headers[key] = value
	me.m.Unlock()
	i := ICall(me)

	return &i
}

// AddHeader method - add new header key value pair
func (me *Curl) AddCustomHeader(key string, value string) *ICall {
	if me.customheader == nil {
		me.customheader = make(map[string]string)
	}

	me.m.Lock()
	me.customheader[key] = value
	me.m.Unlock()
	i := ICall(me)

	return &i
}

// Call method - triggers the curl call
func (me *Curl) Call() (response *http.Response, rawBody []byte, err error) {
	if me.method == "" || me.url == "" {
		panic("method, url, and endpoint are required values")
	}
	readPayload := strings.NewReader(me.payload)
	fullEndpoint := strings.TrimRight(me.url, "/") + "/" + strings.TrimLeft(me.endpoint, "/")
	fullEndpoint = strings.TrimRight(fullEndpoint, "/")
	request, err := http.NewRequest(me.method, fullEndpoint, readPayload)

	if err != nil {
		return
	}

	for k, v := range me.headers {
		request.Header.Add(k, v)
	}

	for k, v := range me.customheader {
		request.Header[k] = []string{v}
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
func (me *Curl) SetClient(client *http.Client) *ICall {
	me.client = client
	i := ICall(me)
	return &i
}

// SetMethod method ...
func (me *Curl) SetMethod(method string) *ICall {
	me.method = method
	i := ICall(me)
	return &i
}

// SetUrl method ...
func (me *Curl) SetUrl(url string) *ICall {
	me.url = url
	i := ICall(me)
	return &i
}

// SetEndpoint method ...
func (me *Curl) SetEndpoint(endpoint string) *ICall {
	me.endpoint = endpoint
	i := ICall(me)
	return &i
}

// SetAuthentication ...
func (me *Curl) SetAuthentication(authType string, args ...interface{}) *ICall {
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

// SetPayload method ...
func (me *Curl) SetPayload(payload string) *ICall {
	me.payload = payload
	i := ICall(me)
	return &i
}

// setDefaultClient method ...
func (me *Curl) setDefaultClient() {
	me.SetClient(&http.Client{
		Timeout: time.Duration(RequestTimeoutSeconds) * time.Second,
	})
}

// StructToPayload convert struct to paylaod
func (me *Curl) StructToPayload(params map[string]interface{}) *ICall {

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
func (me *Curl) ByteToPayload(s []byte) *ICall {

	me.SetPayload(string(s))
	i := ICall(me)
	return &i
}
