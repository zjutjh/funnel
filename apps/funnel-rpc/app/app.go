package funnelrpc

import (
	"funnel/apps/funnel-rpc/internal/config"
	"funnel/apps/funnel-rpc/internal/server"
	"funnel/apps/funnel-rpc/internal/svc"
	"funnel/apps/funnel-rpc/pb"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config = config.Config

func NewServer(c Config) *zrpc.RpcServer {
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterWorkerZfServiceServer(grpcServer, server.NewWorkerZfServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	return s
}
