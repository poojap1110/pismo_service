package validation

import (
	"encoding/json"
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	moduleConfig "bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/config"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/container"
	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/tcp/config/getProduct"
	"bitbucket.org/matchmove/go-tools/array"
	"gopkg.in/validator.v2"
)

const (
	TagJSONKey = "json"
	TagValidationKey = "validation_key"
	TagAddValidationKey = "add_validation"
)

// init the validation
func init() {
	validator.SetValidationFunc("minlen", MinLen)
	validator.SetValidationFunc("maxlen", MaxLen)
	validator.SetValidationFunc("reqlen", ReqLen)
}

// Required validate if field exist
func Required(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if st.String() == "" {
		return errors.New("Parameter " + param + " is a required field.")
	}

	return nil
}

// MaxLen - Max Length validation
func MaxLen(v interface{}, param string) (err error) {
	st := reflect.ValueOf(v)
	if st.String() == "" {
		return nil
	}
	sl, _ := strconv.Atoi(param)
	if sl < len(st.String()) {
		return errors.New("")
	}

	return nil
}

// MinLen - Min Length validation
func MinLen(v interface{}, param string) (err error) {
	st := reflect.ValueOf(v)
	if st.String() == "" {
		return nil
	}
	sl, _ := strconv.Atoi(param)
	if sl > len(st.String()) {
		return errors.New("min_len: " + param)
	}

	return nil
}

// ReqLen - Required Length validation
func ReqLen(v interface{}, param string) (err error) {
	st := reflect.ValueOf(v)
	if st.String() == "" {
		return nil
	}
	sl, _ := strconv.Atoi(param)
	if sl != len(st.String()) {
		return errors.New("Required length")
	}

	return nil
}

// Detail ...
type Detail struct {
	Code        string
	Description string
}

// GetDetails for enum values with Code and Description
func GetDetails(jsonDetails string) (detailsArray []string) {

	var details []Detail

	json.Unmarshal([]byte(jsonDetails), &details)
	for _, value := range details {
		detailsArray = append(detailsArray, value.Code)
	}

	return
}

// VerifyPassword ...
func VerifyPassword(password string, allowNumber bool, allowCharacter bool, atleastUpperCase int) (errs []error) {
	var uppercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	var upperCaseCount = 0

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
		case unicode.IsUpper(ch):
			uppercasePresent = true
			upperCaseCount++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
		case ch == ' ' || unicode.IsLower(ch):
			continue
		}
	}

	if allowNumber {
		if !numberPresent {
			errs = []error{errors.New(constant.ErrorInvalidPassword + ": At least one numeric character should be present")}
			return
		}
	} else {
		if numberPresent {
			errs = []error{errors.New(constant.ErrorInvalidPassword + " : Numeric characters are not allowed")}
			return
		}
	}

	if uppercasePresent {
		if atleastUpperCase == 0 {
			errs = []error{errors.New(constant.ErrorInvalidPassword + " : Uppercase characters are not allowed")}
			return
		} else {
			if upperCaseCount < atleastUpperCase {
				errs = []error{errors.New(constant.ErrorInvalidPassword + " : At least " + strconv.Itoa(atleastUpperCase) + " uppercase character should be present")}
				return
			}
		}

	} else {
		if atleastUpperCase > 0 {
			errs = []error{errors.New(constant.ErrorInvalidPassword + " : At least " + strconv.Itoa(atleastUpperCase) + " uppercase character should be present")}
			return
		}
	}

	if allowCharacter {
		if !specialCharPresent {
			errs = []error{errors.New(constant.ErrorInvalidPassword + " : At least one special character should be present")}
			return
		}
	} else {
		if specialCharPresent {
			errs = []error{errors.New(constant.ErrorInvalidPassword + " : Special characters are not allowed")}
			return
		}
	}

	return
}

