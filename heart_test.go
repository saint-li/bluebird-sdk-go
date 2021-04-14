package tspsdk

import (
	"fmt"
	"github.com/xu5g/bluebird-sdk-go/query"
	"testing"
)

// 最新心率数据
func TestHeart_GetHeart(t *testing.T) {
	query := &query.HeartGetQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}
	res, err := NewClient(gateWay, appKey, token).Heart().GetHeart(query)
	fmt.Println(res, err)
}

// 心率数据列表
func TestHeart_GetHearts(t *testing.T) {
	query := &query.HeartsGetQuery{
		//StartTime: "2021-02-01 00:00:00",
		//EndTime:   "2021-03-31 00:00:00",
		Page:      1,
		Limit:     10,
	}

	res, err := NewClient(gateWay, appKey, token).Heart().GetHearts(query)
	fmt.Println(res, err)
}

// 获取心率上报时间间隔
func TestHeart_GetHeartUpload(t *testing.T) {
	query := &query.HeartUploadGetQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}

	res, err := NewClient(gateWay, appKey, token).Heart().GetHeartUpload(query)
	fmt.Println(res, err)
}

func TestHeart_UpdateHeartUpload(t *testing.T) {
	param := &query.HeartUploadSetQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
		Second: 400,
	}

	res, err := NewClient(gateWay, appKey, token).Heart().UpdateHeartUpload(param)
	fmt.Println(res, err)
}
