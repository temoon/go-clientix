package responses

import (
	"errors"
)

type AddEditResponse struct {
	Status   bool                `json:"status"`
	Messages []string            `json:"messages"`
	Errors   map[string][]string `json:"errors"`
}

func (r *AddEditResponse) IsOk() bool {
	return r.Status && len(r.Errors) == 0
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
