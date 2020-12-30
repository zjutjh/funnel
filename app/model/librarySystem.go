package model

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type LibrarySystem struct {
	RootUrl       string
	ClassTableUrl string
}

func (t *LibrarySystem) Login(user *LibraryUser) error {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	url0 := "http://210.32.205.60/login.aspx"
	loginData := url.Values{"TextBox1": {user.Username}, "TextBox2": {user.Password}}
	response, _ := client.PostForm(url0,loginData)
	session := response.Cookies()[0]
	user.Session=*session
	return nil
}


func (t *LibrarySystem) GetBorrowHistory(user *LibraryUser) string {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}
	requestUrl := "http://210.32.205.60/login.aspx"
	request, _ := http.NewRequest("GET",requestUrl,nil)
	request.AddCookie(&user.Session)
	response, _ := client.Do(request)
	s, _ := ioutil.ReadAll(response.Body)
	return string(s)
}