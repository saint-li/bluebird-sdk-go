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

type MacBook struct {
	Cfg *Config
}

// 获取macbook列表
func (p *MacBook) GetMacBooks(query *query.MacBooksGetQuery) *result.MacBooksGetResult {
	params := url.Values{}
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("mac", query.Mac)
	params.Set("name", query.Name)
	params.Set("status", query.Status)
	params.Set("sort", query.Sort)
	params.Set("city_id", strconv.Itoa(int(query.CityId)))
	params.Set("total_count", strconv.Itoa(int(query.TotalCount)))

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TspMacBooksGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.MacBooksGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.MacBooksGetResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.MacBooksGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 获取macbook详情
func (p *MacBook) GetMacBook(query *query.MacBookGetQuery) *result.MacBookGetResult {
	params := url.Values{}
	params.Set("mac", query.Mac)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TspMacBookGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.MacBookGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.MacBookGetResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.MacBookGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 更新macbook
func (p *MacBook) UpdateMacBook(query *query.MacbookUpdateQuery) *result.Result {
	var data = gmap.New(true)
	data.Set("mac", query.Mac)
	data.Set("name", query.Name)
	data.Set("province_id", query.ProvinceId)
	data.Set("city_id", query.CityId)
	data.Set("district_id", query.DistrictId)
	data.Set("lat", query.Lat)
	data.Set("lng", query.Lng)
	data.Set("status", query.Status)
	data.Set("total_count", query.TotalCount)
	data.Set("remark", query.Remark)
	data.Set("street", query.Street)
	data.Set("address", query.Address)

	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TspMacBookUpdatePath).SetData(data).HttpRequest()
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

// 重绘macbook
func (p *MacBook) DrawMacBook(query *query.MacbookDrawQuery) *result.Result {
	var data = gmap.New(true)
	data.Set("mac", query.Mac)
	data.Set("name", query.Name)

	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TspMacBookDrawPath).SetData(data).HttpRequest()
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

// 删除macbook
func (p *MacBook) DeleteMacBook(query *query.MacbookDeleteQuery) *result.Result {
	var data = gmap.New(true)
	data.Set("macaddr", query.MakAddr)
	data.Set("mac", query.Mac)

	res, err := p.Cfg.HttpClient.SetMethod("delete").SetUrl(p.Cfg.HttpClient.GateWay + util.TspMacBookDeletePath).SetData(data).HttpRequest()
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
