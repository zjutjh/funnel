package oauth

import (
	"funnel/comm"

	"github.com/go-resty/resty/v2"
	"github.com/zjutjh/mygo/nesty"
)

type OauthClient struct {
	BaseURL    string
	HTTPClient *resty.Client
}

func New() *OauthClient {
	return &OauthClient{
		BaseURL:    comm.BizConf.ZF.BaseURL,
		HTTPClient: nesty.Pick(),
	}
}
