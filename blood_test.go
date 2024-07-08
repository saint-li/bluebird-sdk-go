package tspsdk

import (
	"fmt"
	"github.com/saint-li/bluebird-sdk-go/query"
	"testing"
)

// 最新血压数据
func TestBlood_GetBlood(t *testing.T) {
	query := &query.BloodGetQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}
	res := NewClient(gateWay, appKey, token).Blood().GetBlood(query)
	fmt.Println(res)
}

// 血压数据列表
func TestBlood_GetBloods(t *testing.T) {
	query := &query.BloodsGetQuery{
		Page:  1,
		Limit: 10,
	}
	res := NewClient(gateWay, appKey, token).Blood().GetBloods(query)
	fmt.Println(res)
}

// 获取血压上报间隔
func TestBlood_GetBloodUpload(t *testing.T) {
	query := &query.BloodUploadGetQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}
	res := NewClient(gateWay, appKey, token).Blood().GetBloodUpload(query)
	fmt.Println(res)
}

// 设置血压上报间隔
func TestBlood_UpdateBloodUpload(t *testing.T) {
	query := &query.BloodUploadSetQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
		Second: 500,
	}
	res := NewClient(gateWay, appKey, token).Blood().UpdateBloodUpload(query)
	fmt.Println(res)
}
