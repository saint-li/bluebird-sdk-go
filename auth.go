package tspsdk

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gsha1"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/xu5g/bluebird-sdk-go/result"
	"github.com/xu5g/bluebird-sdk-go/util"
	"net/url"
)

type Auth struct {
	Cfg *Config
}

// 获取token
func (p *Auth) GetToken() (*result.AuthResult, error) {
	appKey := p.Cfg.HttpClient.AppKey
	tranSid := fmt.Sprintf("%s%s%s", appKey, gtime.Now().Format("YmdHis"), grand.Digits(12))
	sign := gsha1.Encrypt(fmt.Sprintf("%s%s%s%s", p.Cfg.HttpClient.Secret, appKey, tranSid, p.Cfg.HttpClient.Secret))

	params := url.Values{}
	params.Set("transid", tranSid)
	params.Set("appkey", appKey)
	params.Set("sign", sign)

	res, err := p.Cfg.HttpClient.SetUrl(p.Cfg.HttpClient.GateWay + util.TSPAuthPath + "?" + params.Encode()).SetMethod("get").HttpRequest()
	if err != nil {
		return nil, err
	}

	jsonString, _ := res.ToJsonString()
	var result = new(result.AuthResult)
	err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		return nil, err
	}

	return result, err
}
