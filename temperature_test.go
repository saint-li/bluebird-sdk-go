package tspsdk

import (
	"fmt"
	"github.com/saint-li/bluebird-sdk-go/query"
	"testing"
)

// 最新体温数据
func TestTemperature_GetTemperature(t *testing.T) {
	query := &query.TemperatureGetQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}
	res := NewClient(gateWay, appKey, token).Temperature().GetTemperature(query)
	fmt.Println(res)
}

// 体温数据列表
func TestTemperature_GetTemperatures(t *testing.T) {
	query := &query.TemperaturesGetQuery{
		Page:  1,
		Limit: 10,
	}

	res := NewClient(gateWay, appKey, token).Temperature().GetTemperatures(query)
	fmt.Println(res)
}

// 获取体温上报时间间隔
func TestTemperature_GetTemperatureUpload(t *testing.T) {
	query := &query.TemperatureUploadGetQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}

	res := NewClient(gateWay, appKey, token).Temperature().GetTemperatureUpload(query)
	fmt.Println(res)
}

func TestTemperature_UpdateTemperatureUpload(t *testing.T) {
	param := &query.TemperatureUploadSetQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
		Second: 400,
	}

	res := NewClient(gateWay, appKey, token).Temperature().UpdateTemperatureUpload(param)
	fmt.Println(res)
}
