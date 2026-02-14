package register

import (
	"fmt"

	"github.com/zjutjh/mygo/config"
	"github.com/zjutjh/mygo/feishu"
	"github.com/zjutjh/mygo/foundation/kernel"
	"github.com/zjutjh/mygo/kit"
	"github.com/zjutjh/mygo/ndb"
	"github.com/zjutjh/mygo/nesty"
	"github.com/zjutjh/mygo/nlog"

	"app/comm"
	"app/register/generate"
)

func Boot() kernel.BootList {
	return kernel.BootList{
		// 基础引导器
		feishu.Boot(),   // 飞书Bot (消息提醒)
		nlog.Boot(),     // 业务日志
		generate.Boot(), // 导入生成代码

		// Client引导器
		ndb.Boot(), // DB
		// nedis.Boot(), // Redis
		nesty.Boot(), // HTTP Client

		// 业务引导器
		BizConfBoot(),
		AppBoot(),
	}
}

// BizConfBoot 初始化应用业务配置引导器
func BizConfBoot() func() error {
	return func() error {
		comm.BizConf = &comm.BizConfig{}
		err := config.Pick().UnmarshalKey("biz", comm.BizConf)
		if err != nil {
			return fmt.Errorf("%w: 解析应用业务配置错误: %w", kit.ErrDataUnmarshal, err)
		}
		return nil
	}
}

// AppBoot 应用定制引导器
func AppBoot() func() error {
	return func() error {
		// 可以在这里编写业务初始引导逻辑
		return nil
	}
}
