package types

import (
	"strings"
)

type Phone string

func (t *Phone) UnmarshalJSON(b []byte) (err error) {
	if string(b) == "null" {
		return
	}

	*t = ParsePhone(string(b[1 : len(b)-1]))

	return
}

func (t Phone) IsValid() bool {
	if len(t) == 0 {
		return false
	}

	for i := 0; i < len(t); i++ {
		if t[i] < '0' || t[i] > '9' {
			return false
		}
	}

	return true
}

func ParsePhone(data string) Phone {
	builder := strings.Builder{}
	for i := 0; i < len(data); i++ {
		if data[i] >= '0' && data[i] <= '9' {
			builder.WriteByte(data[i])
		}
	}

	return Phone(builder.String())
}
