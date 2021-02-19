package libraryService

import (
	"funnel/app/apis/library"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"funnel/app/utils/fetch"
)

func login(username string, password string) (*model.User, error) {
	f := fetch.Fetch{}
	f.Init()

	loginData := genLoginData(username, password)
	_, err := f.PostForm(library.LibraryLogin, loginData)

	if len(f.Cookie) == 0 || err != nil {
		return nil, errors.ERR_WRONG_PASSWORD
	}
	return service.SetUser(service.LibraryPrefix, username, password, f.Cookie[0])
}
