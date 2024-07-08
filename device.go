package tspsdk

import (
	"encoding/json"
	"fmt"
	"github.com/saint-li/bluebird-sdk-go/query"
	"github.com/saint-li/bluebird-sdk-go/result"
	"github.com/saint-li/bluebird-sdk-go/util"
	"net/url"
	"strconv"
)

type Device struct {
	Cfg *Config
}

// 设备详情
func (p *Device) GetDevice(query *query.DeviceGetQuery) *result.DeviceGetResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("attence_sn", query.AttenceSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.DeviceGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.DeviceGetResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.DeviceGetResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 更新设备信息
func (p *Device) DeviceUpdate(query *query.DeviceUpdateQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["truename"] = query.Truename
	data["mobile"] = query.Mobile
	data["appkey"] = query.AppKey
	data["model_id"] = query.ModelId
	data["ic_card_sn"] = query.IccardSn
	data["attence_sn"] = query.AttenceSn
	data["partner_id"] = query.PartnerId

	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceUpdatePath).SetData(data).HttpRequest()
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

// 创建设备
func (p *Device) DeviceCreate(query *query.DeviceCreateQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["model_id"] = query.ModelId
	data["attence_sn"] = query.AttenceSn
	data["appkey"] = query.AppKey
	data["iccard_sn"] = query.IccardSn
	data["locate_upload"] = query.LocateUpload
	data["partner_id"] = query.PartnerId
	data["engine"] = query.Engine

	res, err := p.Cfg.HttpClient.SetMethod("post").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceCreatePath).SetData(data).HttpRequest()
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

// 获取设备列表
func (p *Device) GetDevices(query *query.DevicesGetQuery) *result.DevicesResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("attence_sn", query.AttenceSn)
	params.Set("mobile", query.Mobile)
	params.Set("is_online", query.IsOnline)
	params.Set("uuid", query.Uuid)
	params.Set("page", strconv.Itoa(int(query.Page)))
	params.Set("limit", strconv.Itoa(int(query.Limit)))
	params.Set("model_id", strconv.FormatInt(query.ModelId, 10))
	params.Set("appkey", strconv.FormatInt(query.AppKey, 10))
	params.Set("partner_id", strconv.FormatInt(query.PartnerId, 10))
	params.Set("start_active_time", query.StartActiveTime)
	params.Set("end_active_time", query.EndActiveTime)
	params.Set("start_firstlink_time", query.StartFirstlinkTime)
	params.Set("end_firstlink_time", query.EndFirstlinkTime)
	params.Set("start_time", query.StartTime)
	params.Set("end_time", query.EndTime)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDevicesGetPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.DevicesResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.DevicesResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.DevicesResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 下发定位指令到终端
func (p *Device) SendLocate(query *query.DeviceLocateQuery) *result.Result {

	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceLocatePath + "?" + params.Encode()).HttpRequest()

	if err != nil {
		return &result.Result{
			Status:  1,
			Message: err.Error(),
		}
	}
	jsonString, _ := res.ToJsonString()
	fmt.Println(jsonString)

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

// 判断设备是否在线
func (p *Device) GetDeviceIsOnline(query *query.DeviceIsOnlineQuery) *result.DeviceIsOnlineResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceOnlinePath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.DeviceIsOnlineResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.DeviceIsOnlineResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.DeviceIsOnlineResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 透传报文
func (p *Device) SendMessage(query *query.DeviceMessageQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["message"] = query.Message
	res, err := p.Cfg.HttpClient.SetMethod("post").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceMessagePath).SetData(data).HttpRequest()

	fmt.Println(res, err)
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

// 获取设备功能清单
func (p *Device) GetDeviceModules(query *query.DeviceModulesQuery) *result.DeviceModulesResult {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceModulesPath + "?" + params.Encode()).HttpRequest()
	if err != nil {
		return &result.DeviceModulesResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}

	jsonString, _ := res.ToJsonString()
	var resData = new(result.DeviceModulesResult)
	err = json.Unmarshal([]byte(jsonString), resData)
	if err != nil {
		return &result.DeviceModulesResult{
			Result: result.Result{
				Status:  1,
				Message: err.Error(),
			},
		}
	}
	return resData
}

// 设备绑定
func (p *Device) BindDevice(query *query.DeviceBindQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["mobile"] = query.Mobile
	data["uuid"] = query.Uuid
	data["truename"] = query.TrueName
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceBindPath).SetData(data).HttpRequest()

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

// 设备解绑
func (p *Device) UnBindDevice(query *query.DeviceUnBindQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceUnBindPath).SetData(data).HttpRequest()

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

// 下发设备寻找指令
func (p *Device) SendFindDevice(query *query.DeviceFindQuery) *result.Result {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceFindPath + "?" + params.Encode()).HttpRequest()

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

// 下发定位上报间隔指令
func (p *Device) SendLocateUpload(query *query.DeviceLocateUploadQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["second"] = query.Second
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceLocateUploadPath).SetData(data).HttpRequest()

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

