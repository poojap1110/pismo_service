package api_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"bitbucket.org/matchmove/integration-svc-aub/app/errs"
	"bitbucket.org/matchmove/integration-svc-aub/app/resource/api"

	"github.com/stretchr/testify/assert"
)

// mockErrors ...
func mockErrors(m string, product string, q string, c *api.PismoErrors, f func(c *api.PismoErrors)) *http.Response {
	var body io.Reader
	url := "localhost:8082/api/errors"

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

// mockErrors ...
func mockError(m string, product string, q string, c *api.PismoError, f func(c *api.PismoError)) *http.Response {
	var body io.Reader
	url := "localhost:8082/api/error"

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

func TestAPI_Get_Errors_Success(t *testing.T) {
	mysqlConn, cont := GetEnvironmentDB()
	rName, rRoute, ci := api.NewPismoErrors(mysqlConn, cont)
	c := ci().(*api.PismoErrors)
	c.Vars = make(map[string]string)

	// success scenario
	assert.Equal(t, "pismo-error", rName)
	assert.Equal(t, "/v1/errors", rRoute)

	resp := mockErrors("GET", "", "", c, func(c *api.PismoErrors) {
		c.Get()
	})

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, resp)
}

func TestAPI_Get_Error_Success(t *testing.T) {

	mysqlConn, cont := GetEnvironmentDB()
	rName, rRoute, ci := api.NewPismoError(mysqlConn, cont)
	c := ci().(*api.PismoError)
	assert.Equal(t, "pismo-error", rName)

	// Get all existing errors
	errorCodes := errs.GetErrors()

	// Loop each error code and fetch the error Struct
	for eCode := range errorCodes {
		c.Vars = map[string]string{"error_code": eCode}
		assert.Equal(t, "/v1/error/{error_code}", rRoute)

		resp := mockError("GET", "", "", c, func(c *api.PismoError) {
			c.Get()
		})

		b, _ := ioutil.ReadAll(resp.Body)
		body := string(b[:])

		var raw map[string]interface{}
		json.Unmarshal([]byte(body), &raw)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, eCode, raw["code"])
		assert.NotEmpty(t, resp)

		t.Logf("%s: '%s'", raw["code"], raw["message"])
	}
}
