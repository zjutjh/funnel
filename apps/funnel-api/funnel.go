package main

import (
	"flag"
	"fmt"

	funnelapi "funnel/apps/funnel-api/app"

	"github.com/zeromicro/go-zero/core/conf"
)

type rootConfig struct {
	FunnelAPI funnelapi.Config
}

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	flag.Parse()

	var c rootConfig
	conf.MustLoad(*configFile, &c)
	server := funnelapi.NewServer(c.FunnelAPI)
	defer server.Stop()

	fmt.Printf("Starting funnel API server at %s:%d...\n", c.FunnelAPI.Host, c.FunnelAPI.Port)
	server.Start()
}
