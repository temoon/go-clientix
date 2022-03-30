package models

import (
	"context"
	"encoding/json"
	"errors"
	netUrl "net/url"
	"strconv"
	"strings"
	"time"

	"github.com/temoon/go-clientix"
	"github.com/temoon/go-clientix/types"
)

type AddAppointmentResponse struct {
	clientix.AddResponse
	Object Client `json:"object"`
}

type EditAppointmentResponse struct {
	clientix.EditResponse
}

type AppointmentsListResponse struct {
	clientix.ListResponse
	Items []Appointment `json:"items"`
}

type Appointment struct {
	Id             int                     `json:"id"`
	ClientId       int                     `json:"client_id"`
	ExecutorId     int                     `json:"executor_id"`
	StartDatetime  types.DateTime          `json:"start_datetime"`
	FinishDatetime types.DateTime          `json:"finish_datetime"`
	Status         types.AppointmentStatus `json:"status"`
	Services       []AppointedService      `json:"appointed_services"`
	Created        types.DateTime          `json:"created"`
	Modified       types.DateTime          `json:"modified"`
	Paid           bool                    `json:"paid"`
	Urgent         bool                    `json:"urgent"`
	IsSale         bool                    `json:"is_sale"`
}

type AppointedService struct {
	Id        int    `json:"id"`
	ServiceId int    `json:"service_id"`
	Name      string `json:"name"`
	Type      string `json:"type,omitempty"`
}

//goland:noinspection GoUnusedExportedFunction
func AddAppointment(ctx context.Context, c *clientix.Client, appointment *Appointment) (err error) {
	url := "https://" + c.GetDomain() + "/clientix/Restapi/add" +
		"/a/" + c.GetAccountId() +
		"/u/" + c.GetUserId() +
		"/t/" + c.GetAccessToken() +
		"/m/Appointments/"

	// region Values
	values := netUrl.Values{}

	if appointment.ClientId > 0 {
		values.Add("client_id", strconv.Itoa(appointment.ClientId))
	} else {
		return errors.New("client id required or invalid")
	}

	if appointment.ExecutorId > 0 {
		values.Add("executor_id", strconv.Itoa(appointment.ExecutorId))
	} else {
		return errors.New("executor id required or invalid")
	}

	if appointment.Status.IsValid() {
		values.Add("status", string(appointment.Status))
	} else {
		return errors.New("status required or invalid")
	}

	startDatetime := time.Time(appointment.StartDatetime).Round(15 * time.Minute)
	if !startDatetime.IsZero() {
		values.Add("start_datetime", startDatetime.Format("2006-01-02"))
		values.Add("start_time", startDatetime.Format("15:04:05"))
	} else {
		return errors.New("start datetime required or invalid")
	}

	finishDatetime := time.Time(appointment.FinishDatetime).Round(15 * time.Minute)
	if finishDatetime.After(startDatetime) {
		values.Add("finish_datetime", finishDatetime.Format("2006-01-02"))
		values.Add("finish_time", finishDatetime.Format("15:04:05"))
	} else {
		return errors.New("finish datetime required or invalid")
	}

	if servicesCount := len(appointment.Services); servicesCount != 0 {
		serviceIds := make([]string, servicesCount)
		for i := 0; i < servicesCount; i++ {
			serviceIds[i] = strconv.Itoa(appointment.Services[i].ServiceId)
		}

		values.Add("appointed_services", "["+strings.Join(serviceIds, ",")+"]")
	}
	// endregion

	var data []byte
	if data, err = c.HttpRequest(ctx, "POST", url, &values); err != nil {
		return
	}

	var res AddAppointmentResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return
	}
	if !res.IsOk() {
		return res.GetError()
	}

	appointment.Id = res.Object.Id

	return
}

//goland:noinspection GoUnusedExportedFunction
func EditAppointment(ctx context.Context, c *clientix.Client, appointment *Appointment) (err error) {
	url := "https://" + c.GetDomain() + "/clientix/Restapi/edit" +
		"/a/" + c.GetAccountId() +
		"/u/" + c.GetUserId() +
		"/t/" + c.GetAccessToken() +
		"/m/Appointments/"

	// region Values
	values := netUrl.Values{}

	if appointment.Id > 0 {
		values.Add("id", strconv.Itoa(appointment.Id))
	} else {
		return errors.New("id required or invalid")
	}

	if appointment.ClientId > 0 {
		values.Add("client_id", strconv.Itoa(appointment.ClientId))
	}

	if appointment.ExecutorId > 0 {
		values.Add("executor_id", strconv.Itoa(appointment.ExecutorId))
	}

	if appointment.Status.IsValid() {
		values.Add("status", string(appointment.Status))
	}

	startDatetime := time.Time(appointment.StartDatetime).Round(15 * time.Minute)
	if !startDatetime.IsZero() {
		values.Add("start_datetime", startDatetime.Format("2006-01-02 15:04:05"))
	}

	finishDatetime := time.Time(appointment.FinishDatetime).Round(15 * time.Minute)
	if finishDatetime.After(startDatetime) {
		values.Add("finish_datetime", finishDatetime.Format("2006-01-02 15:04:05"))
	}
	// endregion

	var data []byte
	if data, err = c.HttpRequest(ctx, "POST", url, &values); err != nil {
		return
	}

	var res EditAppointmentResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return
	}
	if !res.IsOk() {
		return res.GetError()
	}

	return
}

//goland:noinspection GoUnusedExportedFunction
func GetAppointmentsList(ctx context.Context, c *clientix.Client, datetime time.Time, offset, limit int) (appointments []Appointment, err error) {
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

	var res AppointmentsListResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return
	}

	appointments = res.Items

	return
}
