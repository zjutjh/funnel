package schoolcardService

import (
	"funnel/app/model"
	"github.com/PuerkitoBio/goquery"
	"net/url"
)

func genLoginData(doc *goquery.Document, username, password string) url.Values {
	code := string(doc.Find("#UserLogin_ImgFirst").AttrOr("src", "0")[7])
	code += string(doc.Find("#UserLogin_imgSecond").AttrOr("src", "0")[7])
	code += string(doc.Find("#UserLogin_imgThird").AttrOr("src", "0")[7])
	code += string(doc.Find("#UserLogin_imgFour").AttrOr("src", "0")[7])

	return url.Values{
		"__VIEWSTATE":              {doc.Find("#__VIEWSTATE").AttrOr("value", "")},
		"__VIEWSTATEGENERATOR":     {doc.Find("#__VIEWSTATEGENERATOR").AttrOr("value", "")},
		"__EVENTVALIDATION":        {doc.Find("#__EVENTVALIDATION").AttrOr("value", "")},
		"UserLogin:txtUser":        {username},
		"UserLogin:txtPwd":         {password},
		"UserLogin:ddlPerson":      {"\xBF\xA8\xBB\xA7"},
		"UserLogin:txtSure":        {code},
		"UserLogin:ImageButton1.x": {"48"},
		"UserLogin:ImageButton1.y": {"8"}}
}
func genHistoryForm(doc *goquery.Document, year string, month string) url.Values {
	return url.Values{
		"__VIEWSTATE":          {doc.Find("#__VIEWSTATE").AttrOr("value", "")},
		"__VIEWSTATEGENERATOR": {doc.Find("#__VIEWSTATEGENERATOR").AttrOr("value", "")},
		"ddlYear":              {year},
		"txtMonth":             {month},
		"ddlMonth":             {month},
		"ImageButton1.x":       {"48"},
		"ImageButton1.y":       {"8"}}
}
func genTodayCardTransaction(doc *goquery.Document) []model.CardTransaction {
	var cardTransactions []model.CardTransaction
	cardTransactions = append(cardTransactions)
	doc.Find("#dgShow").Find("tr").Next().Each(func(i int, selection *goquery.Selection) {
		nodes := selection.Find("td").Nodes
		cardTransaction := model.CardTransaction{
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

func genHistoryCardTransaction(doc *goquery.Document) []model.CardTransaction {
	var cardTransactions []model.CardTransaction
	cardTransactions = append(cardTransactions)
	doc.Find("#dgShow").Find("tr").Next().Each(func(i int, selection *goquery.Selection) {
		nodes := selection.Find("td").Nodes
		cardTransaction := model.CardTransaction{
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
