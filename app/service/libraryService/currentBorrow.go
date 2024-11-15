package libraryService

import (
	"funnel/app/apis/library"
	"github.com/go-resty/resty/v2"
)

// GetCurrentBorrow 获取图书馆当前借书记录
func GetCurrentBorrow(username string, password string, page int) (interface{}, error) {
	var ret Result
	cookies, err := OAuthLogin(username, password)
	if err != nil {
		return nil, err
	}
	client := resty.New()
	_, err = client.R().
		SetBody(map[string]interface{}{
			"page":          page,
			"rows":          10,
			"searchType":    1,
			"searchContent": "",
			"sortType":      0,
			"startDate":     nil,
			"endDate":       nil}).
		SetCookies(cookies).
		SetResult(&ret).
		Post(library.CurrentLoanList)
	return ret.Data, nil
}
