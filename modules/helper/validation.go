package helper

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"bitbucket.org/matchmove/go-tools/array"
	"bitbucket.org/matchmove/go-valid"
	"bitbucket.org/matchmove/integration-svc-aub/app/errs"
	"github.com/jmoiron/sqlx"
	"gopkg.in/validator.v2"
)

const (
	TagJSONKey          = "json"
	TagValidationKey    = "validation_key"
	TagAddValidationKey = "add_validation"
)

// init the validation
func init() {
	validator.SetValidationFunc("required", valid.Required)
	validator.SetValidationFunc("pattern", valid.Pattern)
	validator.SetValidationFunc("minlen", MinLen)
	validator.SetValidationFunc("maxlen", MaxLen)
	validator.SetValidationFunc("reqlen", ReqLen)
	validator.SetValidationFunc("numeric", Numeric)
}

// MaxLen - Max Length validation
func MaxLen(v interface{}, param string) (err error) {
	st := reflect.ValueOf(v)
	if st.String() == "" {
		return nil
	}
	sl, _ := strconv.Atoi(param)
	if sl < len(st.String()) {
		return errors.New("Error Maximum Length for param : " + param)
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
		return errors.New("Error Minimum Length for param : " + param)
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
		return errors.New("Error Required Length for param : " + param)
	}

	return nil
}

// Numeric - Number only validation
func Numeric(v interface{}, param string) (err error) {
	st := reflect.ValueOf(v)
	if st.String() == "" {
		return nil
	}

	_, err = strconv.ParseFloat(st.String(), 64)
	if err != nil {
		return errors.New("Only numbers allowed for :" + param)
	}

	return nil
}

// DateDifference get date difference
func DateDifference(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}

	if a.After(b) {
		a, b = b, a
	}

	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}

	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}

	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}

	if month < 0 {
		month += 12
		year--
	}

	return
}

// VerifyPassword ...
func VerifyPassword(password string, allowNumber bool, allowCharacter bool, letterShouldBePresent bool, atleastUpperCase int, atleastLowerCase int) (err error) {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	var whitespace bool
	var upperCaseCount = 0
	var lowerCaseCount = 0

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
		case unicode.IsUpper(ch):
			uppercasePresent = true
			upperCaseCount++
		case unicode.IsLower(ch):
			lowercasePresent = true
			lowerCaseCount++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
		case ch == ' ':
			whitespace = true
		case unicode.IsLower(ch):
			continue
		}
	}

	if allowNumber {
		if !numberPresent {
			err = errors.New(errs.ErrParameterInvalidFormat + ": password" + ": At least one numeric character should be present")
			return
		}
	} else {
		if numberPresent {
			err = errors.New(errs.ErrParameterInvalidFormat + ": password" + ": Numeric characters are not allowed")
			return
		}
	}

	if letterShouldBePresent {
		if !uppercasePresent || !lowercasePresent {
			err = errors.New(errs.ErrParameterInvalidFormat + ": password" + ": At least one letter should be present")
			return
		}
	}

	if atleastUpperCase > 0 {
		if !uppercasePresent {
			err = errors.New(errs.ErrParameterInvalidFormat + ": password" + ": At least " + strconv.Itoa(atleastUpperCase) + " uppercase character should be present")
			return
		}

		if upperCaseCount < atleastUpperCase {
			err = errors.New(errs.ErrParameterInvalidFormat + ": password" + ": At least " + strconv.Itoa(atleastUpperCase) + " uppercase character should be present")
			return
		}
	}

	if atleastLowerCase > 0 {
		if !lowercasePresent {
			err = errors.New(errs.ErrParameterInvalidFormat + ": password" + ": At least " + strconv.Itoa(atleastUpperCase) + " uppercase character should be present")
			return
		}

		if lowerCaseCount < atleastLowerCase {
			err = errors.New(errs.ErrParameterInvalidFormat + ": password" + ": At least " + strconv.Itoa(atleastLowerCase) + " lowercase character should be present")
			return
		}
	}

	if allowCharacter {
		if !specialCharPresent {
			err = errors.New(errs.ErrParameterInvalidFormat + ": password" + ": At least one special character should be present")
			return
		}
	} else {
		if specialCharPresent {
			err = errors.New(errs.ErrParameterInvalidFormat + ": password" + ": Special characters are not allowed")
			return
		}
	}

	if whitespace {
		err = errors.New(errs.ErrParameterInvalidFormat + ": password" + ": Whitespaces are not allowed")
		return
	}

	return
}

