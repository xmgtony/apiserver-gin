package xtime

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

const (
	DateTime   = "2006-01-02 15:04:05"
	DateTimeMs = "2006-01-02 15:04:05.000"
	DateOny    = "2006-01-02"
	TimeOny    = "15:04:05"
)

var timeLayouts = []string{
	DateTime,
	DateTimeMs,
	DateOny,
	TimeOny,
}

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%s"`, time.Time(t).Format(DateTime))
	return []byte(s), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	// TODO(https://go.dev/issue/47353): Properly unescape a JSON string.
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("[xtime.Time].UnmarshalJSON: input is not a JSON string")
	}
	// 因为实际接收到值是"2018-11-25 20:04:51"格式的json字符串，所以这里去除前后各一个"号
	str := string(data[1 : len(data)-1])
	for _, layout := range timeLayouts {
		st, err := time.ParseInLocation(layout, str, time.Local)
		if err == nil {
			*t = Time(st)
			return nil
		}
	}
	return errors.New("[xtime.Time].UnmarshalJSON: time parse error")
}

func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t *Time) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch st := value.(type) {
	case time.Time:
		*t = Time(st)
	case string:
		for _, layout := range timeLayouts {
			st, err := time.ParseInLocation(layout, st, time.Local)
			if err == nil {
				*t = Time(st)
				return nil
			}
		}
	}
	return nil
}
