package funnel

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type TeachingAdministrationSystem struct {
	RootUrl string
}

type RSAPublicKey struct {
	Modulus  string `json:"modulus"`
	Exponent string `json:"exponent"`
}

type yzmServerResponse struct {
	Code string `json:"code"`
	Data string `json:"data"`
}

func (t *TeachingAdministrationSystem) GetClassTable(stu *Student, year string, term string) string {
	return FetchTermRelatedInfo(t.RootUrl+"kbcx/xskbcx_cxXsKb.html", year, term, stu)
}
func (t *TeachingAdministrationSystem) GetExamInfo(stu *Student, year string, term string) string {
	return FetchTermRelatedInfo(t.RootUrl+"/kwgl/kscx_cxXsksxxIndex.html", year, term, stu)
}
func (t *TeachingAdministrationSystem) GetScoreDetail(stu *Student, year string, term string) string {
	return FetchTermRelatedInfo(t.RootUrl+"cjcx/cjcx_cxXsKccjList.html", year, term, stu)
}
func (t *TeachingAdministrationSystem) GetScore(stu *Student, year string, term string) string {
	return FetchTermRelatedInfo(t.RootUrl+"cjcx/cjcx_cxDgXscj.html?doType=query", year, term, stu)
}

func FetchTermRelatedInfo(url0 string, year string, term string, stu *Student) string {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	data4 := url.Values{"xnm": {year}, "xqm": {term}, "queryModel.showCount": {"1500"}}
	request, _ := http.NewRequest("POST", url0, strings.NewReader(data4.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.AddCookie(&stu.Session)

	response, _ := client.Do(request)
	s, _ := ioutil.ReadAll(response.Body)

	return string(s)
}

func (t *TeachingAdministrationSystem) Login(stu *Student) error {

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	url0 := t.RootUrl + "xtgl/login_slogin.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
	response, _ := client.Get(url0)
	JSESSIONID := response.Cookies()[0]

	url1 := t.RootUrl + "xtgl/login_getPublicKey.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
	request, _ := http.NewRequest("GET", url1, nil)
	request.AddCookie(JSESSIONID)
	response, _ = client.Do(request)
	s, _ := ioutil.ReadAll(response.Body)

	p := &RSAPublicKey{}
	json.Unmarshal(s, p)

	nString, _ := base64.StdEncoding.DecodeString(p.Modulus)
	n, _ := new(big.Int).SetString(hex.EncodeToString(nString), 16)
	eString, _ := base64.StdEncoding.DecodeString(p.Exponent)
	e, _ := strconv.ParseInt(hex.EncodeToString(eString), 16, 32)
	pub := rsa.PublicKey{E: int(e), N: n}
	cc, _ := rsa.EncryptPKCS1v15(rand.Reader, &pub, []byte(stu.Password))
	base64Password := base64.StdEncoding.EncodeToString(cc)

	url2 := t.RootUrl + "kaptcha?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
	request, _ = http.NewRequest("GET", url2, nil)
	request.AddCookie(JSESSIONID)

	response, _ = client.Do(request)

	s, _ = ioutil.ReadAll(response.Body)
	ymzbase := base64.StdEncoding.EncodeToString(s)

	url3 := "http://172.16.32.50/yzm"
	response, _ = client.PostForm(url3, url.Values{"img_base64": {"data:image/jpeg;base64," + ymzbase}})
	yzm, _ := ioutil.ReadAll(response.Body)

	yzmRes := &yzmServerResponse{}
	json.Unmarshal(yzm, yzmRes)

	url4 := t.RootUrl + "xtgl/login_slogin.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)

	data4 := url.Values{"yhm": {stu.Sid}, "mm": {base64Password}, "yzm": {yzmRes.Data}}
	request, _ = http.NewRequest("POST", url4, strings.NewReader(data4.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.AddCookie(JSESSIONID)

	response, _ = client.Do(request)
	s, _ = ioutil.ReadAll(response.Body)
	if len(response.Cookies()) < 2 {
		return errors.New("Login Failed")
	}
	JSESSIONID = response.Cookies()[1]
	stu.Session = *JSESSIONID
	return nil
}
