package funnel

import "net/http"

type Student struct {
	Sid string
	Password string
	Name string
	Session http.Cookie
}

