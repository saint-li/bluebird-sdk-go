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

type Message struct {
	Cfg *Config
}

// 获取报文数据列表
func (p *Message) GetMessages(query *query.MessagesGetQuery) *result.MessagesGetResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("uuid", query.Uuid)
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("sort", query.Sort)
	params.Set("status", query.Status)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPMessagesGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.MessagesGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.MessagesGetResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.MessagesGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 删除报文
func (p *Message) DeleteMessage(query *query.MessageDeleteQuery) *result.Result {
	var data = gmap.New(true)
	data.Set("id", query.Id)
	data.Set("primary_key", query.PrimaryKey)

	res, err := p.Cfg.HttpClient.SetMethod("delete").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPMessagesDeletePath).SetData(data).HttpRequest()
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
