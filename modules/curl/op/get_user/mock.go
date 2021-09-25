package get_user

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"errors"
	"net/http"
)

type MockCurl struct {
	curl.Curl
}

// AddHeader method - add new header key value pair
func (me *MockCurl) AddHeader(key string, value string) *curl.ICall {
	i := curl.ICall(me)
	return &i
}

// Call method - triggers the curl call
func (me *MockCurl) Call() (response *http.Response, rawBody []byte, err error) {
	response = &http.Response{StatusCode: 500}
	rawBody = []byte("{}")
	err = errors.New("Internal Service Error!")

	return
}

// SetClient method ...
func (me *MockCurl) SetClient(client *http.Client) *curl.ICall {
	i := curl.ICall(me)
	return &i
}

// SetMethod method ...
func (me *MockCurl) SetMethod(method string) *curl.ICall {
	i := curl.ICall(me)
	return &i
}

// SetUrl method ...
func (me *MockCurl) SetUrl(url string) *curl.ICall {
	i := curl.ICall(me)
	return &i
}

// SetEndpoint method ...
func (me *MockCurl) SetEndpoint(endpoint string) *curl.ICall {
	i := curl.ICall(me)
	return &i
}

// SetAuthentication ...
func (me *MockCurl) SetAuthentication(authType string, args ...interface{}) *curl.ICall {
	i := curl.ICall(me)
	return &i
}

// SetPayload method ...
func (me *MockCurl) SetPayload(payload string) *curl.ICall {
	i := curl.ICall(me)
	return &i
}

// StructToPayload convert struct to paylaod
func (me *MockCurl) StructToPayload(params map[string]interface{}) *curl.ICall {
	i := curl.ICall(me)
	return &i
}
