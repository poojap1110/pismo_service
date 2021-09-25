package tcp

import (
	"github.com/go-errors/errors"
)

type Mock struct {
	TCP
	success bool
	response string
	consecutiveResponse []string
}

// SuccessScenario mocker method ...
func (me *Mock) SuccessScenario(b bool) {
	me.success = b
}

// SetMockResponse mocker method ...
func (me *Mock) SetMockResponse(r string) {
	me.response = r
}

// SetMockResponse mocker method ...
func (me *Mock) SetMultipleMockResponse(r string) {
	me.consecutiveResponse = append(me.consecutiveResponse, r)
}

// GetHost mocked method ...
func (me *Mock) GetHost() string {
	return ""
}

// SetHost mocked method ...
func (me *Mock) SetHost(h string) {

}

// GetPort mocked method ...
func (me *Mock) GetPort() string {
	return ""
}

// SetPort mocked method ...
func (me *Mock) SetPort(p string) {

}

// GetHeaderVersion mocked method ...
func (me *Mock) GetHeaderVersion() int {
	return 0
}

// SetHeaderVersion mocked method ...
func (me *Mock) SetHeaderVersion(hv int) {

}

// GetResourceName mocked method ...
func (me *Mock) GetResourceName() string {
	return ""
}

// SetResourceName mocked method ...
func (me *Mock) SetResourceName(rn string) {

}

// GetResourceVersion mocked method ...
func (me *Mock) GetResourceVersion() string {
	return ""
}

// SetResourceVersion mocked method ...
func (me *Mock) SetResourceVersion(rv string) {

}

// GetFlag method ...
func (me *Mock) GetFlag() int {
	return me.TCP.GetFlag()
}

// SetFlag method ...
func (me *Mock) SetFlag(f int) {

}

// GetPayload mocked method ...
func (me *Mock) GetPayload() string {
	return me.TCP.GetPayload()
}

// SetPayload mocked method ...
func (me *Mock) SetPayload(p string) {

}

// GetRawRequest mocked method ...
func (me *Mock) GetRawRequest() string {
	return ""
}

// SetRawRequest mocked method ...
func (me *Mock) SetRawRequest(rr string) {

}

// GetRawResponse mocked method ...
func (me *Mock) GetRawResponse() string {
	if me.response != "" {
		return me.response
	}

	return `{}`
}

// SetRawResponse mocked method ...
func (me *Mock) SetRawResponse(rr string) {

}

// FormatRequest mocked method ...
func (me *Mock) FormatRequest() (err error) {
	return nil
}

// FormatResponse mocked method ...
func (me *Mock) FormatResponse() (err error) {
	return me.TCP.FormatResponse()
}

// Call mocked method ...
func (me *Mock) Call() (err error) {
	if me.success {
		if len(me.consecutiveResponse) == 0 {
			me.TCP.rawResponse = me.response
			return nil
		} else {
			me.TCP.rawResponse = UnshiftSlice(&me.consecutiveResponse)
			return nil
		}

	}

	return errors.New("Mock_TCP: Error")
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

