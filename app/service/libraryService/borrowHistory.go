package libraryService

import (
	"funnel/app/apis/library"
	"funnel/app/errors"
	"funnel/app/service/libraryService/request"
)

// GetBorrowHistory 获取图书馆当前借书记录
func GetBorrowHistory(username string, password string, page int) (interface{}, error) {
	var ret Result
	cookies, err := OAuthLogin(username, password)
	if err != nil {
		return nil, errors.ERR_WRONG_PASSWORD
	}

	client := request.New()

	_, err = client.Request().
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
		Post(library.BorrowHistoryList)
	return ret.Data.SearchResult, nil
}
