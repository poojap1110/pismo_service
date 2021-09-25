package resource

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"/home/pooja/git/pismo-service/app/errs"

	"bitbucket.org/matchmove/integration-svc-aub/modules/container"
	"bitbucket.org/matchmove/integration-svc-aub/modules/entity"

	"bitbucket.org/matchmove/integration-svc-aub/modules/helper"

	"github.com/gorilla/schema"

	logs "bitbucket.org/matchmove/fmt-logs"
	gorm "bitbucket.org/matchmove/go-model"
	"bitbucket.org/matchmove/go-resource/out"
	"bitbucket.org/matchmove/go-tools/secure"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/willf/pad"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/matchmove/rest.v2"
)

const (
	// TimestampFormat Timestamp Format
	TimestampFormat = "20060102150405"
	// ContentType defines Content-Type
	ContentType = "Content-Type"
	// ContentTypePEMKeys defines the content type for our chosen key format
	ContentTypePEMKeys = `application/pem-keys`
	// ContentTypeEncrypted denotes the content type for encrypted payloads
	ContentTypeEncrypted = `application/x-tls-encrypted`

	// ResourcePin ResourcePin
	ResourcePin = "*api.Pin"

	//MaxPadLeft ...
	MaxPadLeft = 15

	DocPath = "/api/documentation"
)

// Resource is an extension of rest.Resource with an addional DB implementation
// and outputs the appropriate header and status
type Resource struct {
	rest.Resource
	Name       string
	DB         *sqlx.DB
	Container  *container.Container
	Log        *logs.Log
	Identifier string
	Key        string
	Time       int64
	Status     int
	Public     bool
	Validate   *validator.Validate
	RawBody    interface{}
}

// New creates a new ResourceDB instance
func New(name string, public bool, db *sqlx.DB, cont *container.Container) Resource {
	return Resource{rest.Resource{}, name, db, cont, logs.New(), "", "", 0, 0, public, nil, nil}
}

// LoadConfig load json config
func LoadConfig(path string, out interface{}) error {
	buff, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	return json.Unmarshal(buff, out)
}

// FormatException ...
func (r *Resource) FormatException(resource interface{}, err error, errList ...error) {

	var (
		errorString      = err.Error()
		errorStrings     []string
		values           []interface{}
		dynamicParameter = false
		mErr             errs.Error
	)

	if len(errList) > 0 {
		for _, item := range errList {
			r.Record("Error", item)
		}
	}
	r.Record("Error", err)

	switch {
	case strings.Contains(errorString, "cannot unmarshal string into Go struct field"):
		err = errors.New(errs.ErrRequestBodyInvalid)
	case strings.Contains(errorString, "invalid character"):
		err = errors.New(errs.ErrRequestBodyInvalid)
	case strings.Contains(strings.ToLower(errorString), "timeout"):
		err = errors.New(errs.ErrGatewayTimeout)
	case strings.Contains(errorString, ": "):
		dynamicParameter = true
		errorStrings = strings.Split(errorString, ": ")
		errorString = errorStrings[0]
	}

	mErr, err = errs.GetErrorByCode(errorString)

	if err != nil {
		mErr, _ = errs.GetErrorByCode(errs.ErrCodeNotFound)
	}

	// dynamically replace `%s`.
	if dynamicParameter && strings.Contains(mErr.Message, "%s") {
		if strings.Count(mErr.Message, "%s") > 1 {
			for i := 1; i <= strings.Count(mErr.Message, "%s"); i++ {
				values = append(values, errorStrings[i])
			}
		} else {
			e := strings.Join(errorStrings[1:], " ")
			values = append(values, e)
		}
	}

	r.Status = mErr.HTTPCode
	r.RawBody = errs.FormateErrorResponse(mErr, values...)

}

// ParentInit allows the generic Init() to be extended
func (r *Resource) ParentInit() bool {
	r.RawBody = nil
	r.Status = 0
	return r.AccessVerification(r.Route.Pattern)
}

// AccessVerification ...
func (r *Resource) AccessVerification(path string) bool {

	return true
}

// Init implements the initialize method of ReST resource
func (r *Resource) Init() bool {
	var (
		group string
	)

	r.Log = logs.New()

	r.Identify()
	r.Validate = validator.New()

	ip, _ := r.GetClientIP()

	r.Record("Start", time.Now().Format(gorm.SQLDatetime))
	r.Record("IP", ip)
	r.Record("Key", r.Key)

	if group = os.Getenv("DB_GROUP"); group == "" {
		group = "default"
	}

	r.Record("Resource", r.Name)
	r.Record("Method", r.Request.Method)
	r.Record("URL", r.Request.URL.String())
	r.Record("Param", r.Vars)

	var body = strings.Replace(string(r.Body()), "\n", "", -1)
	body = helper.MaskSensitive(body)
	r.Record("Request", body)

	return r.ParentInit()
}

