package models

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/go-clientix"
	"github.com/temoon/go-clientix/types"
)

type UsersResponse struct {
	clientix.ListResponse
	Items []User `json:"items"`
}

type User struct {
	Id       int32          `json:"id"`
	Name     string         `json:"name"`
	Phone    types.Phone    `json:"phone"`
	Email    string         `json:"email"`
	Created  types.DateTime `json:"created"`
	Modified types.DateTime `json:"modified"`
	Archived bool           `json:"archived"`
}

func GetUsers(ctx context.Context, c *clientix.Client, offset, limit int) (res *UsersResponse, err error) {
	url := "https://" + c.GetDomain() + "/clientix/Restapi/list" +
		"/a/" + c.GetAccountId() +
		"/u/" + c.GetUserId() +
		"/t/" + c.GetAccessToken() +
		"/m/Users/" +
		"?offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit)

	var data []byte
	if data, err = c.HttpRequest(ctx, "GET", url, nil); err != nil {
		return
	}

	if err = json.Unmarshal(data, &res); err != nil {
		return
	}

	return
}
