package tspsdk

import (
	"fmt"
	"github.com/saint-li/bluebird-sdk-go/query"
	"testing"
)

// 计步数据列表
func TestStep_GetSteps(t *testing.T) {
	query := &query.StepsGetQuery{
		EndTime:   "2021-04-30 00:00:00",
		StartTime: "2021-02-01 00:00:00",
		Page:      1,
		Limit:     10,
	}

	res := NewClient(gateWay, appKey, token).Step().GetSteps(query)
	fmt.Println(res)
}
