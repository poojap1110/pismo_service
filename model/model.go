package model

import (
	"bitbucket.org/matchmove/go-tools/secure"
	"database/sql"
	"fmt"
	"time"
)

// NullStringToString function ...
func NullStringToString(nullString sql.NullString) string {
	return nullString.String
}

// StringToNullString function ...
func StringToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// Int64ToNullInt64 function ...
func Int64ToNullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}

// NullInt64ToInt64 function ...
func NullInt64ToInt64(nullInt64 sql.NullInt64) int64 {
	return nullInt64.Int64
}

// IntToNullInt64 function ...
func IntToNullInt64(i int) sql.NullInt64 {
	return Int64ToNullInt64(int64(i))
}

// UnixTimestamp function - returns utc int timestamp ...
func UnixTimestamp() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

// GenerateRandomHash function ...
func GenerateRandomHash(hashLength int) string {
	unixTime := UnixTimestamp()
	return fmt.Sprintf("%s%s", unixTime, secure.RandomString(hashLength-len(unixTime), []rune(secure.RuneAlNumCS)))
}
