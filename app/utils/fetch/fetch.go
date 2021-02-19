package fetch

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Fetch struct {
	Cookie []*http.Cookie
	client *http.Client
}

func (f *Fetch) Init() {
	f.client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
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
	s, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return s, nil
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

func (f *Fetch) PostForm(url string, requestData url.Values) ([]byte, error) {
	request, _ := http.NewRequest("POST", url, strings.NewReader(requestData.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, v := range f.Cookie {
		request.AddCookie(v)
	}
	response, err := f.client.Do(request)

	if err != nil {
		return nil, err
	}
	f.Cookie = cookieMerge(f.Cookie, response.Cookies())
	return ioutil.ReadAll(response.Body)
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