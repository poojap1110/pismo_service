package process_transaction

import (
	"net/http"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/soap"
)

type MockSoap struct {
	soap.Soap
	Error               error
	Response            string
	success             bool
	consecutiveResponse []string
	StatusCode          int
}

// SuccessScenario mocker method ...
func (me *MockSoap) SuccessScenario(b bool) {
	me.success = b
}

// SetMockResponse mocker method ...
func (me *MockSoap) SetMockResponse(r string) {
	me.Response = r
}

// SetMultipleMockResponse mocker method ...
func (me *MockSoap) SetMultipleMockResponse(r string) {
	me.consecutiveResponse = append(me.consecutiveResponse, r)
}

// AddHeader method ...
func (me *MockSoap) AddHeader(k string, v string) *soap.ICall {
	i := soap.ICall(me)
	return &i
}

// Call method ...
func (me *MockSoap) Call() (response *http.Response, body []byte, err error) {
	response = &http.Response{StatusCode: me.StatusCode}

	err = me.Error

	if me.success {
		if len(me.consecutiveResponse) == 0 {
			body = []byte(me.Response)
		} else {
			me.Response = UnshiftSlice(&me.consecutiveResponse)
			body = []byte(me.Response)
		}
	}
	return
}

// SetClient method ...
func (me *MockSoap) SetClient(*http.Client) *soap.ICall {
	i := soap.ICall(me)
	return &i
}

// SetMethod method ...
func (me *MockSoap) SetMethod(m string) *soap.ICall {
	i := soap.ICall(me)
	return &i
}

// SetUrl ...
func (me *MockSoap) SetUrl(u string) *soap.ICall {
	i := soap.ICall(me)
	return &i
}

// SetAuthentication ...
func (me *MockSoap) SetAuthentication(string, ...interface{}) *soap.ICall {
	i := soap.ICall(me)
	return &i
}

// SetPayload ...
func (me *MockSoap) SetPayload(p interface{}) *soap.ICall {
	i := soap.ICall(me)
	return &i
}

// UnshiftSlice ...
func UnshiftSlice(s *[]string) (e string) {
	if len(*s) == 0 {
		return ""
	}
	e = (*s)[0] // get the first element
	if len(*s) > 1 {
		*s = (*s)[1:] // get the first element if there are more than 1 element
	}
	return
}
