package oauth

import (
	"context"
	"funnel/worker/internal/config"

	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type OauthClient struct {
	logx.Logger
	BaseURL    string
	HTTPClient *resty.Client
}

func New(ctx context.Context, cfg config.OauthConfig) *OauthClient {
	if ctx == nil {
		ctx = context.Background()
	}

	httpClient := resty.New()
	httpClient.SetBaseURL(cfg.BaseURL)

	return &OauthClient{
		Logger:     logx.WithContext(ctx),
		BaseURL:    cfg.BaseURL,
		HTTPClient: httpClient,
	}
}
