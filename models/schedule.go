package models

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/temoon/go-clientix"
	"github.com/temoon/go-clientix/types"
)

type ScheduleResponse struct {
	Items Schedule `json:"items"`
}

type Schedule map[string]types.DateTime

type ScheduleFilterOpts struct {
	StartDatetime  time.Time
	FinishDatetime time.Time
	ExecutorId     int
	ServiceId      int
}

//goland:noinspection GoUnusedExportedFunction
func GetAvailableTime(ctx context.Context, c *clientix.Client, filter *ScheduleFilterOpts) (schedule Schedule, err error) {
	url := "https://" + c.GetDomain() + "/clientix/Restapi/list" +
		"/a/" + c.GetAccountId() +
		"/u/" + c.GetUserId() +
		"/t/" + c.GetAccessToken() +
		"/m/availableTimes"

	if !filter.StartDatetime.IsZero() {
		url += "?start_day=" + filter.StartDatetime.Format("2006-01-02")
	} else {
		return nil, errors.New("start datetime required or invalid")
	}

	if filter.FinishDatetime.After(filter.StartDatetime) {
		url += "&finish_day=" + filter.FinishDatetime.Format("2006-01-02")
	}

	if filter.ExecutorId > 0 {
		url += "&executor_id=" + strconv.Itoa(filter.ExecutorId)
	}

	if filter.ServiceId > 0 {
		url += "&service_id=" + strconv.Itoa(filter.ServiceId)
	}

	var data []byte
	if data, err = c.HttpRequest(ctx, "GET", url, nil); err != nil {
		return
	}

	var res ScheduleResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return
	}

	schedule = res.Items

	return
}
