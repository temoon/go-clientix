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
	Items interface{} `json:"items"`
}

type Schedule map[string][]types.DateTime

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

	switch items := res.Items.(type) {
	case map[string]interface{}:
		var values []interface{}
		var value string
		var datetime time.Time
		var ok bool

		schedule = make(Schedule, len(items))
		for day, item := range items {
			if values, ok = item.([]interface{}); ok {
				schedule[day] = make([]types.DateTime, 0, len(values))
				for i := 0; i < len(values); i++ {
					if value, ok = values[i].(string); ok {
						if datetime, err = time.Parse("2006-01-02 15:04:05", value); err != nil {
							continue
						}

						schedule[day] = append(schedule[day], types.DateTime(datetime))
					} else {
						continue
					}
				}
			} else {
				continue
			}
		}
	case string:
		schedule = nil
		if items != "No data" {
			err = errors.New(items)
		}
	default:
		schedule = nil
		err = errors.New("no data")
	}

	return
}
