package model

import "net/http"

type User struct {
	Username string
	Password string
	Session  http.Cookie
}
