package tspsdk

import (
	"fmt"
	"github.com/saint-li/bluebird-sdk-go/query"
	"testing"
)

// 轨迹数据列表
func TestTrack_GetTracks(t *testing.T) {
	query := &query.TracksGetQuery{
		StartTime: "2021-02-01 00:00:00",
		EndTime:   "2021-04-30 23:00:00",
		Page:      1,
		Limit:     10,
		ImeiSn:    "863659041488510",
	}

	res := NewClient(gateWay, appKey, token).Track().GetTracks(query)
	fmt.Println(res)
}

// TestTrack_TrackUpdate 更新轨迹数据
func TestTrack_TrackUpdate(t *testing.T) {
	query := &query.TracksUpdateQuery{
		Id:         0,
		PrimaryKey: "000",
		IsIgnore:   0,
	}
	res := NewClient(gateWay, appKey, token).Track().TrackUpdate(query)
	fmt.Println(res)

}
