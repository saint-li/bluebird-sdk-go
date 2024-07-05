package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
	"net/url"
	"strconv"
)

type Fence struct {
	Cfg *Config
}

// 获取围栏列表
func (p *Fence) GetFences(query *query.FencesGetQuery) *result.FencesGetResult {
	params := url.Values{}
	params.Set("uuid", query.Uuid)
	params.Set("page", strconv.Itoa(query.Page))
	params.Set("limit", strconv.Itoa(query.Limit))
	params.Set("fence_type", strconv.Itoa(int(query.FenceType)))
	params.Set("shape_type", strconv.Itoa(int(query.ShapeType)))

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPFencesGetPath + "?" + params.Encode()).HttpRequest()

	if err != nil {
		return &result.FencesGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.FencesGetResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.FencesGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 创建围栏
func (p *Fence) CreateFence(query *query.FenceCreateQuery) *result.FenceCreateResult {
	var data = make(map[string]interface{})
	data["truename"] = query.Truename
	data["fence_type"] = query.FenceType
	data["shape_type"] = query.ShapeType
	data["collide_type"] = query.CollideType
	data["points"] = query.Points
	data["point_radius"] = query.PointRadius
	data["near_radius"] = query.NearRadius
	data["valid_start"] = query.ValidStart
	data["valid_end"] = query.ValidEnd
	data["valid_week"] = query.ValidWeek
	data["partner_id"] = query.PartnerId
	data["appkey"] = query.AppKey
	data["uuid"] = query.Uuid
	res, err := p.Cfg.HttpClient.SetMethod("post").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPFenceCreatePath).SetData(data).HttpRequest()

	if err != nil {
		return &result.FenceCreateResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	jsonString, _ := res.ToJsonString()

	var resData = new(result.FenceCreateResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.FenceCreateResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 删除围栏
func (p *Fence) DeleteFence(query *query.FenceDeleteQuery) *result.Result {
	var data = make(map[string]interface{})
	data["id"] = query.Id
	res, err := p.Cfg.HttpClient.SetMethod("delete").SetUrl(p.Cfg.HttpClient.GateWay + util.TspFenceDeletePath).SetData(data).HttpRequest()

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
