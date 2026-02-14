package main

import (
	"sync"

	"github.com/spf13/cobra"
	"github.com/zjutjh/mygo/foundation/command"
	"github.com/zjutjh/mygo/foundation/crontab"
	"github.com/zjutjh/mygo/foundation/httpserver"

	"app/register"
)

func main() {
	command.Execute(
		register.Boot,    // 应用引导注册器
		register.Command, // 应用命令注册器
		// httpserver.CommandRegister(register.Route), // 默认注入HTTP Server注册器
		func(cmd *cobra.Command, args []string) error {
			wg := &sync.WaitGroup{}

			// 启动HTTP Server
			wg.Add(1)
			go func() {
				defer wg.Done()
				httpserver.StartHTTPServer(register.Route)
			}()

			// 启动HTTP Server伴生定时任务
			wg.Add(1)
			go func() {
				defer wg.Done()
				crontab.Run(register.CronWithHTTPServer)
			}()

			// 如有需要 可以额外启动其他服务

			wg.Wait()
			return nil
		},
	)
}