// 下发定位时间段指令
func (p *Device) SendUdtime(query *query.DeviceUdtimeQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["start"] = query.Start
	data["end"] = query.End
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceUdtimePath).SetData(data).HttpRequest()

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

// 下发设置亲情号码指令
func (p *Device) SendFamily(query *query.DeviceFamilyQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["families"] = query.Families
	data["status"] = query.Status
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceFamilyPath).SetData(data).HttpRequest()

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

// 下发设置定位模式指令
func (p *Device) SendLocateMode(query *query.DeviceLocateModeQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["locate_mode"] = query.LocateMode
	data["locate_upload"] = query.LocateUpload
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceLocateModePath).SetData(data).HttpRequest()

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

// 下发设置终端host指令
func (p *Device) SendHost(query *query.DeviceHostQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["host"] = query.Host
	data["port"] = query.Port
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceHostPath).SetData(data).HttpRequest()

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

// 下发关机指令
func (p *Device) SendPowerOff(query *query.DevicePowerOffQuery) *result.Result {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDevicePowerOffPath + "?" + params.Encode()).HttpRequest()

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

// 下发重启指令
func (p *Device) SendRestart(query *query.DeviceRestartQuery) *result.Result {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceRestartPath + "?" + params.Encode()).HttpRequest()

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

// 下发聆听指令
func (p *Device) SendMonitor(query *query.DeviceMonitorQuery) *result.Result {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("mobile", query.Mobile)

	res, err := p.Cfg.HttpClient.SetMethod("get").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceMonitorPath + "?" + params.Encode()).HttpRequest()

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

// 下发设置免打扰时间段指令
func (p *Device) SendDnd(query *query.DeviceDndQuery) *result.Result {
	params := url.Values{}
	params.Set("imei_sn", query.ImeiSn)
	params.Set("dnd", query.Dnd)

	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceDndPath + "?" + params.Encode()).HttpRequest()

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

// 变更设备状态
func (p *Device) DeviceUpdateStatus(query *query.DeviceStatusQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["status"] = query.Status
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceStatusPath).SetData(data).HttpRequest()

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

// 删除设备
func (p *Device) DeviceDelete(query *query.DeviceDeleteQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	res, err := p.Cfg.HttpClient.SetMethod("delete").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceDeletePath).SetData(data).HttpRequest()

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

// 下发睡眠时间段指令
func (p *Device) DeviceSleepTime(query *query.SendSleepTimeQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["start"] = query.Start
	data["end"] = query.End
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceSleepTimePath).SetData(data).HttpRequest()

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

// DeviceWechat 下发传输微聊音频文件到设备的指令
func (p *Device) DeviceWechat(query *query.DeviceWechatQuery) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["wchat_audio_url"] = query.WechatAudioUrl
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceDeviceWechatPath).SetData(data).HttpRequest()

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

// DeviceWhitelistStatus 变更通话白名单状态
func (p *Device) DeviceWhitelistStatus(query *query.DeviceWhitelistStatus) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["whitelist_status"] = query.WhitelistStatus
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceDeviceWhitelistStatusPath).SetData(data).HttpRequest()

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

// DeviceBatchLocateMode 批量设置设备的定位模式指令
func (p *Device) DeviceBatchLocateMode(query *query.DeviceBatchLocateMode) *result.Result {
	var data = make(map[string]interface{})
	data["locatemode"] = query.LocateMode
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceBatchLocateMode).SetData(data).HttpRequest()

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

// DeviceRemind 设置设备闹钟
func (p *Device) DeviceRemind(query *query.DeviceRemind) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["remind"] = query.Remind
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceRemind).SetData(data).HttpRequest()

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

// DeviceRemind 设置设备定时开关机
func (p *Device) SetCronshutdown(query *query.TspSetCronshutdownRequest) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["is_handshutdown"] = query.IsHandshutdown
	data["is_cronshutdown"] = query.IsCronshutdown
	data["poweron_time"] = query.PoweronTime
	data["shutdown_time"] = query.ShutdownTime
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TSPDeviceCronshutdown).SetData(data).HttpRequest()

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

// DeviceRemind 发送文字消息
func (p *Device) DeviceWordsmessage(query *query.DeviceWordsmessage) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["message"] = query.Message
	res, err := p.Cfg.HttpClient.SetMethod("post").SetUrl(p.Cfg.HttpClient.GateWay + util.TspDeviceWordsmessage).SetData(data).HttpRequest()

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

// DeviceRemind 下发单词
func (p *Device) DeviceWordbook(query *query.TspSetWordBookRequest) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["id"] = query.Id
	data["title"] = query.Title
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TspDeviceWordbook).SetData(data).HttpRequest()

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

// DeviceRemind 下发古诗词
func (p *Device) DevicePoetrybook(query *query.TspSetPoetryBookRequest) *result.Result {
	var data = make(map[string]interface{})
	data["imei_sn"] = query.ImeiSn
	data["id"] = query.Id
	data["title"] = query.Title
	res, err := p.Cfg.HttpClient.SetMethod("put").SetUrl(p.Cfg.HttpClient.GateWay + util.TspDevicePoetrybook).SetData(data).HttpRequest()

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
