package service

import (
	"funnel/app/apis"
	"funnel/app/errors"
	"funnel/app/model"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type LibrarySystem struct {
}

type ReaderInfo struct {
	CurrentBorrowedCount string
	ExtendedCount        string
}
type Book struct {
	Name          string
	LibraryID     string
	LibraryPlace  string
	Time          string
	ReturnTime    string
	RenewalsTimes string
	IsExtended    string
}

func (t *LibrarySystem) GetBorrowHistory(user *model.User) []Book {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}
	requestUrl := apis.LibraryBorrowHistory
	request, _ := http.NewRequest("GET", requestUrl, nil)
	request.AddCookie(&user.Session)
	response, _ := client.Do(request)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var books []Book
	doc.Find("tr[onmouseout=\"this.style.backgroundColor=c\"]").Each(func(i int, s *goquery.Selection) {
		onmouseout, _ := s.Attr("onmouseout")
		if onmouseout == "this.style.backgroundColor=c" {
			bookName := strings.Trim(s.Find("a").Text(), " \n\r")
			bookId := strings.Trim(s.Find("td").Last().Prev().Prev().Text(), " \r\n")
			borrowTime := strings.Trim(s.Find("td").Last().Prev().Text(), " \r\n")
			book := Book{Name: bookName, LibraryID: bookId, Time: borrowTime}
			books = append(books, book)
		}
	})
	return books
}

func (t *LibrarySystem) GetCurrentBorrow(user *model.User) []Book {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}
	requestUrl := apis.LibraryBorrowing
	request, _ := http.NewRequest("GET", requestUrl, nil)
	request.AddCookie(&user.Session)
	response, _ := client.Do(request)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var books []Book
	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		style, _ := s.Attr("style")
		if strings.Contains(style, "border-style") {
			bookName := strings.Trim(s.Find("a").Text(), " \r\n")
			token := s.Find("span").Nodes
			libraryID := strings.Trim(token[0].FirstChild.Data, " \r\n")
			libraryPlace := strings.Trim(token[1].FirstChild.Data, " \r\n")
			time := strings.Trim(token[2].FirstChild.Data, " \r\n")
			returnTime := strings.Trim(token[3].FirstChild.Data, " \r\n")
			renewalsTimes := strings.Trim(token[4].FirstChild.Data, " \r\n")
			isExtended := strings.Trim(token[5].FirstChild.Data, " \r\n")
			book := Book{
				Name:          bookName,
				LibraryID:     libraryID,
				LibraryPlace:  libraryPlace,
				Time:          time,
				ReturnTime:    returnTime,
				RenewalsTimes: renewalsTimes,
				IsExtended:    isExtended}
			books = append(books, book)
		}

	})
	return books
}

func (t *LibrarySystem) login(username string, password string) (*model.User, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	url0 := apis.LibraryLogin
	loginData := url.Values{
		"__VIEWSTATE":          {apis.Library__VIEWSTATE},
		"__VIEWSTATEGENERATOR": {apis.Library__VIEWSTATEGENERATOR},
		"__EVENTVALIDATION":    {apis.Library__EVENTVALIDATION},
		"TextBox1":             {username},
		"TextBox2":             {password},
		"ImageButton1.x":       {"29"},
		"ImageButton1.y":       {"8"}}
	response, _ := client.PostForm(url0, loginData)

	if len(response.Cookies()) == 0 {
		return nil, errors.ERR_WRONG_PASSWORD
	}
	cookie := response.Cookies()[0]
	return SetUser(LibraryPrefix, username, password, cookie)
}

func (t *LibrarySystem) GetUser(username string, password string) (*model.User, error) {
	user, err := GetUser(LibraryPrefix, username, password)
	if err != nil {
		return t.login(username, password)
	}
	return user, err
}
