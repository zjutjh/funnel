/*
用来向控制器提供客流量相关数据的服务
*/

package canteenService

import (
	"encoding/json"
	"funnel/app/apis/canteen"
	"funnel/app/utils/fetch"
)

type response struct {
	Data struct {
		BeginTime string `json:"beginTime"`
		EndTime   string `json:"endTime"`
		Data      []struct {
			DealCount      string `json:"dealCount"`
			RestaurantName string `json:"restaurantName"`
		} `json:"data"`
	}
}

func FetchFlow() (interface{}, error) {
	var client fetch.Fetch
	client.Init()
	result, err := client.Get(canteen.FlowApi)
	if err != nil { // 获取数据失败
		return "fetchError", err
	} else { // 获取数据成功
		var f response
		err = json.Unmarshal(result, &f)
		return f.Data, err // 返回抓取到的食堂客流量数据
	}
}
