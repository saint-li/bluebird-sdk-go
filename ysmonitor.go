package tspsdk

import (
	"encoding/json"
	"github.com/saint-li/bluebird-sdk-go/result"
	"github.com/saint-li/bluebird-sdk-go/util"
)

type YsMonitor struct {
	Cfg *Config
}

// 获取萤石token
func (p *YsMonitor) TSPYsmonitorAccesstokenPath() *result.YsmonitorTokenGetResult {
	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPYsmonitorAccesstokenGetPath).HttpRequest()
	if err != nil {
		return &result.YsmonitorTokenGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	jsonString, _ := res.ToJsonString()
	var resData = new(result.YsmonitorTokenGetResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.YsmonitorTokenGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}
