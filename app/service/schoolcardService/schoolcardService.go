package schoolcardService

import (
	"funnel/app/apis/schoolcard"
	"funnel/app/model"
	"funnel/app/utils/fetch"
	strings2 "funnel/app/utils/strings"
	"github.com/PuerkitoBio/goquery"
)

func GetCurrentBalance(user *model.User) (string, error) {

	f := fetch.Fetch{}
	f.InitUnSafe()

	response, err := f.GetRaw(schoolcard.CardBalance)
	if err != nil {
		return "", err
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}
	return doc.Find("#lblOne0").Text(), nil
}

func GetCardToday(user *model.User) ([]model.CardTransaction, error) {
	f := fetch.Fetch{}
	f.InitUnSafe()

	response, err := f.GetRaw(schoolcard.CardToday)
	if err != nil {
		return nil, err
	}
	utf8Body, err := strings2.DecodeHTMLBody(response.Body, "GBK")
	doc, err := goquery.NewDocumentFromReader(utf8Body)

	if err != nil {
		return nil, err
	}

	cardTransactions := genTodayCardTransaction(doc)
	return cardTransactions, nil
}

func GetCardHistory(user *model.User, year string, month string) ([]model.CardTransaction, error) {
	f := fetch.Fetch{}
	f.InitUnSafe()

	response, err := f.GetRaw(schoolcard.CardHistoryQuery)
	if err != nil {
		return nil, err
	}
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	q := genHistoryForm(doc, year, month)

	response, err = f.PostFormRaw(schoolcard.CardHistoryQuery, q)
	if err != nil {
		return nil, err
	}
	utf8Body, err := strings2.DecodeHTMLBody(response.Body, "GBK")
	doc, err = goquery.NewDocumentFromReader(utf8Body)
	if err != nil {
		return nil, err
	}
	response, err = f.GetRaw(schoolcard.CardHistory)
	if err != nil {
		return nil, err
	}
	utf8Body, err = strings2.DecodeHTMLBody(response.Body, "GBK")
	doc, err = goquery.NewDocumentFromReader(utf8Body)

	if err != nil {
		return nil, err
	}

	cardTransactions := genHistoryCardTransaction(doc)
	return cardTransactions, nil
}
