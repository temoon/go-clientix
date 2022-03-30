package models

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/go-clientix"
	"github.com/temoon/go-clientix/responses"
	"github.com/temoon/go-clientix/types"
)

type UsersListResponse struct {
	responses.ListResponse
	Items []User `json:"items"`
}

type User struct {
	Id       int            `json:"id"`
	Name     string         `json:"name"`
	Phone    types.Phone    `json:"phone"`
	Email    string         `json:"email"`
	Created  types.DateTime `json:"created"`
	Modified types.DateTime `json:"modified"`
	Archived bool           `json:"archived"`
}

//goland:noinspection GoUnusedExportedFunction
func GetUsersList(ctx context.Context, c *clientix.Client, offset, limit int) (users []User, err error) {
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

	var res UsersListResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return
	}

	users = res.Items

	return
}