// Rule for validation purposes
func Rule(requestKey string, requestValue interface{}, rule string, value string, db *sqlx.DB) (err error) {
	if rule != "required" {
		if requestValue == nil || requestValue.(string) == "" {
			return
		}
	}

	switch rule {
	case "enum":
		enum := strings.Split(value, "enum=")
		enumArray := strings.Split(enum[1], ",")

		// Check if value passed exists in array of enumValues
		if exists, _ := array.InArray(requestValue.(string), enumArray); exists != true {
			err = errors.New(errs.ErrParameterNotAcceptableValues + ": " + requestKey + ": " + strings.Join(enumArray, ","))
			return err
		}

	case "equal":
		value := strings.Split(value, "equal=")
		if value[1] != requestValue.(string) {
			err = errors.New(errs.ErrParameterShouldBeEqual + ": " + requestKey + ": " + value[1])
			return err
		}

		break
	case "date":
		format := strings.Split(value, "date=")
		vals := strings.Split(format[1], "|")

		input := requestValue.(string)
		layout := vals[0]
		t, err := time.Parse(layout, input)

		var l string

		switch layout {
		case "2006-01-02":
			l = "ISO 8601 Date, e.g `2006-01-02`"
			break
		case "2006-01-02T15:04:05-07:00":
			l = "ISO 8601 Date with Timezone, e.g `2006-01-02T15:04:05-07:00`"
			break
		}

		if err != nil {
			err = errors.New(errs.ErrParameterInvalidFormat + ": " + requestKey + ": Date should be in format " + l)
			return err
		}

		utcInput := t.UTC()

		for i := range vals {
			if strings.Contains(vals[i], "future") {
				future := strings.Split(vals[i], "=")
				// Check if not allowed for future
				today := time.Now().Format(layout)
				parsedDateToday, _ := time.Parse(layout, today)

				utcToday := parsedDateToday.UTC()

				if future[1] == "1" {
					if utcInput.Before(utcToday) {
						err = errors.New(errs.ErrParameterInvalidFormat + ": " + requestKey + ": Expired/Past date is not allowed ")
						return err
					}
				} else {
					if utcInput.After(utcToday) {
						err = errors.New(errs.ErrParameterInvalidFormat + ": " + requestKey + ": Future date is not allowed ")
						return err
					}
				}
			}

			if strings.Contains(vals[i], "min") {
				minValue := strings.Split(vals[i], "=")
				min, _ := strconv.Atoi(minValue[1])

				// calculate years, month, days and time betwen dates
				year, _, _, _, _, _ := DateDifference(t, time.Now())

				if year < min {
					err = errors.New(errs.ErrParameterInvalidFormat + ": " + requestKey + ": Minimum age should be " + minValue[1] + " years old")
					return err
				}
			}
		}
		break
	case "password":
		val := strings.Split(value, "password=")
		vals := strings.Split(val[1], "|")

		var (
			allowNumber           = false
			allowCharacter        = false
			letterShouldBePresent = false
			atleastUpperCase      = 0
			atleastLowerCase      = 0
		)

		if val[1] == "default" {
			allowNumber = true
			allowCharacter = true
			letterShouldBePresent = true
		} else {
			for _, k := range vals {
				if k == "number" {
					allowNumber = true
				}

				if k == "char" {
					allowCharacter = true
				}

				if k == "letter" {
					letterShouldBePresent = true
				}

				if strings.Contains(k, "upper=") {
					count := strings.Split(k, "=")
					atleastUpperCase, _ = strconv.Atoi(count[1])
				}

				if strings.Contains(k, "lower=") {
					count := strings.Split(k, "=")
					atleastLowerCase, _ = strconv.Atoi(count[1])
				}
			}
		}

		err = VerifyPassword(requestValue.(string), allowNumber, allowCharacter, letterShouldBePresent, atleastUpperCase, atleastLowerCase)
		if err != nil {
			err = errors.New(errs.ErrParameterInvalidFormat + ": " + requestKey + ": At least should contain one number, one upper case, one lower case, and one special character")
			return err
		}

		break
	case "pattern":
		val := strings.Split(value, "pattern=")
		if val[1] == "email" {
			if err = validator.Valid(requestValue, value); err != nil {
				err = errors.New(errs.ErrParameterInvalid + ": " + requestKey)
				return
			}
		} else {
			if ok, _ := regexp.MatchString(`^`+val[1]+`$`, requestValue.(string)); !ok {
				err = errors.New(errs.ErrParameterInvalid + ": " + requestKey)
				return
			}
		}

		break

	case "numeric":
		if err = validator.Valid(requestValue, rule); err != nil {
			err = errors.New(errs.ErrParameterNumeric + ": " + requestKey)
			return
		}

		break

	case "required":
		if requestValue == nil {
			err = errors.New(errs.ErrParameterRequired + ": " + requestKey)
			return
		}

		if err = validator.Valid(requestValue, rule); err != nil {
			err = errors.New(errs.ErrParameterRequired + ": " + requestKey)
			return
		}

		break

	default:
		var errMsg string
		if err = validator.Valid(requestValue, value); err != nil {

			val := strings.Split(value, "=")

			switch rule {
			case "reqlen":
				errMsg = errs.ErrParameterRequiredLength + ": " + requestKey + ": " + val[1]
			case "min":
				errMsg = errs.ErrParameterMinValue + ": " + requestKey + ": " + val[1]
			case "max":
				errMsg = errs.ErrParameterMaxValue + ": " + requestKey + ": " + val[1]
			case "minlen":
				errMsg = errs.ErrParameterMinLengthValue + ": " + requestKey + ": " + val[1]
			case "maxlen":
				errMsg = errs.ErrParameterMaxLengthValue + ": " + requestKey + ": " + val[1]
			}

			return errors.New(errMsg)
		}

		break
	}

	return
}

