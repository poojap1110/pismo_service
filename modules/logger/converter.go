package logger

import "encoding/json"

type HTTPReq struct {
	Method string `json:"method"`
	Endpoint string `json:"endpoint"`
	Payload string `json:"payload"`
	Headers []Header `json:"headers"`
}

type Header struct {
	Name string `json:"name"`
	Value interface{} `json:"value"`
}

// ConvertToJson method ...
func (me *HTTPReq) ConvertToJson() string {
	j, _ := json.Marshal(me)

	return string(j)
}
