package schoolcardService

import (
	"funnel/app/apis/schoolcard"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"funnel/app/utils/fetch"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func GetUser(username string, password string) (*model.User, error) {
	user, err := service.GetUser(service.CardPrefix, username, password)
	if err != nil {
		return login(username, password)
	}
	return user, err
}

func login(username string, password string) (*model.User, error) {
	f := fetch.Fetch{}
	f.InitUnSafe()
	res, err := f.GetRaw(schoolcard.CardHome)
	if err != nil {
		return nil, err
	}
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	loginData := genLoginData(doc, username, password)
	response, err := f.PostFormRaw(schoolcard.CardHome, loginData)

	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusFound {
		return nil, errors.ERR_WRONG_PASSWORD
	}

	cookie := res.Cookies()[0]
	return service.SetUser(service.CardPrefix, username, password, cookie)
}
