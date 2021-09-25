package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"github.com/google/uuid"
)

func ValidatePasswordUserFields(password string, firstname string, lastname string) (err error) {
	fname_exist := strings.Contains(password, firstname)
	lname_exist := strings.Contains(password, lastname)

	if fname_exist == true || lname_exist == true {
		err = errors.New(constant.UserPasswordFirstLastName)
		return
	}

	return nil
}

func Sanitize(original string) string {
	extraWhiteSpaces := regexp.MustCompile(`\s+`)
	updatedString := extraWhiteSpaces.ReplaceAllString(original, " ")

	return updatedString
}

func ComputeHmac256(message string) string {
	h := sha256.New()
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

// GetRandomID ...
func GetRandomID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		return ""
	}

	return id.String()
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

func GenerateRandomNumberString(n int) (string, error) {
	const letters = "0123456789"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
