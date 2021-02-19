package schoolcardService

import (
	"crypto/tls"
	"funnel/app/apis/schoolcard"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"github.com/PuerkitoBio/goquery"
	"net/http"
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

	loginData := genLoginData(doc, username, password)

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
