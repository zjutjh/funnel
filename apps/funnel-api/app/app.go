package funnelapi

import (
	"funnel/apps/funnel-api/internal/config"
	"funnel/apps/funnel-api/internal/handler"
	"funnel/apps/funnel-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

type Config = config.Config

func NewServer(c Config) *rest.Server {
	server := rest.MustNewServer(c.RestConf)
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	return server
}
