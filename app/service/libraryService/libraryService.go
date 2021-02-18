package libraryService

import (
	"funnel/app/apis/library"
	"funnel/app/model"
	"funnel/app/service"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func GetBorrowHistory(user *model.User) []model.BorrowedBookInfo {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}
	requestUrl := library.LibraryBorrowHistory
	request, _ := http.NewRequest("GET", requestUrl, nil)
	request.AddCookie(&user.Session)
	response, _ := client.Do(request)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var books []model.BorrowedBookInfo
	doc.Find("tr[onmouseout=\"this.style.backgroundColor=c\"]").Each(func(i int, s *goquery.Selection) {
		onmouseout, _ := s.Attr("onmouseout")
		if onmouseout == "this.style.backgroundColor=c" {
			bookName := strings.Trim(s.Find("a").Text(), " \n\r")
			bookId := strings.Trim(s.Find("td").Last().Prev().Prev().Text(), " \r\n")
			borrowTime := strings.Trim(s.Find("td").Last().Prev().Text(), " \r\n")
			book := model.BorrowedBookInfo{Name: bookName, LibraryID: bookId, Time: borrowTime}
			books = append(books, book)
		}
	})
	return books
}

func GetCurrentBorrow(user *model.User) []model.BorrowedBookInfo {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}
	requestUrl := library.LibraryBorrowing
	request, _ := http.NewRequest("GET", requestUrl, nil)
	request.AddCookie(&user.Session)
	response, _ := client.Do(request)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var books []model.BorrowedBookInfo
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
			book := model.BorrowedBookInfo{
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

func GetUser(username string, password string) (*model.User, error) {
	user, err := service.GetUser(service.LibraryPrefix, username, password)
	if err != nil {
		return login(username, password)
	}
	return user, err
}
