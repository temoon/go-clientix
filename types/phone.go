package types

import (
	"strings"
)

type Phone string

func (t *Phone) UnmarshalJSON(b []byte) (err error) {
	if string(b) == "null" {
		return
	}

	cleanedPhone := strings.Builder{}
	cleanedPhone.WriteByte('+')

	for i := 1; i < len(b)-1; i++ {
		if b[i] >= '0' && b[i] <= '9' {
			cleanedPhone.WriteByte(b[i])
		}
	}

	*t = Phone(cleanedPhone.String())

	return
}

func (t Phone) IsValid() bool {
	if len(t) < 2 || t[0] != '+' {
		return false
	}

	for _, c := range t[1:] {
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
}
