package midware

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

// LogMiddleware 日志中间件
func LogMiddleware(client *resty.Client, resp *resty.Response) error {
	fmt.Println("==================================================================")
	method := resp.Request.Method
	url := resp.Request.URL
	fmt.Printf("%s %s \n", method, url)
	//fmt.Println(resp.Request.Header)
	//fmt.Println(resp)
	fmt.Printf("Time Spent: %v\n", resp.Time())

	return nil
}
