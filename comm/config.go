package comm

// BizConf 业务配置
var BizConf *BizConfig

type ZFConfig struct {
	BaseURL string   `mapstructure:"base_url"`
	Public  struct { // 用于公共查询的相关配置
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"public"`
	SessionPool struct {
		TTLMinute   int `mapstructure:"ttl_minute"`
		MaxSize     int `mapstructure:"max_size"`
		FillWorkers int `mapstructure:"fill_workers"`
	} `mapstructure:"session_pool"`
}
type BizConfig struct {
	ZF ZFConfig `mapstructure:"zf"`
}