// Rule for validation purposes
func Rule(requestKey string, requestValue interface{}, rule string, value string) (errs []error) {
	if rule != "required" {
		if requestValue == nil {
			return
		} else {
			if requestValue.(string) == "" {
				return
			}
		}
	}

	switch rule {
	case "enum":
		enum := strings.Split(value, "enum=")
		enumValues := GetDetails(enum[1])

		// Check if value passed exists in array of enumValues
		if exists, _ := array.InArray(requestValue.(string), enumValues); exists != true {
			errs = []error{errors.New(constant.ErrorNotAcceptableValues + ": " + requestKey + ": " + strings.Join(enumValues, ","))}
			return
		}

		break

	case "equal":
		value := strings.Split(value, "equal=")
		if value[1] != requestValue.(string) {
			errs = []error{errors.New(constant.ErrorParameterNotEqual + ": " + requestKey + ": " + value[1])}
			return
		}

		break

	case "pattern":
		val := strings.Split(value, "pattern=")
		if val[1] == "email" {
			if err := validator.Valid(requestValue, value); err != nil {
				errs = []error{errors.New(constant.ErrorInvalidPattern + ": " + requestKey)}
				return
			}
		} else {
			if ok, _ := regexp.MatchString(`^`+val[1]+`$`, requestValue.(string)); !ok {
				errs = []error{errors.New(constant.ErrorInvalidPattern + ": " + requestKey)}
				return
			}
		}

		break

	case "date":
		format := strings.Split(value, "date=")
		vals := strings.Split(format[1], "|")

		input := requestValue.(string)
		layout := vals[0]
		t, err := time.Parse(layout, input)
		if err != nil {
			errs = []error{errors.New(constant.ErrorInvalidParamValue + ": " + requestKey + ": date should be in same format with `" + layout + "`")}
			return
		}

		utcInput := t.UTC()

		if strings.Contains(vals[1], "future") {
			future := strings.Split(vals[1], "=")
			// Check if not allowed for future
			today := time.Now().UTC().Format(layout)
			parsedDateToday, _ := time.Parse(layout, today)

			utcToday := parsedDateToday.UTC()

			if future[1] == "1" {
				if utcInput.Before(utcToday) {
					errs = []error{errors.New(constant.ErrorInvalidParamValue + ": " + requestKey + ": Expired/Past date is not allowed ")}
					return
				}
			} else {
				if utcInput.After(utcToday) {
					errs = []error{errors.New(constant.ErrorInvalidParamValue + ": " + requestKey + ": Future date is not allwoed ")}
					return
				}
			}
		}

	case "password":
		val := strings.Split(value, "password=")
		vals := strings.Split(val[1], "|")

		var (
			allowNumber      = false
			allowCharacter   = false
			atleastUpperCase = 0
		)

		if val[1] == "default" {
			allowNumber = true
			allowCharacter = true
		} else {
			for _, k := range vals {
				if k == "number" {
					allowNumber = true
				}

				if k == "char" {
					allowCharacter = true
				}

				if strings.Contains(k, "upper=") {
					count := strings.Split(k, "=")
					atleastUpperCase, _ = strconv.Atoi(count[1])
				}
			}
		}

		errs = VerifyPassword(requestValue.(string), allowNumber, allowCharacter, atleastUpperCase)
		if len(errs) > 0 {
			return
		}

	case "required":
		if requestValue == nil {
			errs = []error{errors.New(constant.ErrorParameterRequired + ": " + requestKey)}
			return
		}

		if err := validator.Valid(requestValue, "required"); err != nil {
			errs = []error{errors.New(constant.ErrorParameterRequired + ": " + requestKey)}
			return
		}

		break

	default:
		var errMsg string
		if err := validator.Valid(requestValue, value); err != nil {

			val := strings.Split(value, "=")

			switch rule {
			case "reqlen":
				errMsg = constant.ErrorParameterRequiredLength + ": " + requestKey + ": " + val[1]
			case "min":
				errMsg = constant.ErrorParameterMinValue + ": " + requestKey + ": " + val[1]
			case "max":
				errMsg = constant.ErrorParameterMaxValue + ": " + requestKey + ": " + val[1]
			case "minlen":
				errMsg = constant.ErrorParameterMinLengthValue + ": " + requestKey + ": " + val[1]
			case "maxlen":
				errMsg = constant.ErrorParameterMaxLengthValue + ": " + requestKey + ": " + val[1]
			}

			errs = []error{errors.New(errMsg)}

			return
		}

		break
	}

	return
}


// GetValidationRules function - returns the validation rules for the specified groupname & fields
func GetValidationRules(groupName string, fields []string, additionalRules []string, productConfigs getProduct.ProductConfigs) (validationRules map[string]string) {
	validationRules = make(map[string]string)

	for _, v := range fields {
		validationRules[v] = ""
	}

	for _, v := range additionalRules {
		validationRules[v] = v
	}

	for _, v := range productConfigs.Products {
		if v.GroupName == groupName {
			if _, ok := validationRules[v.PropertyName]; ok {
				validationRules[v.PropertyName] = v.Validation
			}
		}
	}

	return
}

