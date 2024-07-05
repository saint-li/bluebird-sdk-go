package tspsdk

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xu5g/bluebird-sdk-go/query"
	"testing"
)

// 获取macbook列表
func TestMacBook_GetMacBooks(t *testing.T) {
	params := &query.MacBooksGetQuery{
		StartTime: "2021-03-12",
		EndTime:   "2021-04-22",
		Page:      1,
		Limit:     10,
	}

	res := NewClient(gateWay, appKey, token).MacBook().GetMacBooks(params)
	fmt.Println(res)
}

// 获取macbook详情
func TestMacBook_GetMacBook(t *testing.T) {
	params := &query.MacBookGetQuery{
		Mac: "xxxxxxxxxxxxxxxxxxxxxxxxx",
	}

	res := NewClient(gateWay, appKey, token).MacBook().GetMacBook(params)
	fmt.Println(res)
	fmt.Println(gconv.Map(res.Data))
}

// 更新macbook
func TestMacBook_UpdateMacBook(t *testing.T) {
	params := &query.MacbookUpdateQuery{
		Name:       "11",
		Mac:        "xxxxxxxxxxxxxxxxxxxxxxxxx",
		ProvinceId: 410000,
		CityId:     410100,
		DistrictId: 410102,
		Lat:        0,
		Lng:        0,
		Status:     1,
		TotalCount: 1,
		Remark:     "备注信息",
	}

	res := NewClient(gateWay, appKey, token).MacBook().UpdateMacBook(params)
	fmt.Println(res)
}

// 重绘macbook
func TestMacBook_DrawMacBook(t *testing.T) {
	params := &query.MacbookDrawQuery{
		Name: "11",
		Mac:  "xxxxxxxxxxxxxxxxxxxxxxxxx",
	}

	res := NewClient(gateWay, appKey, token).MacBook().DrawMacBook(params)
	fmt.Println(res)
}

// 删除macbook
func TestMacBook_DeleteMacBook(t *testing.T) {
	params := &query.MacbookDeleteQuery{
		MakAddr: "xxxxxxxxxxxxxxxxxxxxxxxxxx",
		Mac:     "xxxxxxxxxxxxxxxxxxxxxxxxx",
	}

	res := NewClient(gateWay, appKey, token).MacBook().DeleteMacBook(params)
	fmt.Println(res)
}
