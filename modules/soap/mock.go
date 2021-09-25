package soap

import (
	"net/http"
)

type Mock struct {
	Soap
	Error               error
	Response            string
	success             bool
	consecutiveResponse []string
	StatusCode          int
}

// SuccessScenario mocker method ...
func (me *Mock) SuccessScenario(b bool) {
	me.success = b
}

// SetMockResponse mocker method ...
func (me *Mock) SetMockResponse(r string) {
	me.Response = r
}

// SetMultipleMockResponse mocker method ...
func (me *Mock) SetMultipleMockResponse(r string) {
	me.consecutiveResponse = append(me.consecutiveResponse, r)
}

// AddHeader method - add new header key value pair
func (me *Mock) AddHeader(key string, value string) *ICall {
	i := ICall(me)
	return &i
}

// Call method - triggers the soap call
func (me *Mock) Call() (response *http.Response, rawBody []byte, err error) {
	response = &http.Response{StatusCode: me.StatusCode}

	err = me.Error

	if me.success {
		if len(me.consecutiveResponse) == 0 {
			rawBody = []byte(me.Response)
		} else {
			me.Response = UnshiftSlice(&me.consecutiveResponse)
			rawBody = []byte(me.Response)
		}

	}

	return
}

// SetClient method ...
func (me *Mock) SetClient(client *http.Client) *ICall {
	i := ICall(me)
	return &i
}

// SetMethod method ...
func (me *Mock) SetMethod(method string) *ICall {
	i := ICall(me)
	return &i
}

// SetUrl method ...
func (me *Mock) SetUrl(url string) *ICall {
	i := ICall(me)
	return &i
}

// SetAuthentication ...
func (me *Mock) SetAuthentication(authType string, args ...interface{}) *ICall {
	i := ICall(me)
	return &i
}

// SetPayload method ...
func (me *Mock) SetPayload(payload interface{}) *ICall {
	i := ICall(me)
	return &i
}

// StructToPayload convert struct to paylaod
func (me *Mock) StructToPayload(params map[string]interface{}) *ICall {
	i := ICall(me)
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
