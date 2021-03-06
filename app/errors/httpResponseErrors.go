package errors

type HttpResponseError struct {
	Code    int
	Message string
}

var OK = HttpResponseError{200, "OK"}
var UnKnown = HttpResponseError{500, "未知错误"}
var RequestFailed = HttpResponseError{510, "请求错误"}

var InvalidArgs = HttpResponseError{410, "参数错误"}
var WrongPassword = HttpResponseError{412, "密码错误"}
var CaptchaFailed = HttpResponseError{413, "验证码错误"}
var SessionExpired = HttpResponseError{414, "缓存过期"}