// Record logs the entry on io and new relic
func (r *Resource) Record(key string, value interface{}) {
	r.Log.Print(pad.Right(key, MaxPadLeft, " "), value)
}

// Identify ...
func (r *Resource) Identify() {
	if r.Identifier == "" {
		key, secret, _ := r.Request.BasicAuth()
		r.Identifier = fmt.Sprintf("%s\t", secure.MD5(time.Now().UTC().Format(gorm.SQLDatetime)+
			key+
			secret+
			r.Request.URL.String()+
			r.Request.Method+
			string(r.Body())))
	}

	r.Log.Identify(r.Identifier)
}

// Defer centralizes the processing of response
func (r *Resource) Defer() {
	var b bytes.Buffer

	r.Identify()
	r.Record("Content-Type", r.Response.Header().Get(ContentType))
	defer func() {
		if r.Status == 0 {
			r.Record("Status", http.StatusInternalServerError)
		} else {
			r.Record("Status", r.Status)
		}

		if r.RawBody != nil {

			if fmt.Sprint(r.RawBody) == "[]" {
				emptyResponse, _ := json.Marshal(make([]int64, 0))
				r.Record("Response", string(emptyResponse))
			} else {
				enc := json.NewEncoder(&b)
				enc.SetEscapeHTML(false)
				enc.Encode(r.RawBody)
				var response = strings.Replace(string(b.Bytes()), "\n", "", -1)
				response = helper.MaskSensitive(response)
				r.Record("Response", response)
			}
		}

		r.Record("End", time.Now().Format(gorm.SQLDatetime))

		r.Log.ShowCount(true)
		r.Log.Dump()

	}()

	if rec := recover(); rec != nil {
		r.Record("Recovery", fmt.Sprint(rec))
		r.FormatException(r, errors.New(fmt.Sprint(rec)))
		r.Identifier = fmt.Sprintf("%s\t", secure.MD5(time.Now().UTC().Format(gorm.SQLDatetime)+
			r.Request.URL.String()+
			r.Request.Method+
			string(r.Body())))
	}

	r.Log.Identify(r.Identifier)
}

// Defer centralizes the processing of response
func (r *Resource) Defer() {
	var b bytes.Buffer

	r.Identify()
	r.Record("Content-Type", r.Response.Header().Get(ContentType))
	defer func() {
		if r.Status == 0 {
			r.Record("Status", http.StatusInternalServerError)
		} else {
			r.Record("Status", r.Status)
		}

		if r.RawBody != nil {

			if fmt.Sprint(r.RawBody) == "[]" {
				emptyResponse, _ := json.Marshal(make([]int64, 0))
				r.Record("Response", string(emptyResponse))
			} else {
				enc := json.NewEncoder(&b)
				enc.SetEscapeHTML(false)
				enc.Encode(r.RawBody)
				var response = strings.Replace(string(b.Bytes()), "\n", "", -1)
				response = helper.MaskSensitive(response)
				r.Record("Response", response)
			}
		}

		r.Record("End", time.Now().Format(gorm.SQLDatetime))

		r.Log.ShowCount(true)
		r.Log.Dump()

	}()

	if rec := recover(); rec != nil {
		r.Record("Recovery", fmt.Sprint(rec))
		r.FormatException(r, errors.New(fmt.Sprint(rec)))

		if r.Status == 0 {
			r.FormatException(r, errors.New(errs.ErrInternalServerError))
			out.JSON(r.Response, http.StatusInternalServerError, r.RawBody)
			return
		}

		if r.RawBody != nil {
			out.JSON(r.Response, r.Status, r.RawBody)
			return
		}
		out.Status(r.Response, r.Status)
	}
}

// Done will handle the primary response processing
func (r *Resource) Done() {
	defer func() {
		if recover := recover(); recover != nil {
			r.FormatException(r, errors.New(fmt.Sprint(recover)))
		}
	}()

	body := r.RawBody
	status := r.Status

	if body == nil {
		out.Status(r.Response, status)
		return
	}

	out.JSON(r.Response, r.Status, body)
}

// GetClientIP returns the client's IP
func (r *Resource) GetClientIP() (string, error) {
	req := r.Request
	ipAddress := req.RemoteAddr

	if ip := req.Header.Get("X-Forwarded-For"); "" != ip {
		ipAddress = ip

		// X-Forwarded-For might contain multiple IPs. Get the last one.
		if strings.Contains(ipAddress, ",") {
			ips := strings.Split(ipAddress, ",")
			ipAddress = strings.Trim(ips[len(ips)-1], " ")
		}
	}

	var (
		ip  net.IP
		err error
	)

	if -1 != strings.Index(ipAddress, ":") {
		if ipAddress, _, err = net.SplitHostPort(ipAddress); nil != err {
			return "", fmt.Errorf("GetClientIP SplitHostPort Error: %v", err)
		}
	}

	if err := ip.UnmarshalText([]byte(ipAddress)); nil != err {
		return "", fmt.Errorf("GetClientIP UnmarshalText Error: %v", err)
	}

	return ipAddress, nil
}

