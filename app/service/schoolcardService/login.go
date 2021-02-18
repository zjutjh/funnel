package schoolcardService

import (
	"crypto/tls"
	"funnel/app/apis/schoolcard"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
)

func login(username string, password string) (*model.User, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Transport:     transport,
	}

	res, _ := client.Get(schoolcard.CardHome)
	SSIONID := res.Cookies()[0]
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	code := string(doc.Find("#UserLogin_ImgFirst").AttrOr("src", "0")[7])
	code += string(doc.Find("#UserLogin_imgSecond").AttrOr("src", "0")[7])
	code += string(doc.Find("#UserLogin_imgThird").AttrOr("src", "0")[7])
	code += string(doc.Find("#UserLogin_imgFour").AttrOr("src", "0")[7])
	loginData := url.Values{
		"__VIEWSTATE":              {doc.Find("#__VIEWSTATE").AttrOr("value", "")},
		"__VIEWSTATEGENERATOR":     {doc.Find("#__VIEWSTATEGENERATOR").AttrOr("value", "")},
		"__EVENTVALIDATION":        {doc.Find("#__EVENTVALIDATION").AttrOr("value", "")},
		"UserLogin:txtUser":        {username},
		"UserLogin:txtPwd":         {password},
		"UserLogin:ddlPerson":      {"\xBF\xA8\xBB\xA7"},
		"UserLogin:txtSure":        {code},
		"UserLogin:ImageButton1.x": {"48"},
		"UserLogin:ImageButton1.y": {"8"}}
	request, _ := http.NewRequest("POST", schoolcard.CardHome, strings.NewReader(loginData.Encode()))
	request.AddCookie(SSIONID)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, _ := client.Do(request)

	if response.StatusCode != http.StatusFound {
		return nil, errors.ERR_WRONG_PASSWORD
	}

	cookie := SSIONID
	return service.SetUser(service.CardPrefix, username, password, cookie)
}
