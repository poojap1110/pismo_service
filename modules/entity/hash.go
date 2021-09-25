package entity

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"time"

	"bitbucket.org/matchmove/go-tools/secure"
)

// GenerateRandomHash function ...
func GenerateRandomHash(hashLength int) string {
	unixTime := UnixTimestamp()
	return fmt.Sprintf("%s%s", unixTime, secure.RandomString(hashLength-len(unixTime), []rune(secure.RuneAlNumCS)))
}

// UnixTimestamp function - returns utc int timestamp ...
func UnixTimestamp() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

// GeneratePureHash function ...
func GeneratePureHash(hashLength int) string {
	return secure.RandomString(hashLength, []rune(secure.RuneAlNumCS))
}

func GenerateRansomMD5() string {
	unixTime := UnixTimestamp()
	randString := secure.RandomString(32-len(unixTime), []rune(secure.RuneAlNumCS))
	data := []byte(randString)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func GenerateRandomSHA() string {
	unixTime := UnixTimestamp()
	randString := secure.RandomString(32-len(unixTime), []rune(secure.RuneAlNumCS))
	data := []byte(randString)
	return fmt.Sprintf("%x", sha1.Sum(data))
}
