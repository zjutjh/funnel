package errors

type ResponseError struct {
	Code    int
	Message string
}

func (r ResponseError) Error() string {
	panic("implement me")
}

var OK = ResponseError{200, "OK"}
var UnKnown = ResponseError{500, "未知错误"}
var RequestFailed = ResponseError{500, "请求错误"}
var InvalidArgs = ResponseError{400, "参数错误"}

var NotLogin = ResponseError{403, "未登录"}
var WrongPassword = ResponseError{403, "密码错误"}
