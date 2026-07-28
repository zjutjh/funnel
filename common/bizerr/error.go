package bizerr

import (
	"context"
	"errors"
	"net/http"
)

const CodeOK = 0

// 系统错误码
const (
	CodeUnknownError           = 10000
	CodeThirdServiceError      = 10001
	CodeDatabaseError          = 10002
	CodeRedisError             = 10003
	CodeMiddlewareServiceError = 10004
)

// 业务通用错误码
const (
	CodeNotLoggedIn        = 20000
	CodeLoginExpired       = 20001
	CodePermissionDenied   = 20002
	CodeParameterInvalid   = 20003
	CodeDataParseError     = 20004
	CodeDataNotFound       = 20005
	CodeDataConflict       = 20006
	CodeServiceMaintenance = 20007
	CodeTooFrequently      = 20008
)

// 业务错误码从 30000 开始。
const (
	CodeWrongUsernameOrPassword = 30000
	CodeOauthClosed             = 30001
	CodeOauthPasswordNeedEdit   = 30002
	CodeOauthNotActivated       = 30003
)

var (
	ErrWrongUsernameOrPassword = New(CodeWrongUsernameOrPassword, "用户名或密码错误")
	ErrOauthClosed             = New(CodeOauthClosed, "统一系统在夜间关闭")
	ErrOauthPasswordNeedEdit   = New(CodeOauthPasswordNeedEdit, "统一密码需要修改, 请手动登录统一修改")
	ErrOauthNotActivated       = New(CodeOauthNotActivated, "统一账号未激活")
	ErrUnknown                 = New(CodeUnknownError, "未知错误")
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func New(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	return e.Msg
}

func (e *Error) Response() Response {
	if e == nil {
		return Response{}
	}

	return Response{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

func From(err error) (*Error, bool) {
	var bizErr *Error
	if errors.As(err, &bizErr) {
		return bizErr, true
	}

	return nil, false
}

func HTTPErrorHandler(_ context.Context, err error) (int, any) {
	if bizErr, ok := From(err); ok {
		return http.StatusOK, bizErr.Response()
	}

	return http.StatusOK, err
}
