package balance_enquiry

import (
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/curl"
	"net/http"
)

type MockCurl struct {
	curl.Curl
	Error               error
	Response            string
	success             bool
	consecutiveResponse []string
	StatusCode          int
}

// SuccessScenario mocker method ...
func (me *MockCurl) SuccessScenario(b bool) {
	me.success = b
}

// SetMockResponse mocker method ...
func (me *MockCurl) SetMockResponse(r string) {
	me.Response = r
}

// SetMultipleMockResponse mocker method ...
func (me *MockCurl) SetMultipleMockResponse(r string) {
	me.consecutiveResponse = append(me.consecutiveResponse, r)
}
// AddHeader method ...
func (me *MockCurl) AddHeader(k string, v string) *curl.ICall {
	i := curl.ICall(me)
	return &i
}

// Call method ...
func (me *MockCurl) Call() (response *http.Response, body []byte, err error) {

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
func (me *MockCurl) SetClient(*http.Client) *curl.ICall {
	i := curl.ICall(me)
	return &i
}

// SetMethod method ...
func (me *MockCurl) SetMethod(m string) *curl.ICall {
	me.Curl.SetMethod("OPTIONS")
	i := curl.ICall(me)
	return &i
}

// SetUrl ...
func (me *MockCurl) SetUrl(u string) *curl.ICall {
	i := curl.ICall(me)
	return &i
}

// SetEndpoint ...
func (me *MockCurl) SetEndpoint(p string) *curl.ICall {
	i := curl.ICall(me)
	return &i
}

// SetAuthentication ...
func (me *MockCurl) SetAuthentication(string, ...interface{}) *curl.ICall {
	i := curl.ICall(me)
	return &i
}

// SetPayload ...
func (me *MockCurl) SetPayload(p string) *curl.ICall {
	i := curl.ICall(me)
	return &i
}

// LogToStdout ...
func (me *MockCurl) LogToStdout(res string, url string) {

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