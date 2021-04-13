package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
	"net/url"
	"strconv"
)

type Temperature struct {
	Cfg *Config
}

// 获取最新体温数据
func (p *Temperature) GetTemperature(query *query.TemperatureGetQuery) (*result.TemperatureResult, error) {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPTemperatureGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return nil, err
	}

	jsonString := res.Export()
	var result = new(result.TemperatureResult)
	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取体温列表
func (p *Temperature) GetTemperatures(query *query.TemperaturesGetQuery) (*result.TemperaturesResult, error) {
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
		return nil, err
	}

	jsonString := res.Export()
	var result = new(result.TemperaturesResult)
	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取体温测量间隔时间
func (p *Temperature) GetTemperatureUpload(query *query.TemperatureUploadGetQuery) (*result.TemperatureUploadResult, error) {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPTemperatureUploadGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return nil, err
	}

	jsonString := res.Export()
	var result = new(result.TemperatureUploadResult)
	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 设置体温测量间隔时间
func (p *Temperature) UpdateTemperatureUpload(param *query.TemperatureUploadSetQuery) (*result.Result, error) {

	var data = make(map[string]interface{})
	data["imei_sn"] = param.ImeiSn
	data["second"] = param.Second

	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPTemperatureUploadSetPath).SetData(data).HttpRequest()
	if err != nil {
		return nil, err
	}

	jsonString := res.Export()

	var result = new(result.Result)
	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
