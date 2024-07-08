package tspsdk

import "github.com/saint-li/bluebird-sdk-go/util"

type Config struct {
	HttpClient *util.Request
}

type Tsp struct {
	Cfg *Config
}

// 获取其它接口数据时在header里面携带token
func NewClient(gateWay, appKey, token string) *Tsp {
	return &Tsp{
		Cfg: &Config{
			HttpClient: &util.Request{
				Token:   token,
				GateWay: gateWay,
				AppKey:  appKey,
			},
		},
	}
}

// 获取token（只用来获取token）
func NewAuth(gateWay, appKey, secret string) *Tsp {
	return &Tsp{
		Cfg: &Config{
			&util.Request{
				AppKey:  appKey,
				Secret:  secret,
				GateWay: gateWay,
			},
		},
	}
}

// 鉴权
func (p *Tsp) Auth() *Auth {
	return &Auth{
		Cfg: p.Cfg,
	}
}

// 体温
func (p *Tsp) Temperature() *Temperature {
	return &Temperature{
		Cfg: p.Cfg,
	}
}

// 心率
func (p *Tsp) Heart() *Heart {
	return &Heart{
		Cfg: p.Cfg,
	}
}

// 血压
func (p *Tsp) Blood() *Blood {
	return &Blood{
		Cfg: p.Cfg,
	}
}

// 血氧
func (p *Tsp) BloodOxygen() *BloodOxygen {
	return &BloodOxygen{
		Cfg: p.Cfg,
	}
}

// 计步
func (p *Tsp) Step() *Step {
	return &Step{
		Cfg: p.Cfg,
	}
}

// 睡眠
func (p *Tsp) Sleep() *Sleep {
	return &Sleep{
		Cfg: p.Cfg,
	}
}

// 轨迹
func (p *Tsp) Track() *Track {
	return &Track{
		Cfg: p.Cfg,
	}
}

// 设备
func (p *Tsp) Device() *Device {
	return &Device{
		Cfg: p.Cfg,
	}
}

// 围栏
func (p *Tsp) Fence() *Fence {
	return &Fence{
		Cfg: p.Cfg,
	}
}

// maclist
func (p *Tsp) MacList() *MacList {
	return &MacList{
		Cfg: p.Cfg,
	}
}

// macbook
func (p *Tsp) MacBook() *MacBook {
	return &MacBook{
		Cfg: p.Cfg,
	}
}

// 考勤
func (p *Tsp) Attence() *Attence {
	return &Attence{
		Cfg: p.Cfg,
	}
}

// 报文
func (p *Tsp) Message() *Message {
	return &Message{
		Cfg: p.Cfg,
	}
}

// api日志
func (p *Tsp) ApiLog() *ApiLog {
	return &ApiLog{
		Cfg: p.Cfg,
	}
}

// core日志
func (p *Tsp) CoreLog() *CoreLog {
	return &CoreLog{
		Cfg: p.Cfg,
	}
}

// guard日志
func (p *Tsp) GuardLog() *GuardLog {
	return &GuardLog{
		Cfg: p.Cfg,
	}
}

// 萤石token
func (p *Tsp) YsMonitor() *YsMonitor {
	return &YsMonitor{
		Cfg: p.Cfg,
	}
}

// 跳绳
func (p *Tsp) Skip() *Skip {
	return &Skip{
		Cfg: p.Cfg,
	}
}

// 跳绳蓝牙
func (p *Tsp) Bluetooth() *Bluetooth {
	return &Bluetooth{
		Cfg: p.Cfg,
	}
}
