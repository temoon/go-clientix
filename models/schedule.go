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

const NoData = "No data"

type ScheduleResponse struct {
	Items interface{} `json:"items"`
}

type Schedule map[string][]time.Time

type ScheduleFilterOpts struct {
	StartDatetime  time.Time
	FinishDatetime time.Time
	ExecutorId     int
	ServiceId      int
}

func (r *ScheduleResponse) GetItems() (schedule Schedule, err error) {
	switch items := r.Items.(type) {
	case map[string]interface{}:
		var values []interface{}
		var value string
		var datetime time.Time
		var ok bool

		schedule = make(Schedule, len(items))
		for day, item := range items {
			if values, ok = item.([]interface{}); ok {
				schedule[day] = make([]time.Time, 0, len(values))
				for i := 0; i < len(values); i++ {
					if value, ok = values[i].(string); ok {
						if datetime, err = time.ParseInLocation("2006-01-02 15:04:05", value, types.MskLocation); err != nil {
							continue
						}

						schedule[day] = append(schedule[day], datetime)
					} else {
						continue
					}
				}
			} else {
				continue
			}
		}
	case string:
		if items != NoData {
			err = errors.New(items)
		}
	default:
		err = errors.New("invalid response")
	}

	return
}

//goland:noinspection GoUnusedExportedFunction
func GetAvailableTimes(ctx context.Context, c *clientix.Client, filter *ScheduleFilterOpts) (schedule Schedule, err error) {
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
	} else {
		return nil, errors.New("finish datetime required or invalid")
	}

	if filter.ExecutorId > 0 {
		url += "&executor_id=" + strconv.Itoa(filter.ExecutorId)
	} else {
		return nil, errors.New("executor id required or invalid")
	}

	if filter.ServiceId > 0 {
		url += "&service_id=" + strconv.Itoa(filter.ServiceId)
	} else {
		return nil, errors.New("service id required or invalid")
	}

	var data []byte
	if data, err = c.HttpRequest(ctx, "GET", url, nil); err != nil {
		return
	}

	var res ScheduleResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return
	}

	return res.GetItems()
}