// GetValidationRules function - returns the validation rules for the specified groupname & fields
func GetValidationRules(fields []string, additionalRules []string) (validationRules map[string]string) {
	validationRules = make(map[string]string)

	for _, v := range fields {
		validationRules[v] = ""
	}

	for _, v := range additionalRules {
		validationRules[v] = v
	}

	return
}

// GetValidationKeys function - gets validation_key and add_validation tags on the request struct.
func GetValidationKeys(request interface{}, group string) (valKeys []string, addRules []string) {
	sv := reflect.ValueOf(request)
	st := reflect.TypeOf(request)

	if sv.Kind() == reflect.Ptr && !sv.IsNil() {
		return GetValidationKeys(sv.Elem().Interface(), group)
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
		groupTag := st.Field(i).Tag.Get(group + "_validation")

		if tag == "-" {
			continue
		}

		if tag != "" {
			valKeys = append(valKeys, tag)
		}

		if groupTag != "" {
			addRules = append(addRules, groupTag)
		}

		if addValTag != "" && addValTag != "-" {
			addRules = append(addRules, addValTag)
		}
	}

	return
}

// Validate function - prepares and performs complete validation on the request struct.
// request struct must have the `json` and `validation_key` tags declared on its members.
func Validate(request interface{}, db *sqlx.DB, group string) error {
	validationKeys, additionalRules := GetValidationKeys(request, group)
	rules := GetValidationRules(validationKeys, additionalRules)

	return ManualValidate(rules, request, db, group)
}

// ManualValidate function - performs validation manually
// request struct must have the `json` and `validation_key` tags declared on its members.
func ManualValidate(rules map[string]string, request interface{}, db *sqlx.DB, group string) (err error) {
	sv := reflect.ValueOf(request)
	st := reflect.TypeOf(request)

	if sv.Kind() == reflect.Ptr && !sv.IsNil() {
		return ManualValidate(rules, sv.Elem().Interface(), db, group)
	}

	if sv.Kind() != reflect.Struct && sv.Kind() != reflect.Interface {
		panic("GetValidationKeys error: " + validator.ErrUnsupported.Error())
	}

	nfields := sv.NumField()
	for i := 0; i < nfields; i++ {
		jTag := st.Field(i).Tag.Get(TagJSONKey)
		valTag := st.Field(i).Tag.Get(TagValidationKey)
		addValTag := st.Field(i).Tag.Get(TagAddValidationKey)
		groupTag := st.Field(i).Tag.Get(group + "_validation")

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

		if groupTag != "" {
			addVals := BreakValidation(groupTag)

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
							err = Rule(jTag, sv.Field(i).String(), rule, value, db)

							if err != nil {
								return
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
