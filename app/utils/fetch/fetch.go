package fetch

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Fetch struct {
	Cookie []*http.Cookie
	client *http.Client
}

func (f Fetch) init() {
	f.client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}
}

func (f Fetch) Get(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for _, v := range f.Cookie {
		request.AddCookie(v)
	}
	response, err := f.client.Do(request)
	if err != nil {
		return nil, err
	}
	s, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	f.Cookie = append(response.Cookies(), f.Cookie...)
	return s, nil
}

func (f Fetch) PostForm(url string, requestData url.Values) ([]byte, error) {
	request, _ := http.NewRequest("POST", url, strings.NewReader(requestData.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, v := range f.Cookie {
		request.AddCookie(v)
	}
	response, err := f.client.Do(request)

	if err != nil {
		return nil, err
	}
	f.Cookie = append(response.Cookies(), f.Cookie...)
	return ioutil.ReadAll(response.Body)
}
