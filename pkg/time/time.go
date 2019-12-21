package time

import (
	"fmt"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%s"`, time.Time(t).Format(timeLayout))
	return []byte(s), nil
}

func (t *JsonTime) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) <= 1 {
		dateTime, _ := time.Parse(timeLayout, "0000-00-00 00:00:00")
		*t = JsonTime(dateTime)
		return nil
	}
	// 因为实际接收到值是""2018-11-25 20:04:51""格式的，所以这里去除前后各一个"号
	str := string(data[1 : len(data)-1])
	st, err := time.Parse(timeLayout, str)
	if err == nil {
		*t = JsonTime(st)
	} else {
		return err
	}
	return nil
}
