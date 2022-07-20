package libraryService

import (
	"funnel/app/apis/library"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"funnel/app/utils/fetch"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
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

	books, next := genBorrowedBookHistoryInfoFromDoc(doc)
	for next {
		form := genBorrowHistoryForm(doc)
		response, _ = f.PostFormRaw(library.LibraryBorrowHistory, form)
		doc, err = goquery.NewDocumentFromReader(response.Body)
		booksNew, nextNew := genBorrowedBookHistoryInfoFromDoc(doc)
		books = append(books, booksNew...)
		next = nextNew
	}
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

	books, next := genBorrowedBookInfoFromDoc(doc)

	var form = url.Values{}
	if next {
		form = genCurrentBorrowFormFirst(doc)
		response, _ = f.PostFormRawAsynchronous(library.LibraryBorrowing, form)
		doc, err = goquery.NewDocumentFromReader(response.Body)
		booksNew, nextNew := genBorrowedBookInfoFromDoc(doc)
		books = append(books, booksNew...)
		next = nextNew
	}
	for next {
		form = genCurrentBorrowForm(doc)
		response, _ = f.PostFormRawAsynchronous(library.LibraryBorrowing, form)
		doc, err = goquery.NewDocumentFromReader(response.Body)
		booksNew, nextNew := genBorrowedBookInfoFromDoc(doc)
		books = append(books, booksNew...)
		next = nextNew
	}
	return books, nil
}

func DoReBorrow(user *model.User, bookid string) error {
	f := fetch.Fetch{}
	f.Init()
	f.Cookie = append(f.Cookie, &user.Session)
	response, err := f.GetRaw(library.LibraryBorrowing)
	if response != nil && response.StatusCode != 200 {
		service.ForgetUserByUsername(service.LibraryPrefix, user.Username)
		return errors.ERR_Session_Expired
	}
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		return err
	}

	tkID := getReBorrowedBookInfoIDFromDoc(doc, bookid)

	if tkID == "" {
		return errors.ERR_INVALID_ARGS
	}
	q := genReBorrowForm(doc, tkID)
	print(q.Encode())
	raw, err := f.PostForm(library.LibraryBorrowing, q)
	if err != nil {
		return err
	}
	print(string(raw))
	if strings.Contains(string(raw), "续借成功!") {
		return nil
	}
	return errors.ERR_INVALID_ARGS
}
