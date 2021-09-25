package entity

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"bitbucket.org/matchmove/go-tools/array"
)

// payload const declaration
const (
	CT          = "Content-Type"
	CTJson      = "application/json"
	CTPlain     = "text/plain"
	CTFormData  = "multipart/form-data"
	CTUrlEncode = "application/x-www-form-urlencoded"
)

// Helper variable declaration
var (
	JSONSupportCT = []string{CTJson}
	FormSupportCT = []string{CTFormData, CTUrlEncode}
)

// GetContentType ...
func GetContentType(req *http.Request) (ct string) {
	ct = req.Header.Get(CT)
	if flg := strings.Contains(ct, ";"); flg {
		ctWithBountry := strings.Split(ct, ";")
		ct = ctWithBountry[0]
	}
	return
}

// ValidContentType ...
func ValidContentType(ct string) bool {
	if CheckJSONCT(ct) || CheckFormDataCT(ct) {
		return true
	}
	return false
}

// CheckJSONCT - check json related content type
func CheckJSONCT(ct string) bool {
	exist, _ := array.InArray(ct, JSONSupportCT)
	return exist
}

// CheckFormDataCT - check form data related content type
func CheckFormDataCT(ct string) bool {
	exist, _ := array.InArray(ct, FormSupportCT)
	return exist
}

// ParseForm ...
func ParseForm(ct string, r *http.Request, isResource bool) (res url.Values, err error) {

	if r.Body != nil {
		// read all bytes from content body and create new stream using it.
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		url := r.URL.String()
		var req2 *http.Request

		if isResource {
			if r.Method == http.MethodDelete && ct == CTUrlEncode {
				url = url + "?" + string(bodyBytes)
				// create new request for parsing the body
				req2, _ = http.NewRequest(r.Method, url, nil)
			} else {
				// create new request for parsing the body
				req2, _ = http.NewRequest(r.Method, url, bytes.NewReader(bodyBytes))
			}
		} else {
			// create new request for parsing the body
			req2, _ = http.NewRequest(r.Method, url, bytes.NewReader(bodyBytes))
		}

		req2.Header = r.Header

		if ct == CTUrlEncode {
			if err = req2.ParseForm(); err != nil {
				return
			}

			if isResource {
				if r.Method == http.MethodDelete && req2.PostForm != nil {
					res = req2.Form
				} else {
					res = req2.PostForm
				}
			} else {
				res = req2.Form
			}
		} else if ct == CTFormData {
			if err = req2.ParseMultipartForm(200000); err != nil {
				return
			}
			res = req2.PostForm
		}
	}

	return
}
