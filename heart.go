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

type Heart struct {
	Cfg *Config
}

// 获取最新心率数据
func (p *Heart) GetHeart(query *query.HeartGetQuery) *result.HeartResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPHeartGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.HeartResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.HeartResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.HeartResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 获取心率列表
func (p *Heart) GetHearts(query *query.HeartsGetQuery) *result.HeartsResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("sort", query.Sort)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPHeartsGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.HeartsResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.HeartsResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.HeartsResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 获取心率测量间隔时间
func (p *Heart) GetHeartUpload(query *query.HeartUploadGetQuery) *result.HeartUploadResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPHeartUploadGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.HeartUploadResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.HeartUploadResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.HeartUploadResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 设置心率测量间隔时间
func (p *Heart) UpdateHeartUpload(param *query.HeartUploadSetQuery) *result.Result {

	var data = make(map[string]interface{})
	data["imei_sn"] = param.ImeiSn
	data["second"] = param.Second

	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPHeartUploadSetPath).SetData(data).HttpRequest()
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

// 删除心率
func (p *Heart) DeleteHeart(query *query.HeartDeleteQuery) *result.Result {
	var data = gmap.New(true)
	data.Set("id", query.Id)
	data.Set("primary_key", query.PrimaryKey)

	res, err := p.Cfg.HttpClient.SetMethod("delete").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPHeartDeletePath).SetData(data).HttpRequest()
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
