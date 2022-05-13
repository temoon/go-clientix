package responses

import (
	"errors"
)

type AddEditResponse struct {
	Status   interface{}         `json:"status"`
	Messages []string            `json:"messages"`
	Errors   map[string][]string `json:"errors"`
}

func (r *AddEditResponse) IsOk() bool {
	return r.GetStatus() && len(r.Errors) == 0
}

func (r *AddEditResponse) GetStatus() bool {
	switch status := r.Status.(type) {
	case string:
		return status == "ok"
	case bool:
		return status
	}

	return false
}

func (r *AddEditResponse) GetFirstMessage() string {
	if len(r.Messages) != 0 {
		return r.Messages[0]
	}

	return ""
}

func (r *AddEditResponse) GetFirstError() error {
	for _, values := range r.Errors {
		for _, value := range values {
			return errors.New(value)
		}
	}

	return nil
}
