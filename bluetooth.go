package tspsdk

import (
	"encoding/json"
	"github.com/xu5g/bluebird-sdk-go/query"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
)

type Bluetooth struct {
	Cfg *Config
}

// 蓝牙配对
func (p *Bluetooth) SetBluetoothLink(param *query.TspBluetoothLinkRequest) *result.Result {
	var data = make(map[string]interface{})

	data["imei_sn"] = param.ImeiSn
	data["status"] = param.Status
	data["mac"] = param.Mac

	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TspBluetoothLinkPath).SetData(data).HttpRequest()
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
