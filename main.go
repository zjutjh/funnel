package main

import (
	"flag"
	"fmt"
	"sync"

	funnelapi "funnel/apps/funnel-api/app"
	funnelrpc "funnel/apps/funnel-rpc/app"

	"github.com/zeromicro/go-zero/core/conf"
)

type rootConfig struct {
	FunnelAPI funnelapi.Config
	FunnelRPC funnelrpc.Config
}

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	flag.Parse()

	var c rootConfig
	conf.MustLoad(*configFile, &c)

	rpcServer := funnelrpc.NewServer(c.FunnelRPC)
	defer rpcServer.Stop()

	fmt.Printf("Starting funnel RPC server at %s...\n", c.FunnelRPC.ListenOn)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		rpcServer.Start()
	}()

	httpServer := funnelapi.NewServer(c.FunnelAPI)
	defer httpServer.Stop()

	fmt.Printf("Starting funnel API server at %s:%d...\n", c.FunnelAPI.Host, c.FunnelAPI.Port)

	go func() {
		defer wg.Done()
		httpServer.Start()
	}()

	wg.Wait()
}