// GetProductConfig function - retrieve product config
func GetProductConfig(cont *container.Container, productID string, userGroup string, tenantHashID string) (pc getProduct.ProductConfigs, errs []error) {
	var (
		resp string
		err  error
	)

	if resp, err = getProduct.Call(cont, getProduct.FormatRequest(productID)); err != nil {
		errs = []error{errors.New(constant.ErrorCodeInternalServerProductNotFound)}
		return
	}

	res, err := getProduct.FormatResponse(resp)
	pc.RetrieveProductConfig(res, userGroup)

	if len(pc.Products) == 0 {
		errs = []error{errors.New(constant.ErrorCodeInternalServerProductNotDefined)}
		return
	}

	return
}

// GetValidationKeys function - gets validation_key and add_validation tags on the request struct.
func GetValidationKeys(request interface{}) (valKeys []string, addRules []string) {
	sv := reflect.ValueOf(request)
	st := reflect.TypeOf(request)

	if sv.Kind() == reflect.Ptr && !sv.IsNil() {
		return GetValidationKeys(sv.Elem().Interface())
	}

	if sv.Kind() != reflect.Struct && sv.Kind() != reflect.Interface {
		panic("GetValidationKeys error: " + validator.ErrUnsupported.Error())
	}

	nfields := sv.NumField()
	valKeys = []string{}
	addRules = []string{}

	for i := 0; i < nfields; i++ {
		tag := st.Field(i).Tag.Get(TagValidationKey)
		addValTag := st.Field(i).Tag.Get(TagAddValidationKey)

		if tag == "-" {
			continue
		}

		if tag != "" {
			valKeys = append(valKeys, tag)
		}

		if addValTag != "" && addValTag != "-" {
			addRules = append(addRules, addValTag)
		}
	}

	return
}

// Validate function - prepares and performs complete validation on the request struct.
// request struct must have the `json` and `validation_key` tags declared on its members.
func Validate(cont *container.Container, productID string, userGroup string, tenantHashID string, request interface{}) []error {
	pc, errs := GetProductConfig(cont, productID, userGroup, tenantHashID)

	if len(errs) > 0 {
		return errs
	}

	// populate config object's tenant configs on container
	configIns := moduleConfig.GetInstance(cont)
	configIns.RetrieveTenantConfigs(cont, tenantHashID, nil, nil)

	validationKeys, additionalRules := GetValidationKeys(request)
	rules := GetValidationRules(userGroup, validationKeys, additionalRules, pc)

	return ManualValidate(rules, request)
}

// ManualValidate function - performs validation manually
// request struct must have the `json` and `validation_key` tags declared on its members.
func ManualValidate(rules map[string]string, request interface{}) []error {
	sv := reflect.ValueOf(request)
	st := reflect.TypeOf(request)

	if sv.Kind() == reflect.Ptr && !sv.IsNil() {
		return ManualValidate(rules, sv.Elem().Interface())
	}

	if sv.Kind() != reflect.Struct && sv.Kind() != reflect.Interface {
		panic("GetValidationKeys error: " + validator.ErrUnsupported.Error())
	}

	nfields := sv.NumField()
	errs := []error{}

	for i := 0; i < nfields; i++ {
		jTag := st.Field(i).Tag.Get(TagJSONKey)
		valTag := st.Field(i).Tag.Get(TagValidationKey)
		addValTag := st.Field(i).Tag.Get(TagAddValidationKey)

		// | valT | addValT | result
		//   ""      ""       true
		//   "-"     "-"      true
		//   "a"     ""       false
		//   "a"     "-"      false
		//   ""      "a"      false
		//   "-"     "a"      false
		//   "a"     "a"      false
		if (jTag == "-" || jTag == "") || ((valTag == "" || valTag == "-") && (addValTag == "" || addValTag == "-")) {
			continue
		}

		validations := BreakValidation(rules[valTag])

		if addValTag != "" && addValTag != "-" {
			addVals := BreakValidation(addValTag)

			if len(addVals) > 0 {
				for _, val := range addVals {
					validations = append(validations, val)
				}
			}
		}

		if len(validations) > 0 {
			for _, val := range validations {
				if len(val) > 0 {
					rule, value := BreakRuleValue(val)

					if rule != "" || value != "" {
						if sv.Field(i).Kind() == reflect.String { // as Rule() only supports string as of writing
							errs = Rule(jTag, sv.Field(i).String(), rule, value)

							if len(errs) > 0 {
								return errs
							}
						}
					}
				}
			}
		}
	}

	return nil
}

// BreakRuleValue function - breaks ruleval combination.
func BreakRuleValue(ruleVal string) (rule, value string) {
	if strings.Contains(ruleVal, "=") {
		str := strings.Split(ruleVal, "=")
		rule = str[0]
		value = ruleVal
	} else {
		rule = ruleVal
	}

	return
}

// BreakValidation function - breaks validation rules
func BreakValidation(val string) []string {
	return strings.Split(val, "||")
}