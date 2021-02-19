package schoolcardService

import (
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
