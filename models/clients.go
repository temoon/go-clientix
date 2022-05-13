package models

import (
	"context"
	"encoding/json"
	"errors"
	netUrl "net/url"
	"strconv"
	"time"

	"github.com/temoon/go-clientix"
	"github.com/temoon/go-clientix/responses"
	"github.com/temoon/go-clientix/types"
)

type AddClientResponse struct {
	responses.AddEditResponse
	Status string `json:"status"`
	Object Client `json:"object"`
}

func (r *AddClientResponse) IsOk() bool {
	return r.Status == "ok" && len(r.Errors) == 0
}

type ClientsListResponse struct {
	responses.ListResponse
	Items []Client `json:"items"`
}

type Client struct {
	Id         int            `json:"id"`
	Phone      types.Phone    `json:"phone"`
	Email      string         `json:"email"`
	FirstName  string         `json:"first_name"`
	PatronName string         `json:"patron_name"`
	SecondName string         `json:"second_name"`
	LeadSource string         `json:"lead_source"`
	Created    types.DateTime `json:"created"`
	Modified   types.DateTime `json:"modified"`
	Deleted    bool           `json:"deleted"`
	Blocked    bool           `json:"blocked"`
}

//goland:noinspection GoUnusedExportedFunction
func AddClient(ctx context.Context, c *clientix.Client, client *Client) (err error) {
	url := "https://" + c.GetDomain() + "/clientix/Restapi/add" +
		"/a/" + c.GetAccountId() +
		"/u/" + c.GetUserId() +
		"/t/" + c.GetAccessToken() +
		"/m/Clients/"

	// region Values
	values := netUrl.Values{}

	if client.Phone.IsValid() {
		values.Add("phone", string(client.Phone))
	} else {
		return errors.New("phone required or invalid")
	}

	if client.FirstName != "" {
		values.Add("first_name", client.FirstName)
	} else {
		return errors.New("first name required or invalid")
	}

	if client.PatronName != "" {
		values.Add("patron_name", client.PatronName)
	}

	if client.SecondName != "" {
		values.Add("second_name", client.SecondName)
	}

	if client.Email != "" {
		values.Add("email", client.Email)
	}

	if client.LeadSource != "" {
		values.Add("lead_source", client.LeadSource)
	}
	// endregion

	var data []byte
	if data, err = c.HttpRequest(ctx, "POST", url, &values); err != nil {
		return
	}

	var res AddClientResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return
	}
	if !res.IsOk() {
		return res.GetFirstError()
	}

	client.Id = res.Object.Id

	return
}

//goland:noinspection GoUnusedExportedFunction
func GetClientsList(ctx context.Context, c *clientix.Client, datetime time.Time, offset, limit int) (clients []Client, err error) {
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

	var res ClientsListResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return
	}

	clients = res.Items

	return
}
