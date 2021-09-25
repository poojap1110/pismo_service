package logger

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	logs "bitbucket.org/matchmove/fmt-logs"
	"bitbucket.org/matchmove/integration-svc-aub/modules/constant"
	"bitbucket.org/matchmove/integration-svc-aub/modules/container"
	"github.com/go-errors/errors"
	"github.com/willf/pad"
)

type logger struct {
	sync.Mutex
	name string
}

const (
	LoggerInstanceKey = "Logger"
	LogGeneral        = "GENERAL"

	year      = "2006"
	month     = "01"
	day       = "02"
	separator = "/"
)

// GetInstance function ...
func GetInstance(c *container.Container, keyName string) (l *logger) {
	if l, ok := c.Get(LoggerInstanceKey).(map[string]*logger); ok == true {
		if l[keyName] != nil {
			return l[keyName]
		}

		l[keyName] = &logger{name: keyName}
		c.StoreToGlobal(LoggerInstanceKey, l)
	}

	l = &logger{name: keyName}
	contValues := c.Get(LoggerInstanceKey)

	if contValues != nil {
		contValues.(map[string]*logger)[keyName] = l
	} else {
		contValues = map[string]*logger{keyName: l}
	}

	c.StoreToGlobal(LoggerInstanceKey, contValues)

	return
}

// WriteLog method ...
func (me *logger) WriteLog(c []byte) error {
	me.Lock()
	defer me.Unlock()

	timezone := os.Getenv("LOG_TIMEZONE")

	if timezone == "" {
		timezone = "Asia/Manila"
	}

	loc, _ := time.LoadLocation(timezone)
	datetime, _ := time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))

	filepath := os.Getenv(constant.EnvAppLogFolder) + separator + datetime.Format(year) + separator + datetime.Format(month) + separator +
		datetime.Format(day) + separator
	fullpath := filepath + me.name

	if _, err := os.OpenFile(fullpath, os.O_RDONLY|os.O_CREATE, 0744); err != nil {
		err = os.MkdirAll(filepath, 0755)

		if err != nil {
			panic(errors.New("Failed to create Log file on " + filepath + " because of " + err.Error()))
		}

		os.Create(fullpath)
	}

	file, err := os.OpenFile(fullpath, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.Write([]byte("--- " + fmt.Sprintf("%d", datetime.Unix()) + " ---\n" + string(c) + "\n")); err != nil {
		panic(errors.New("Failed to write in Log file on " + filepath + " because of " + err.Error()))
		return errors.New("Failed to write in Log file on " + filepath + " because of " + err.Error())
	}

	return nil
}

// Dump method ...
func (me *logger) Dump(log *logs.Log) {
	var (
		entries      = log.Entries()
		line, format string
		params       []interface{}
		entireLog    string
	)

	for k, v := range entries {
		format = "%s\t%s"
		params = []interface{}{
			v.Time.Format(logs.DateFormat),
			v.Level,
		}

		if log.GetIdentify() != "" {
			params = append(params, log.GetIdentify())
			format = format + "\t%s"
		}

		if log.GetCount() {
			params = append(params, pad.Left(strconv.Itoa(k), 3, "0"))
			format = format + "  %s"
		}

		params = append(params, v.Message)
		format = format + "  %s\n"
		line = fmt.Sprintf(format, params...)

		entireLog = entireLog + line
	}

	me.WriteLog([]byte(entireLog))
}
