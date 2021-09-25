package helper

import (
	"encoding/json"
	"strings"

	"bitbucket.org/matchmove/go-tools/array"
)

// Mask mask specific length of a string
func Mask(s string, start, limit int, exceptCharacter []string) string {
	rs := []rune(s)
	for i := start; i < len(rs)-limit; i++ {
		if len(exceptCharacter) > 0 {
			if exists, _ := array.InArray(rs[i], exceptCharacter); exists {
				continue
			}
		}

		rs[i] = 'X'
	}

	return string(rs)
}

// MaskSensitive mask sensitive values on logs
func MaskSensitive(values string) string {
	var (
		body      = make(map[string]interface{})
		sensitive = []string{
			"password",
			"data",
			"mask_assoc_number",
			"mask_cvv",
			"mask_expiry",
			"mask_number",
		}
	)

	err := json.Unmarshal([]byte(values), &body)
	if err == nil {
		for _, s := range sensitive {
			var maskKey string
			if strings.Contains(s, "mask_") {
				maskKey = strings.Replace(s, "mask_", "", -1)
			}
			if _, ok := body[maskKey]; ok {
				if maskKey != "" {
					body[maskKey] = Mask(body[maskKey].(string), 0, 4, nil)
				} else {
					body[maskKey] = ""
				}
			} else {
				if _, ok := body[s]; ok {
					body[s] = ""
				}
			}
		}

		bodyByte, _ := json.Marshal(body)
		return string(bodyByte)
	}

	return values
}
