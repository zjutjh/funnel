package comm

import (
	"github.com/zjutjh/mygo/kit"
)

var (
	WrongUsernameOrPasswordError = NewError(CodeWrongUsernameOrPassword)
	OauthClosedError             = NewError(CodeOauthClosed)
	OauthPasswordNeedEditError   = NewError(CodeOauthPasswordNeedEdit)
	OauthNotActivatedError       = NewError(CodeOauthNotActivated)
	UnknownError                 = NewError(CodeUnknownError)
)

type Error struct {
	Msg string
}

func (e Error) Error() string {
	return e.Msg
}

func NewError(code kit.Code) *Error {
	return &Error{
		Msg: code.Message,
	}
}
