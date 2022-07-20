package libraryService

import (
	"funnel/app/model"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
)

func genLoginData(doc *goquery.Document, username, password string) url.Values {
	return url.Values{
		"__VIEWSTATE":          {doc.Find("#__VIEWSTATE").AttrOr("value", "")},
		"__VIEWSTATEGENERATOR": {doc.Find("#__VIEWSTATEGENERATOR").AttrOr("value", "")},
		"__EVENTVALIDATION":    {doc.Find("#__EVENTVALIDATION").AttrOr("value", "")},
		"TextBox1":             {username},
		"TextBox2":             {password},
		"ImageButton1.x":       {"29"},
		"ImageButton1.y":       {"8"}}
}

func genBorrowHistoryForm(doc *goquery.Document) url.Values {
	return url.Values{
		"__EVENTTARGET":                {"ctl00$ContentPlaceHolder1$GridView1"},
		"__EVENTARGUMENT":              {"Page$Next"},
		"ctl00_TreeView1_ExpandState":  {doc.Find("#ctl00_TreeView1_ExpandState").AttrOr("value", "")},
		"ctl00_TreeView1_SelectedNode": {doc.Find("#ctl00_TreeView1_SelectedNode").AttrOr("value", "")},
		"ctl00_TreeView1_PopulateLog":  {doc.Find("#ctl00_TreeView1_PopulateLog").AttrOr("value", "")},
		"__VIEWSTATE":                  {doc.Find("#__VIEWSTATE").AttrOr("value", "")},
		"__VIEWSTATEGENERATOR":         {doc.Find("#__VIEWSTATEGENERATOR").AttrOr("value", "")},
		"__EVENTVALIDATION":            {doc.Find("#__EVENTVALIDATION").AttrOr("value", "")}}
}

func genCurrentBorrowFormFirst(doc *goquery.Document) url.Values {
	return url.Values{
		"ctl00$ScriptManager1":         {"ctl00$ContentPlaceHolder1$UpdatePanel1|ctl00$ContentPlaceHolder1$GridView1"},
		"__EVENTTARGET":                {"ctl00$ContentPlaceHolder1$GridView1"},
		"__EVENTARGUMENT":              {"Page$Next"},
		"ctl00_TreeView1_ExpandState":  {"eennnnnnennnnnnnennnnennen"},
		"ctl00_TreeView1_SelectedNode": {""},
		"ctl00_TreeView1_PopulateLog":  {""},
		"__LASTFOCUS":                  {doc.Find("#__LASTFOCUS").AttrOr("value", "")},
		"__VIEWSTATE":                  {doc.Find("#__VIEWSTATE").AttrOr("value", "")},
		"__VIEWSTATEGENERATOR":         {doc.Find("#__VIEWSTATEGENERATOR").AttrOr("value", "")},
		"__VIEWSTATEENCRYPTED":         {doc.Find("#__VIEWSTATEENCRYPTED").AttrOr("value", "")},
		"__EVENTVALIDATION":            {doc.Find("#__EVENTVALIDATION").AttrOr("value", "")},
		"__ASYNCPOST":                  {"true"}}
}

func genCurrentBorrowForm(doc *goquery.Document) url.Values {
	var value []string
	hidden := strings.Split(doc.Text(), "|hiddenField|")
	for i, v := range hidden {
		if i != 0 {
			str := strings.Split(v, "|")
			value = append(value, str[1])
		}
	}
	return url.Values{
		"ctl00$ScriptManager1":         {"ctl00$ContentPlaceHolder1$UpdatePanel1|ctl00$ContentPlaceHolder1$GridView1"},
		"__EVENTTARGET":                {"ctl00$ContentPlaceHolder1$GridView1"},
		"__EVENTARGUMENT":              {"Page$Next"},
		"ctl00_TreeView1_ExpandState":  {"eennnnnnennnnnnnennnnennen"},
		"ctl00_TreeView1_SelectedNode": {""},
		"ctl00_TreeView1_PopulateLog":  {""},
		"__LASTFOCUS":                  {value[2]},
		"__VIEWSTATE":                  {value[3]},
		"__VIEWSTATEGENERATOR":         {value[4]},
		"__VIEWSTATEENCRYPTED":         {value[5]},
		"__EVENTVALIDATION":            {value[6]},
		"__ASYNCPOST":                  {"true"}}
}