// Body returns the body from the request
func (r *Resource) Body() []byte {
	body, _ := ioutil.ReadAll(r.Request.Body)
	// Restore the io.ReadCloser to its original state
	r.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body
}

// GetParams reads the parameters and transforms to a request structure
func (r *Resource) GetParams(o interface{}, allowEmptyBody ...bool) (err error) {
	ct := entity.GetContentType(r.Request)
	if !entity.ValidContentType(ct) {
		r.Status = http.StatusUnsupportedMediaType
		r.Response.Header().Set("Accept", rest.ContentTypeJSON)
		err = errors.New("Unsupported media type")
		return
	}

	var canbeEmpty = false
	if len(allowEmptyBody) > 0 {
		canbeEmpty = allowEmptyBody[0]
	}

	if !canbeEmpty && len(r.Body()) < 1 {
		r.Status = http.StatusUnprocessableEntity
		err = errors.New(errs.ErrEmptyBodyContent)
		return
	}

	if len(r.Body()) > 0 {
		if entity.CheckJSONCT(ct) {
			err = json.Unmarshal(r.Body(), o)
			if err != nil {
				r.Status = http.StatusBadRequest
				err = errors.New(errs.ErrRequestBodyInvalid)
				return
			}
		} else if entity.CheckFormDataCT(ct) {
			var frmInput url.Values
			frmInput, err = entity.ParseForm(ct, r.Request, true)
			if err == nil {
				decoder := schema.NewDecoder()
				decoder.SetAliasTag("json")
				err = decoder.Decode(o, frmInput)
				if err != nil {
					r.Status = http.StatusBadRequest
					err = errors.New(errs.ErrRequestBodyInvalid)
					return
				}
			}
		}
	}

	return
}

// GetMapParams for no struct request, convert it to map
func (r *Resource) GetMapParams() (params map[string]interface{}, err error) {
	ct := entity.GetContentType(r.Request)
	if !entity.ValidContentType(ct) {
		r.Response.Header().Set("Accept", rest.ContentTypeJSON)
		err = errors.New(errs.ErrUnsupportedRequest)
		return
	}
	if len(r.Body()) < 1 {
		err = errors.New(errs.ErrUnsupportedRequest)
		return
	}

	params = make(map[string]interface{})

	if entity.CheckJSONCT(ct) {
		if err = json.Unmarshal(r.Body(), &params); nil != err {
			err = errors.New(errs.ErrUnsupportedRequest)
			return nil, err
		}
	} else if entity.CheckFormDataCT(ct) {
		var frmInput url.Values
		frmInput, err = entity.ParseForm(ct, r.Request, true)

		// loop frmInput to format the final params
		for key, value := range frmInput {
			paramsMap := make(map[string]interface{})
			if err = json.Unmarshal([]byte(value[0]), &paramsMap); nil == err {
				params[key] = paramsMap
			} else {
				params[key] = value[0]
				err = nil
			}
		}
	}
	return
}

// GetQuery fetches the value from the query string and d if empty
func (r *Resource) GetQuery(key string, d string) string {
	v := r.Request.URL.Query().Get(key)

	if v == "" {
		return d
	}

	return v
}

// ValidContentType checks if the requested content type is supported by the response
func (r *Resource) ValidContentType(expectedContentType string) bool {
	contentType := r.Request.Header.Get(ContentType)

	if contentType == "" || contentType == expectedContentType {
		return true
	}

	r.Response.WriteHeader(http.StatusUnsupportedMediaType)
	r.Response.Header().Set("Accept", expectedContentType)
	return false
}

// TrimStructs trim leading and trailing whitespaces on struct fields
func (r *Resource) TrimStructs(v interface{}) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	var structMap map[string]interface{}
	if err := json.Unmarshal(bytes, &structMap); err != nil {
		return err
	}

	structMap = TrimMapStringInterface(structMap).(map[string]interface{})
	bytes2, err := json.Marshal(structMap)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes2, v); err != nil {
		return err
	}

	return nil
}

// TrimMapStringInterface trim leading and trailing whitespaces on struct fields
func TrimMapStringInterface(data interface{}) interface{} {
	if values, valid := data.([]interface{}); valid {
		for i := range values {
			data.([]interface{})[i] = TrimMapStringInterface(values[i])
		}
	} else if values, valid := data.(map[string]interface{}); valid {
		for k, v := range values {
			data.(map[string]interface{})[k] = TrimMapStringInterface(v)
		}
	} else if value, valid := data.(string); valid {
		data = strings.TrimSpace(value)
	}

	return data
}
