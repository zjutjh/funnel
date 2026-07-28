package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	ZF    ZFConfig    `json:"zf,optional"`
	Oauth OauthConfig `json:"oauth,optional"`
}
