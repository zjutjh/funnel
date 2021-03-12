/*
用来向控制器提供客流量相关数据的服务
*/

package canteenService

import (
	"funnel/app/apis/canteen"
	"funnel/app/utils/fetch"
)

func FetchFlow() (string, error) {
	var client fetch.Fetch
	client.Init()
	result, err := client.Get(canteen.FlowApi)
	if err != nil { // 获取数据失败
		return "fetchError", err
	} else { // 获取数据成功
		return string(result), err // 返回抓取到的食堂客流量数据
	}
}
