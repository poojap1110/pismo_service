package helper

import (
	b64 "encoding/base64"
	"os"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	jose "github.com/dvsekhvalnov/jose2go"
)

// DecryptJWS ...
func DecryptJWS(encrypted string) (payload string, err error) {

	jweKeyEncoded := os.Getenv(constant.JweSymmetricKey)
	jweKey, _ := b64.StdEncoding.DecodeString(jweKeyEncoded)

	payload, _, err = jose.Decode(encrypted, jweKey)

	if err != nil {
		return
	}

	return
}
