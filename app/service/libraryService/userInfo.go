package libraryService

import (
	"funnel/app/apis/library"
	"funnel/app/service/libraryService/request"
	"net/http"
)

type UserInfo struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrCode   int         `json:"errCode"`
	ErrorCode interface{} `json:"errorCode"`
	Data      interface{} `json:"data"`
}

func GetUserInfo(cookies []*http.Cookie) (UserInfo, error) {
	var userInfo UserInfo
	session := request.New()
	_, err := session.Request().
		SetCookies(cookies).
		SetResult(&userInfo).
		Post(library.UserInfo)

	return userInfo, err
}

// CheckCookie 判断cookie是否还有效
func CheckCookie(cookies []*http.Cookie) bool {
	userInfo, err := GetUserInfo(cookies)
	if err != nil {
		return false
	}
	return userInfo.Success
}
