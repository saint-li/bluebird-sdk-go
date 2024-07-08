package tspsdk

import (
	"fmt"
	"github.com/saint-li/bluebird-sdk-go/query"
	"testing"
)

// 获取围栏列表
func TestFence_GetFences(t *testing.T) {
	query := &query.FencesGetQuery{
		Page:  1,
		Limit: 10,
	}

	res := NewClient(gateWay, appKey, token).Fence().GetFences(query)
	fmt.Println(res)
}

// 创建围栏
func TestFence_CreateFence(t *testing.T) {
	query := &query.FenceCreateQuery{
		Truename:    "围栏1",
		FenceType:   1,
		ShapeType:   2,
		CollideType: 2,
		Points:      "xxxx",
		NearRadius:  1,
		ValidStart:  "xxx",
		ValidEnd:    "xxx",
		ValidWeek:   "xxx",
	}

	res := NewClient(gateWay, appKey, token).Fence().CreateFence(query)
	fmt.Println(res)
}

// 删除围栏
func TestFence_DeleteFence(t *testing.T) {
	query := &query.FenceDeleteQuery{
		Id: 1,
	}

	res := NewClient(gateWay, appKey, token).Fence().DeleteFence(query)
	fmt.Println(res)
}
