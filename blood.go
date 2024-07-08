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

type Blood struct {
	Cfg *Config
}

// 获取最新血压数据
func (p *Blood) GetBlood(query *query.BloodGetQuery) *result.BloodGetResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPBloodGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.BloodGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.BloodGetResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.BloodGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 获取血压列表
func (p *Blood) GetBloods(query *query.BloodsGetQuery) *result.BloodsGetResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("sort", query.Sort)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPBloodsGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.BloodsGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.BloodsGetResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.BloodsGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 获取血压测量间隔时间
func (p *Blood) GetBloodUpload(query *query.BloodUploadGetQuery) *result.BloodUploadResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPBloodUploadGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.BloodUploadResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.BloodUploadResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.BloodUploadResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 设置血压测量间隔时间
func (p *Blood) UpdateBloodUpload(param *query.BloodUploadSetQuery) *result.Result {

	var data = make(map[string]interface{})
	data["imei_sn"] = param.ImeiSn
	data["second"] = param.Second

	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPBloodUploadSetPath).SetData(data).HttpRequest()
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

// 删除血压数据
func (p *Blood) DeleteBlood(query *query.BloodDeleteQuery) *result.Result {
	var data = gmap.New(true)
	data.Set("id", query.Id)
	data.Set("primary_key", query.PrimaryKey)

	res, err := p.Cfg.HttpClient.SetMethod("delete").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPBloodDeletePath).SetData(data).HttpRequest()
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
