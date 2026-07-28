package config

type ZFConfig struct {
	BaseURL     string              `json:"base_url,optional"`
	Public      ZFPublicConfig      `json:"public,optional"`
	SessionPool ZFSessionPoolConfig `json:"session_pool,optional"`
}

type ZFPublicConfig struct {
	Username string `json:"username,optional"`
	Password string `json:"password,optional"`
}

type ZFSessionPoolConfig struct {
	TTLMinute   int `json:"ttl_minute,optional"`
	MaxSize     int `json:"max_size,optional"`
	FillWorkers int `json:"fill_workers,optional"`
}

type OauthConfig struct {
	BaseURL string `json:"base_url,optional"`
}
