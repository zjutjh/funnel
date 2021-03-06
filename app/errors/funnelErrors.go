package errors

import "errors"

var ERR_WRONG_PASSWORD = errors.New("wrong password")
var ERR_TIMEOUT = errors.New("timeout")
var ERR_UNKNOWN_LOGIN_ERROR = errors.New("unknown login error")
var ERR_Session_Expired = errors.New("Session expired")
var ERR_JSON_DESER = errors.New("ERR_JSON_DESER")
var ERR_WRONG_Captcha = errors.New("ERR_WRONG_Captcha")
var ERR_INVALID_ARGS = errors.New("invalid args")
var ERR_SESSION_NOT_EXIST = errors.New("ERR_SESSION_NOT_EXIST")
var ERR_SESSION_EXPIRES = errors.New("ERR_SESSION_EXPIRES")
