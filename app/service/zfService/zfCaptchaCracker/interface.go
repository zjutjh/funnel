package zfCaptchaCracker

import "image"

type ZfCaptchaCracker interface {
	// 通过参数(如配置)初始化破解器
	Init(string) error
	// 破解验证码并返回结果(轨迹包)
	Crack(image image.Image) (string, error)
}
