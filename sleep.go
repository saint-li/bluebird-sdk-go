package tspsdk

import (
	"encoding/json"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
	"net/url"
	"strconv"
)

type Sleep struct {
	Cfg *Config
}

// 获取睡眠列表
func (p *Sleep) GetSleeps(query *query.SleepsGetQuery) *result.SleepsResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("sort", query.Sort)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPSleepsGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.SleepsResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.SleepsResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.SleepsResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 获取最新睡眠数据
func (p *Sleep) GetSleep(query *query.SleepGetQuery) *result.SleepResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPSleepGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.SleepResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.SleepResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.SleepResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 删除睡眠数据
func (p *Sleep) DeleteSleep(query *query.SleepDeleteQuery) *result.Result {
	var data = gmap.New(true)
	data.Set("id", query.Id)
	data.Set("primary_key", query.PrimaryKey)

	res, err := p.Cfg.HttpClient.SetMethod("delete").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPSleepDeletePath).SetData(data).HttpRequest()
	if err != nil {
		return &result.Result{
			Status:  1,
			Message: err.Error(),
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.Result{
			Status:  1,
			Message: err.Error(),
		}
	}
	return resData
}
