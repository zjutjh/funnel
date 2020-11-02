package errors

import "errors"

var ERR_EOF = errors.New("EOF")
var ERR_CLOSED_PIPE = errors.New("io: read/write on closed pipe")
var ERR_NO_PROGRESS = errors.New("multiple Read calls return no data or error")
var ERR_SHORT_BUFFER = errors.New("short buffer")
var ERR_SHORT_WRITE = errors.New("short write")
var ERR_UNEXPECTED_EOF = errors.New("unexpected EOF")

var ERR_WRONG_PASSWORD = errors.New("wrong password")
var ERR_TIMEOUT = errors.New("timeout")
var ERR_UNKNOWN_LOGIN_ERROR = errors.New("unknown login error")
var ERR_REQUEST_FAIL = errors.New("timeout")
