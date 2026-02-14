package register

import (
	"github.com/spf13/cobra"
	"github.com/zjutjh/mygo/foundation/command"
	"github.com/zjutjh/mygo/foundation/crontab"
	"github.com/zjutjh/mygo/foundation/httpserver"
)

func Command(root *cobra.Command) {
	// 基础命令
	command.Add("server", httpserver.CommandRegister(Route))
	command.Add("cron", crontab.CommandRegister(Cron))

	// 业务命令
	//command.Add("xxx", cmd.XXXRun)
}
