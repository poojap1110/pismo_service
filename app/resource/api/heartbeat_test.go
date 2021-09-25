package api_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"bitbucket.org/matchmove/integration-svc-aub/app/resource/api"
	"github.com/stretchr/testify/assert"
)

func TestAPI_Heartbeat_Get(t *testing.T) {
	mysqlConn, cont := GetEnvironmentDB()
	rName, rRoute, ci := api.NewHeartbeat(mysqlConn, cont)
	c := ci().(*api.Heartbeat)
	c.Vars = make(map[string]string)

	// success scenario
	assert.Equal(t, "heartbeat", rName)
	assert.Equal(t, "/api/heartbeat", rRoute)

	resp := mockHeartbeat("GET", "", "", c, func(c *api.Heartbeat) {
		c.Get()
	})

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, resp)
}

// mockHeartbeat ...
func mockHeartbeat(m string, product string, q string, c *api.Heartbeat, f func(c *api.Heartbeat)) *http.Response {
	var body io.Reader
	url := "localhost:8082/api/heartbeat"

	if m == http.MethodGet {
		if q != "" {
			url = url + "?" + q
		}
	} else {
		body = bytes.NewBufferString(q)
	}

	c.Request = httptest.NewRequest(m, url, body)
	w := httptest.NewRecorder()
	c.Response = w

	func() {
		c.Init()
		defer c.Defer()
		f(c)
		c.Done()
	}()

	return w.Result()
}
