package types

import (
	"time"
)

type DateTime time.Time

func (t *DateTime) UnmarshalJSON(b []byte) (err error) {
	if string(b) == "null" {
		return
	}

	var datetime time.Time
	if datetime, err = time.ParseInLocation("2006-01-02 15:04:05.999999", string(b[1:len(b)-1]), time.Local); err != nil {
		return
	}
	*t = DateTime(datetime.Local())

	return
}
