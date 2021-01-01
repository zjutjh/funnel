package service

import (
	"crypto/tls"
	"fmt"
	"funnel/app/apis"
	"funnel/app/errors"
	"funnel/app/helps"
	"funnel/app/model"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type CardSystem struct {
}
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

func (t *CardSystem) Login(user *model.CardUser) error {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Transport:     transport,
	}

	res, _ := client.Get(apis.CardHome)
	SSIONID := res.Cookies()[0]
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	code := string(doc.Find("#UserLogin_ImgFirst").AttrOr("src", "0")[7])
	code += string(doc.Find("#UserLogin_imgSecond").AttrOr("src", "0")[7])
	code += string(doc.Find("#UserLogin_imgThird").AttrOr("src", "0")[7])
	code += string(doc.Find("#UserLogin_imgFour").AttrOr("src", "0")[7])
	loginData := url.Values{
		"__VIEWSTATE":              {doc.Find("#__VIEWSTATE").AttrOr("value", "")},
		"__VIEWSTATEGENERATOR":     {doc.Find("#__VIEWSTATEGENERATOR").AttrOr("value", "")},
		"__EVENTVALIDATION":        {doc.Find("#__EVENTVALIDATION").AttrOr("value", "")},
		"UserLogin:txtUser":        {user.Username},
		"UserLogin:txtPwd":         {user.Password},
		"UserLogin:ddlPerson":      {"\xBF\xA8\xBB\xA7"},
		"UserLogin:txtSure":        {code},
		"UserLogin:ImageButton1.x": {"48"},
		"UserLogin:ImageButton1.y": {"8"}}
	request, _ := http.NewRequest("POST", apis.CardHome, strings.NewReader(loginData.Encode()))
	request.AddCookie(SSIONID)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, _ := client.Do(request)

	s, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(s))
	fmt.Println(response.StatusCode)
	if response.StatusCode != http.StatusFound {
		return errors.WrongPassword
	}
	SSIONID = res.Cookies()[0]
	session := SSIONID
	fmt.Println(SSIONID.Raw)
	user.Session = *session
	return nil
}

func (t *CardSystem) GetCurrentBalance(user *model.CardUser) string {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Transport:     transport,
	}

	request, _ := http.NewRequest("GET", apis.CardBalance, nil)
	request.AddCookie(&user.Session)
	response, _ := client.Do(request)

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	return doc.Find("#lblOne0").Text()
}

func (t *CardSystem) GetCardToday(user *model.CardUser) []CardTransaction {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Transport:     transport,
	}

	request, _ := http.NewRequest("GET", apis.CardToday, nil)
	request.AddCookie(&user.Session)
	response, _ := client.Do(request)
	utf8Body, err := helps.DecodeHTMLBody(response.Body, "GBK")
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
func (t *CardSystem) GetCardHistory(user *model.CardUser, year string, month string) []CardTransaction {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Transport:     transport,
	}

	request, _ := http.NewRequest("GET", apis.CardHistoryQuery, nil)
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

	request, _ = http.NewRequest("POST", apis.CardHistoryQuery, strings.NewReader(loginData.Encode()))

	request.AddCookie(&user.Session)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, _ = client.Do(request)
	utf8Body, err := helps.DecodeHTMLBody(response.Body, "GBK")
	doc, err = goquery.NewDocumentFromReader(utf8Body)

	request, _ = http.NewRequest("GET", apis.CardHistory, nil)
	request.AddCookie(&user.Session)
	response, _ = client.Do(request)
	utf8Body, err = helps.DecodeHTMLBody(response.Body, "GBK")
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