func genReBorrowForm(doc *goquery.Document, tdID string) url.Values {
	return url.Values{
		"__VIEWSTATE":                       {doc.Find("#__VIEWSTATE").AttrOr("value", "")},
		"__VIEWSTATEGENERATOR":              {doc.Find("#__VIEWSTATEGENERATOR").AttrOr("value", "")},
		"__VIEWSTATEENCRYPTED":              {""},
		"__EVENTVALIDATION":                 {doc.Find("#__EVENTVALIDATION").AttrOr("value", "")},
		"ctl00$ContentPlaceHolder1$XuJieBt": {"续借"},
		"ctl00$ContentPlaceHolder1$GridView1$" + tdID + "$CheckBox1": {"on"},
	}
}

func getReBorrowedBookInfoIDFromDoc(doc *goquery.Document, bookid string) string {
	tkID := ""
	doc.Find("#ctl00_ContentPlaceHolder1_GridView1").Each(func(i int, s *goquery.Selection) {
		s.Find("table").Each(func(i int, s *goquery.Selection) {
			bookName := strings.Trim(s.Find("a").Text(), " \r\n")
			if bookName == "" {
				return
			}
			token := s.Find("span").Nodes
			libraryID := strings.Trim(token[0].FirstChild.Data, " \r\n")
			if bookid == libraryID {
				tok := strings.Split(token[0].Attr[0].Val, "_")
				id := tok[len(tok)-2]
				tkID = id
				return
			}
		})
	})
	return tkID
}
func genBorrowedBookInfoFromDoc(doc *goquery.Document) ([]model.BorrowedBookInfo, bool) {
	var books []model.BorrowedBookInfo
	doc.Find("#ctl00_ContentPlaceHolder1_GridView1").Each(func(i int, s *goquery.Selection) {
		s.Find("table").Each(func(i int, s *goquery.Selection) {
			bookName := strings.Trim(s.Find("a").Text(), " \r\n")
			if bookName == "" {
				return
			}
			token := s.Find("span").Nodes
			libraryID := strings.Trim(token[0].FirstChild.Data, " \r\n")
			libraryPlace := strings.Trim(token[1].FirstChild.Data, " \r\n")
			renewalsTimes := strings.Trim(token[2].FirstChild.Data, " \r\n")
			time := strings.Trim(token[3].FirstChild.Data, " \r\n")
			returnTime := strings.Trim(token[4].FirstChild.Data, " \r\n")

			isExtended := strings.Trim(token[5].FirstChild.Data, " \r\n")
			book := model.BorrowedBookInfo{
				Name:          bookName,
				LibraryID:     libraryID,
				LibraryPlace:  libraryPlace,
				Time:          time,
				ReturnTime:    returnTime,
				RenewalsTimes: renewalsTimes,
				OverdueTime:   isExtended}
			books = append(books, book)
		})
	})
	html, _ := doc.Html()
	return books, strings.Contains(html, "pic/NextPage.png")
}

func genBorrowedBookHistoryInfoFromDoc(doc *goquery.Document) ([]model.BorrowedBookInfo, bool) {
	var books []model.BorrowedBookInfo
	doc.Find("tr[onmouseout=\"this.style.backgroundColor=c\"]").Each(func(i int, s *goquery.Selection) {
		onmouseout, _ := s.Attr("onmouseout")
		if onmouseout == "this.style.backgroundColor=c" {
			bookName := strings.Trim(s.Find("a").Text(), " \n\r")
			bookId := strings.Trim(s.Find("td").Last().Prev().Prev().Text(), " \r\n")
			borrowTime := strings.Trim(s.Find("td").Last().Prev().Text(), " \r\n")
			returnTime := strings.Trim(s.Find("td").Last().Text(), " \r\n")
			book := model.BorrowedBookInfo{Name: bookName, LibraryID: bookId, Time: borrowTime, ReturnTime: returnTime}
			books = append(books, book)
		}
	})
	html, _ := doc.Html()
	return books, strings.Contains(html, "pic/NextPage.png")
}
