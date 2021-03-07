package libraryService

import (
	"funnel/app/apis/library"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"funnel/app/utils/fetch"
	"github.com/PuerkitoBio/goquery"
)

func GetBorrowHistory(user *model.User) ([]model.BorrowedBookInfo, error) {
	f := fetch.Fetch{}
	f.Init()
	f.Cookie = append(f.Cookie, &user.Session)
	response, err := f.GetRaw(library.LibraryBorrowHistory)
	if response != nil && response.StatusCode != 200 {
		service.ForgetUserByUsername(service.LibraryPrefix, user.Username)
		return nil, errors.ERR_Session_Expired
	}

	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	books := genBorrowedBookHistoryInfoFromDoc(doc)
	return books, nil
}

func GetCurrentBorrow(user *model.User) ([]model.BorrowedBookInfo, error) {
	f := fetch.Fetch{}
	f.Init()
	f.Cookie = append(f.Cookie, &user.Session)
	response, err := f.GetRaw(library.LibraryBorrowing)
	if response != nil && response.StatusCode != 200 {
		service.ForgetUserByUsername(service.LibraryPrefix, user.Username)
		return nil, errors.ERR_Session_Expired
	}
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		return nil, err
	}

	books := genBorrowedBookInfoFromDoc(doc)
	return books, nil
}
