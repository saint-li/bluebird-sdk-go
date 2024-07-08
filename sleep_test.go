package tspsdk

import (
	"fmt"
	"github.com/saint-li/bluebird-sdk-go/query"
	"testing"
)

// 睡眠数据列表
func TestSleep_GetSleeps(t *testing.T) {
	query := &query.SleepsGetQuery{
		StartTime: "2021-02-01 00:00:00",
		EndTime:   "2021-03-31 00:00:00",
		Page:      1,
		Limit:     10,
	}

	res := NewClient(gateWay, appKey, token).Sleep().GetSleeps(query)
	fmt.Println(res)
}

func TestSleep_GetSleep(t *testing.T) {
	query := &query.SleepGetQuery{
		ImeiSn: "xxxxxxxxxxxxxxx",
	}

	res := NewClient(gateWay, appKey, token).Sleep().GetSleep(query)
	fmt.Println(res)
}
