package clientix

import (
	"errors"
)

type AddResponse struct {
	Status   string   `json:"status"`
	Messages []string `json:"messages"`
}

func (r *AddResponse) IsOk() bool {
	return r.Status == "ok"
}

func (r *AddResponse) GetError() error {
	if len(r.Messages) != 0 {
		return errors.New(r.Messages[0])
	}

	return errors.New("unknown error")
}

type EditResponse struct {
	Status   bool     `json:"status"`
	Messages []string `json:"messages"`
}

func (r *EditResponse) IsOk() bool {
	return r.Status
}

func (r *EditResponse) GetError() error {
	if len(r.Messages) != 0 {
		return errors.New(r.Messages[0])
	}

	return errors.New("unknown error")
}

type ListResponse struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
