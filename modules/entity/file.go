package entity

import (
	"os"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
)

// Writer ...
type Writer func(s string)

const (
	// CharSetUtf8 ...
	CharSetUtf8 = "utf-8"
)

// CreateTempFile function ...
func CreateTempFile() (file *os.File, err error) {
	tempFile := "/tmp/" + os.Getenv(constant.EnvAppName) + "-" + GenerateRandomHash(15)
	file, err = os.Create(tempFile)

	if err != nil {
		return nil, err
	}

	return
}
