/*
用来向控制器提供客流量相关数据的服务
*/

package canteenService

import (
	"funnel/app/apis/canteen"
	"funnel/app/utils/fetch"
)

func FetchFlow() (string, bool) {
	var client fetch.Fetch
	client.Init()
	result, err := client.Get(canteen.CanteenFlow)
	if err != nil { // 获取数据失败
		return "fetchError", false
	} else { // 获取数据成功
		return string(result), true // 返回抓取到的食堂客流量数据
	}
}
