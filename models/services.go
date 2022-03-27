package models

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/go-clientix"
	"github.com/temoon/go-clientix/types"
)

type ServicesListResponse struct {
	clientix.ListResponse
	Items []Service `json:"items"`
}

type Service struct {
	Id       int            `json:"id"`
	Type     string         `json:"type"`
	Name     string         `json:"name"`
	Price    string         `json:"price"`
	Created  types.DateTime `json:"created"`
	Modified types.DateTime `json:"modified"`
	Archived bool           `json:"archived"`
}

//goland:noinspection GoUnusedExportedFunction
func GetServicesList(ctx context.Context, c *clientix.Client, offset, limit int) (services []Service, err error) {
	url := "https://" + c.GetDomain() + "/clientix/Restapi/list" +
		"/a/" + c.GetAccountId() +
		"/u/" + c.GetUserId() +
		"/t/" + c.GetAccessToken() +
		"/m/Services/" +
		"?offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit)

	var data []byte
	if data, err = c.HttpRequest(ctx, "GET", url, nil); err != nil {
		return
	}

	var res ServicesListResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return
	}

	services = res.Items

	return
}
