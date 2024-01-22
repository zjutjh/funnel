package fetch

import (
	"crypto/tls"
	"fmt"
	errors2 "funnel/app/errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Fetch struct {
	Cookie []*http.Cookie
	client *http.Client
}

func (f *Fetch) InitUnSafe() {
	f.client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:       time.Second * 20,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func (f *Fetch) Init() {
	f.client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:       time.Second * 20,
	}
}
func (f *Fetch) SkipTlsCheck() {
	f.client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

func (f *Fetch) Get(url string) ([]byte, error) {
	response, err := f.GetRaw(url)
	if err != nil {
		return nil, err
	}
	s, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (f *Fetch) GetRedirect(url string) (*url.URL, error) {
	response, err := f.GetRaw(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 302 {
		return nil, errors2.ERR_OAUTH_ERROR
	}
	location, err := response.Location()
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (f *Fetch) GetRaw(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for _, v := range f.Cookie {
		request.AddCookie(v)
	}
	f.client.Do(request)
	response, err := f.client.Do(request)
	if err != nil {
		return nil, err
	}
	f.Cookie = cookieMerge(f.Cookie, response.Cookies())
	return response, err
}

func (f *Fetch) PostFormRaw(url string, requestData url.Values) (*http.Response, error) {
	request, _ := http.NewRequest("POST", url, strings.NewReader(requestData.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, v := range f.Cookie {
		request.AddCookie(v)
	}
	return f.client.Do(request)
}

func (f *Fetch) PostFormRawAsynchronous(url string, requestData url.Values) (*http.Response, error) {
	request, _ := http.NewRequest("POST", url, strings.NewReader(requestData.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	for _, v := range f.Cookie {
		request.AddCookie(v)
	}
	return f.client.Do(request)
}

func (f *Fetch) PostFormRedirect(url string, requestData url.Values) (*url.URL, error) {
	response, err := f.PostFormRaw(url, requestData)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 302 {
		fmt.Println(url)
		fmt.Println(response.StatusCode)
		return nil, errors2.ERR_WRONG_PASSWORD
	}
	f.Cookie = cookieMerge(f.Cookie, response.Cookies())
	return response.Location()
}

func (f *Fetch) PostForm(url string, requestData url.Values) ([]byte, error) {
	response, err := f.PostFormRaw(url, requestData)
	if err != nil {
		return nil, err
	}
	f.Cookie = cookieMerge(f.Cookie, response.Cookies())
	return io.ReadAll(response.Body)
}

func cookieMerge(cookieA []*http.Cookie, cookieB []*http.Cookie) []*http.Cookie {

	for _, v := range cookieB {
		for k, v2 := range cookieA {
			if v.Name == v2.Name {
				cookieA = append(cookieA[:k], cookieA[k+1:]...)
				break
			}
		}
	}
	cookieA = append(cookieA, cookieB...)
	return cookieA
}
