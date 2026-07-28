package main

import (
	"flag"
	"fmt"

	funnelrpc "funnel/apps/funnel-rpc/app"

	"github.com/zeromicro/go-zero/core/conf"
)

type rootConfig struct {
	FunnelRPC funnelrpc.Config
}

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	flag.Parse()

	var c rootConfig
	conf.MustLoad(*configFile, &c)
	s := funnelrpc.NewServer(c.FunnelRPC)
	defer s.Stop()

	fmt.Printf("Starting funnel RPC server at %s...\n", c.FunnelRPC.ListenOn)
	s.Start()
}
