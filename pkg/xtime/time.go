package xtime

import (
	"apiserver-gin/pkg/constant"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%s"`, time.Time(t).Format(constant.TimeLayout))
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
	// 因为实际接收到值是"2018-11-25 20:04:51"格式的，所以这里去除前后各一个"号
	str := string(data[1 : len(data)-1])
	st, err := time.Parse(constant.TimeLayout, str)
	if err == nil {
		*t = Time(st)
	} else {
		return err
	}
	return nil
}

func (t Time) Value() (driver.Value, error) {
	tm := time.Time(t)
	//return tm.Format(constant.TimeLayout), nil
	return tm, nil
}

func (t *Time) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch st := value.(type) {
	case time.Time:
		*t = Time(st)
	case string:
		tm, err := time.Parse(constant.TimeLayout, st)
		if err != nil {
			return err
		}
		*t = Time(tm)
	}
	return nil
}
