package helper

import (
	"time"
)

// HadElapsed function - checks if dateExpiry had elapsed on current time.
func HadElapsed(dateExpiry time.Time) (bool) {
	now := time.Now().UTC()
	return now.After(dateExpiry)
}
