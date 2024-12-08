package request

import (
	"funnel/app/service/libraryService/request/midware"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	*resty.Client
}

func New() Client {
	s := Client{
		Client: resty.New(),
	}
	// 利用中间件实现请求日志
	s.OnAfterResponse(midware.LogMiddleware)
	return s
}

func (s Client) Request() *resty.Request {
	return s.R().
		EnableTrace().
		SetHeader("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36")
}
