package loginTokenManager

import (
	"fmt"
	"funnel/app/service/zfService/zfCaptchaCracker"
	"funnel/app/service/zfService/zfLoginTokenRegister"
	"image"
	"time"
)

func RunRegister(host string) (LoginToken, error) {
	// 初始化注册器
	register := zfLoginTokenRegister.ZFLoginTokenRegister{}
	register.Init(host, "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	// 运行注册器并等待验证码图片
	img, err := register.RunAndWaitCaptchaBg()
	if err != nil {
		return LoginToken{}, err
	}

	// 尝试使用多种破解器破解验证码
	solution, err := tryCrackers(img)
	if err != nil {
		return LoginToken{}, err
	}

	// 提交验证码结果
	passed, err := register.SolveAndSubmit(solution)
	if err != nil {
		return LoginToken{}, err
	}

	if passed { // 验证码通过
		return LoginToken{
			RequestTime:    time.Now(),
			JSESSIONID:     register.CookieJSESSIONID,
			Route:          register.CookieRoute,
			CSRFToken:      register.CSRFToken,
			CryptoModulus:  register.CryptoModulus,
			CryptoExponent: register.CryptoExponent,
		}, nil
	}

	// 验证码未通过
	return LoginToken{}, fmt.Errorf("captcha not passed")
}

func tryCrackers(img image.Image) (solution string, err error) {
	solution, err = runCrackerC(img) // CompareCracker 效率较高 优先尝试
	if err == nil {
		return
	}
	solution, err = runCrackerA(img)
	return
}

func runCrackerC(img image.Image) (solution string, err error) {
	cracker := zfCaptchaCracker.ZfCaptchaCracker(&zfCaptchaCracker.CompareCracker{})
	cracker.Init("captcha_templates")
	solution, err = cracker.Crack(img)
	return
}

func runCrackerA(img image.Image) (solution string, err error) {
	cracker := zfCaptchaCracker.ZfCaptchaCracker(&zfCaptchaCracker.AnalyzeCracker{})
	cracker.Init("")
	solution, err = cracker.Crack(img)
	return
}
