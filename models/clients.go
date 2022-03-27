package models

import (
	"context"
	"encoding/json"
	netUrl "net/url"
	"strconv"
	"time"

	"github.com/temoon/go-clientix"
	"github.com/temoon/go-clientix/types"
)

type AddClientResponse struct {
	clientix.AddResponse
	Object Client `json:"object"`
}

type ClientsListResponse struct {
	clientix.ListResponse
	Items []Client `json:"items"`
}

type Client struct {
	Id         int32          `json:"id"`
	Phone      types.Phone    `json:"phone"`
	Created    types.DateTime `json:"created"`
	Modified   types.DateTime `json:"modified"`
	FirstName  string         `json:"first_name"`
	SecondName string         `json:"second_name"`
	Deleted    bool           `json:"deleted"`
	Blocked    bool           `json:"blocked"`
}

//goland:noinspection GoUnusedExportedFunction
func AddClient(ctx context.Context, c *clientix.Client, values *netUrl.Values) (client *Client, err error) {
	url := "https://" + c.GetDomain() + "/clientix/Restapi/add" +
		"/a/" + c.GetAccountId() +
		"/u/" + c.GetUserId() +
		"/t/" + c.GetAccessToken() +
		"/m/Clients/"

	var data []byte
	if data, err = c.HttpRequest(ctx, "POST", url, values); err != nil {
		return
	}

	var res AddClientResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return
	}
	if !res.IsOk() {
		return nil, res.GetError()
	}

	return &res.Object, nil
}

//goland:noinspection GoUnusedExportedFunction
func GetClientsList(ctx context.Context, c *clientix.Client, datetime time.Time, offset, limit int) (res *ClientsListResponse, err error) {
	url := "https://" + c.GetDomain() + "/clientix/Restapi/list" +
		"/a/" + c.GetAccountId() +
		"/u/" + c.GetUserId() +
		"/t/" + c.GetAccessToken() +
		"/m/clients/date/" + datetime.Format("2006-01-02") +
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
