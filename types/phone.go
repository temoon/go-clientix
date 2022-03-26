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
