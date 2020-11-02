package errors

type ResponseError struct {
	Code    int
	Message string
}

var OK = ResponseError{200, "OK"}

var WrongPassword = ResponseError{403, "密码错误"}
var UnKnown = ResponseError{403, "未知错误"}

var NotLogin = ResponseError{403, "未登录"}
