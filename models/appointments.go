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

type EditAppointmentResponse struct {
	clientix.EditResponse
}

type AppointmentsListResponse struct {
	clientix.ListResponse
	Items []Appointment `json:"items"`
}

type Appointment struct {
	Id             int32              `json:"id"`
	ClientId       int32              `json:"client_id"`
	ExecutorId     int32              `json:"executor_id"`
	StartDatetime  types.DateTime     `json:"start_datetime"`
	FinishDatetime types.DateTime     `json:"finish_datetime"`
	Status         string             `json:"status"`
	Services       []AppointedService `json:"appointed_services"`
	Created        types.DateTime     `json:"created"`
	Modified       types.DateTime     `json:"modified"`
	Paid           bool               `json:"paid"`
	Urgent         bool               `json:"urgent"`
	IsSale         bool               `json:"is_sale"`
}

type AppointedService struct {
	Id        int32  `json:"id"`
	ServiceId int32  `json:"service_id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
}

//goland:noinspection GoUnusedExportedFunction
func EditAppointment(ctx context.Context, c *clientix.Client, values *netUrl.Values) (ok bool, err error) {
	url := "https://" + c.GetDomain() + "/clientix/Restapi/edit" +
		"/a/" + c.GetAccountId() +
		"/u/" + c.GetUserId() +
		"/t/" + c.GetAccessToken() +
		"/m/Appointments/"

	var data []byte
	if data, err = c.HttpRequest(ctx, "POST", url, values); err != nil {
		return
	}

	var res EditAppointmentResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return
	}
	if !res.IsOk() {
		return false, res.GetError()
	}

	return true, nil
}

//goland:noinspection GoUnusedExportedFunction
func GetAppointmentsList(ctx context.Context, c *clientix.Client, datetime time.Time, offset, limit int) (res *AppointmentsListResponse, err error) {
	url := "https://" + c.GetDomain() + "/clientix/Restapi/list" +
		"/a/" + c.GetAccountId() +
		"/u/" + c.GetUserId() +
		"/t/" + c.GetAccessToken() +
		"/m/Appointments/?date=" + datetime.Format("2006-01-02") +
		"&offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit)

	var data []byte
	if data, err = c.HttpRequest(ctx, "GET", url, nil); err != nil {
		return
	}

	if err = json.Unmarshal(data, &res); err != nil {
		return
	}

	return
}
