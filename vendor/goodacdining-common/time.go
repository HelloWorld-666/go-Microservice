package common

import (
	"time"
)

const (
	DATE_LAYOUT             = "2006-01-02"
	TIME_LAYOUT             = "15:04:05.000"
	DATETIME_LAYOUT         = "2006-01-02 15:04:05.000"
	DATETIME_FOR_WEB_LAYOUT = "2006-01-02 15:04:05.000"
)

func ParseDateTime(strDateTime string) (time.Time, error) {
	return time.ParseInLocation(DATETIME_LAYOUT, strDateTime, time.Local)
}
