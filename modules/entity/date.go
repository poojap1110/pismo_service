package entity

import (
	"os"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
	"bitbucket.org/matchmove/go-memcached-model"
)

const (
	// DefaultTimestampFormat ...
	DefaultTimestampFormat = "2006-01-02 15:04:05"

	// DefaultDateFormat ...
	DefaultDateFormat = "2006-01-02"

	// DateTimeMicroFormat ...
	DateTimeMicroFormat = "20060102150405"

	// FedDateFormat ...
	FedDateFormat = "02-01-2006"

	RFC8601 = "2006-01-02T15:04:05"
)

// Date ...
type Date struct {
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// NewDate ...
func NewDate(CreatedAt string, UpdatedAt string) Date {
	if "0000-00-00 00:00:00" == UpdatedAt || strings.Contains(UpdatedAt, "1970-01-01") {
		UpdatedAt = ""
	}

	return Date{CreatedAt: CreatedAt, UpdatedAt: UpdatedAt}
}

// GetFirstDayOfWeek function - Gets the first day of a week basing from datetime object received.
func GetFirstDayOfWeek(date time.Time) time.Time {
	var (
		startDayInt time.Weekday
		err         error
	)

	temp, err := strconv.Atoi(os.Getenv(constant.EnvWeekStartInt))
	startDayInt = time.Weekday(temp)

	if err != nil {
		panic(err)
	}

	for date.Weekday() != startDayInt {
		date = date.AddDate(0, 0, -1)
	}

	return date
}

// GetStartAndEndDatesOfQuarter function - Get the start and end dates of current quarter based from datetime object received.
func GetStartAndEndDatesOfQuarter(date time.Time) (startDateTime time.Time, endDateTime time.Time, err error) {
	month, _ := strconv.Atoi(date.Format("1"))

	switch time.Month(month) {
	case time.January, time.February, time.March:
		startDateTime, err = time.Parse(gorm.SQLDatetime, date.Format("2006")+"-01-01 00:00:00")
		endDateTime, err = time.Parse(gorm.SQLDatetime, date.Format("2006")+"-03-31 23:59:59")
	case time.April, time.May, time.June:
		startDateTime, err = time.Parse(gorm.SQLDatetime, date.Format("2006")+"-04-01 00:00:00")
		endDateTime, err = time.Parse(gorm.SQLDatetime, date.Format("2006")+"-06-30 23:59:59")
	case time.July, time.August, time.September:
		startDateTime, err = time.Parse(gorm.SQLDatetime, date.Format("2006")+"-07-01 00:00:00")
		endDateTime, err = time.Parse(gorm.SQLDatetime, date.Format("2006")+"-09-30 23:59:59")
	default: // handles Oct, Nov, Dec
		startDateTime, err = time.Parse(gorm.SQLDatetime, date.Format("2006")+"-10-01 00:00:00")
		endDateTime, err = time.Parse(gorm.SQLDatetime, date.Format("2006")+"-12-31 23:59:59")
	}

	return
}

// GetStartAndEndDatesOfYear function - Get the start and end dates of current year based from datetime object received.
func GetStartAndEndDatesOfYear(date time.Time) (startDateTime time.Time, endDateTime time.Time, err error) {
	startDateTime, err = time.Parse(gorm.SQLDatetime, date.Format("2006")+"-01-01 00:00:00")
	endDateTime, err = time.Parse(gorm.SQLDatetime, date.Format("2006")+"-12-31 23:59:59")

	return
}

// GetDateEndings function - Get the start and end timestamps
func GetDateEndings(date time.Time) (startDateTime time.Time, endDateTime time.Time, err error) {
	startDateTime, err = time.Parse(DefaultTimestampFormat, date.Format(DefaultDateFormat)+" 00:00:00")
	endDateTime, err = time.Parse(DefaultTimestampFormat, date.Format(DefaultDateFormat)+" 23:59:59")

	return
}
