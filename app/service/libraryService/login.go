package libraryService

import (
	"funnel/app/apis/library"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"net/http"
)

func login(username string, password string) (*model.User, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	loginData := genLoginData(username, password)
	response, _ := client.PostForm(library.LibraryLogin, loginData)

	if len(response.Cookies()) == 0 {
		return nil, errors.ERR_WRONG_PASSWORD
	}

	return service.SetUser(service.LibraryPrefix, username, password, response.Cookies()[0])
}
