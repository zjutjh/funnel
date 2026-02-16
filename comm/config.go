package comm

// BizConf 业务配置
var BizConf *BizConfig

type ZFConfig struct {
	BaseURL string `mapstructure:"base_url"`
}
type BizConfig struct {
	ZF ZFConfig `mapstructure:"zf"`
}
