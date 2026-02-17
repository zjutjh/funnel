package comm

// BizConf 业务配置
var BizConf *BizConfig

type ZFConfig struct {
	BaseURL string `mapstructure:"base_url"`
	Public  struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"public"`
}
type BizConfig struct {
	ZF ZFConfig `mapstructure:"zf"`
}
