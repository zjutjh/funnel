package libraryService

import (
	"funnel/app/apis/library"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"funnel/app/utils/fetch"
	"github.com/PuerkitoBio/goquery"
)

func GetUser(username string, password string) (*model.User, error) {
	user, err := service.GetUser(service.LibraryPrefix, username, password)
	if err != nil {
		return login(username, password)
	}
	return user, err
}

func login(username string, password string) (*model.User, error) {
	f := fetch.Fetch{}
	f.Init()

	res, err := f.GetRaw(library.LibraryLogin)
	if err != nil {
		return nil, err
	}
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	loginData := genLoginData(doc, username, password)
	_, err = f.PostForm(library.LibraryLogin, loginData)
	if err != nil {
		return nil, err
	}
	if len(f.Cookie) == 0 || err != nil {
		return nil, errors.ERR_WRONG_PASSWORD
	}
	return service.SetUser(service.LibraryPrefix, username, password, f.Cookie[0])
}
