package model

import "net/http"

type ZFUser struct {
	Username string
	Password string
	Name     string
	Session  http.Cookie
}

type LibraryUser struct {
	Username string
	Password string
	Name     string
	Session  http.Cookie
}

type CardUser struct {
	Username string
	Password string
	Name     string
	Session  http.Cookie
}
