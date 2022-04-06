package types

import (
	"time"
)

var MskLocation *time.Location

type DateTime time.Time

func init() {
	var err error
	if MskLocation, err = time.LoadLocation("Europe/Moscow"); err != nil {
		panic(err)
	}
}

func (t *DateTime) UnmarshalJSON(b []byte) (err error) {
	if string(b) == "null" {
		return
	}

	var datetime time.Time
	if datetime, err = time.ParseInLocation("2006-01-02 15:04:05.999999", string(b[1:len(b)-1]), MskLocation); err != nil {
		return
	}
	*t = DateTime(datetime.Local())

	return
}
