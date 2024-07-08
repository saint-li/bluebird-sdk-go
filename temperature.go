package tspsdk

import (
	"encoding/json"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/saint-li/bluebird-sdk-go/query"
	"github.com/saint-li/bluebird-sdk-go/result"
	"github.com/saint-li/bluebird-sdk-go/util"
	"net/url"
	"strconv"
)

type Temperature struct {
	Cfg *Config
}

// 获取最新体温数据
func (p *Temperature) GetTemperature(query *query.TemperatureGetQuery) *result.TemperatureResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPTemperatureGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.TemperatureResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.TemperatureResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.TemperatureResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 获取体温列表
func (p *Temperature) GetTemperatures(query *query.TemperaturesGetQuery) *result.TemperaturesResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("sort", query.Sort)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPTemperaturesGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.TemperaturesResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.TemperaturesResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.TemperaturesResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 获取体温测量间隔时间
func (p *Temperature) GetTemperatureUpload(query *query.TemperatureUploadGetQuery) *result.TemperatureUploadResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPTemperatureUploadGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.TemperatureUploadResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.TemperatureUploadResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.TemperatureUploadResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 设置体温测量间隔时间
func (p *Temperature) UpdateTemperatureUpload(param *query.TemperatureUploadSetQuery) *result.Result {

	var data = make(map[string]interface{})
	data["imei_sn"] = param.ImeiSn
	data["second"] = param.Second

	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPTemperatureUploadSetPath).SetData(data).HttpRequest()
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

// 删除体温数据
func (p *Temperature) DeleteTemperature(query *query.TemperatureDeleteQuery) *result.Result {
	var data = gmap.New(true)
	data.Set("id", query.Id)
	data.Set("primary_key", query.PrimaryKey)

	res, err := p.Cfg.HttpClient.SetMethod("delete").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPTemperatureDeletePath).SetData(data).HttpRequest()
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
