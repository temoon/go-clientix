package responses

import (
	"errors"
)

type AddEditResponse struct {
	Status   bool     `json:"status"`
	Messages []string `json:"messages"`
	Errors   []string `json:"errors"`
}

func (r *AddEditResponse) IsOk() bool {
	return r.Status
}

func (r *AddEditResponse) GetFirstMessage() string {
	if len(r.Messages) != 0 {
		return r.Messages[0]
	}

	return ""
}

func (r *AddEditResponse) GetFirstError() error {
	if len(r.Errors) != 0 {
		return errors.New(r.Errors[0])
	}

	return nil
}
