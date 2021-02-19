package schoolcardService

import (
	"crypto/tls"
	"fmt"
	"funnel/app/apis/schoolcard"
	"funnel/app/model"
	"funnel/app/service"
	strings2 "funnel/app/utils/strings"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type CardTransaction struct {
	ID              string
	Account         string
	CardType        string
	TransactionType string
	Shop            string
	ShopPlace       string
	Terminal        string
	Transactions    string
	Time            string
	Wallet          string
	Balance         string
}

func GetCurrentBalance(user *model.User) string {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Transport:     transport,
	}

	request, _ := http.NewRequest("GET", schoolcard.CardBalance, nil)
	request.AddCookie(&user.Session)
	response, _ := client.Do(request)

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	return doc.Find("#lblOne0").Text()
}

func GetCardToday(user *model.User) []CardTransaction {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Transport:     transport,
	}

	request, _ := http.NewRequest("GET", schoolcard.CardToday, nil)
	request.AddCookie(&user.Session)
	response, _ := client.Do(request)
	utf8Body, err := strings2.DecodeHTMLBody(response.Body, "GBK")
	doc, err := goquery.NewDocumentFromReader(utf8Body)

	if err != nil {
		log.Fatal(err)
	}
	var cardTransactions []CardTransaction
	cardTransactions = append(cardTransactions)
	doc.Find("#dgShow").Find("tr").Next().Each(func(i int, selection *goquery.Selection) {
		nodes := selection.Find("td").Nodes
		cardTransaction := CardTransaction{
			nodes[0].FirstChild.Data,
			nodes[1].FirstChild.Data,
			nodes[2].FirstChild.Data,
			nodes[3].FirstChild.Data,
			nodes[4].FirstChild.Data,
			nodes[5].FirstChild.Data,
			nodes[6].FirstChild.Data,
			nodes[7].FirstChild.Data,
			nodes[8].FirstChild.Data,
			nodes[9].FirstChild.Data,
			nodes[10].FirstChild.Data}

		cardTransactions = append(cardTransactions, cardTransaction)
	})
	return cardTransactions
}
func GetCardHistory(user *model.User, year string, month string) []CardTransaction {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Transport:     transport,
	}

	request, _ := http.NewRequest("GET", schoolcard.CardHistoryQuery, nil)
	request.AddCookie(&user.Session)
	response, _ := client.Do(request)
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	loginData := url.Values{
		"__VIEWSTATE":          {doc.Find("#__VIEWSTATE").AttrOr("value", "")},
		"__VIEWSTATEGENERATOR": {doc.Find("#__VIEWSTATEGENERATOR").AttrOr("value", "")},
		"ddlYear":              {year},
		"txtMonth":             {month},
		"ddlMonth":             {month},
		"ImageButton1.x":       {"48"},
		"ImageButton1.y":       {"8"}}

	request, _ = http.NewRequest("POST", schoolcard.CardHistoryQuery, strings.NewReader(loginData.Encode()))

	request.AddCookie(&user.Session)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, _ = client.Do(request)
	utf8Body, err := strings2.DecodeHTMLBody(response.Body, "GBK")
	doc, err = goquery.NewDocumentFromReader(utf8Body)

	request, _ = http.NewRequest("GET", schoolcard.CardHistory, nil)
	request.AddCookie(&user.Session)
	response, _ = client.Do(request)
	utf8Body, err = strings2.DecodeHTMLBody(response.Body, "GBK")
	doc, err = goquery.NewDocumentFromReader(utf8Body)

	if err != nil {
		log.Fatal(err)
	}
	var cardTransactions []CardTransaction
	cardTransactions = append(cardTransactions)
	doc.Find("#dgShow").Find("tr").Next().Each(func(i int, selection *goquery.Selection) {
		nodes := selection.Find("td").Nodes
		cardTransaction := CardTransaction{
			nodes[0].FirstChild.Data,
			nodes[1].FirstChild.Data,
			nodes[2].FirstChild.Data,
			nodes[3].FirstChild.Data,
			nodes[4].FirstChild.Data,
			nodes[5].FirstChild.Data,
			nodes[6].FirstChild.Data,
			nodes[7].FirstChild.Data,
			nodes[8].FirstChild.Data,
			nodes[9].FirstChild.Data,
			nodes[10].FirstChild.Data}

		cardTransactions = append(cardTransactions, cardTransaction)
	})
	return cardTransactions
}

func GetUser(username string, password string) (*model.User, error) {
	user, err := service.GetUser(service.CardPrefix, username, password)
	if err != nil {
		fmt.Println(err.Error())
		return login(username, password)
	}
	return user, err
}
