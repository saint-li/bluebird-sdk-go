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

type MacList struct {
	Cfg *Config
}

// 添加maclist
func (p *MacList) MacListCreate(query *query.MacListCreateQuery) *result.Result {
	var data = gmap.New(true)
	data.Set("name", query.Name)
	data.Set("lng", query.Lng)
	data.Set("lat", query.Lat)
	data.Set("mac", query.Mac)
	data.Set("signal", query.Signal)
	data.Set("remark", query.Remark)
	data.Set("status", query.Status)

	res, err := p.Cfg.HttpClient.SetMethod("post").SetUrl(p.Cfg.HttpClient.GateWay + util.TspMacListCreatePath).SetData(data).HttpRequest()
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

// 获取mac列表
func (p *MacList) GetMacLists(query *query.MacListsGetQuery) *result.MacListsGetResult {
	params := url.Values{}
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("mac", query.Mac)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TspMacListsGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.MacListsGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.MacListsGetResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.MacListsGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	return resData
}

// 删除maclist
func (p *MacList) DeleteMacList(query *query.MacListDeleteQuery) *result.Result {
	var data = gmap.New(true)
	data.Set("id", query.Id)
	data.Set("macaddr", query.MacAddr)
	data.Set("mac", query.Mac)

	res, err := p.Cfg.HttpClient.SetMethod("delete").SetUrl(p.Cfg.HttpClient.GateWay + util.TspMacListDeletePath).SetData(data).HttpRequest()
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
