package libraryService

import (
	"funnel/app/apis/library"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type UserInfo struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrCode   int         `json:"errCode"`
	ErrorCode interface{} `json:"errorCode"`
	Data      struct {
		LastLibCode     interface{} `json:"lastLibCode"`
		BranchCode      interface{} `json:"branchCode"`
		DeskId          interface{} `json:"deskId"`
		MainFlag        interface{} `json:"mainFlag"`
		CxuId           interface{} `json:"cxuId"`
		UserId          int         `json:"userId"`
		PrimaryId       string      `json:"primaryId"`
		Name            string      `json:"name"`
		PicUrl          interface{} `json:"picUrl"`
		JobDesc         interface{} `json:"jobDesc"`
		JobType         interface{} `json:"jobType"`
		Password        string      `json:"password"`
		Unit            interface{} `json:"unit"`
		Department      interface{} `json:"department"`
		Position        interface{} `json:"position"`
		Gender          string      `json:"gender"`
		PreferredLang   string      `json:"preferredLang"`
		Email           interface{} `json:"email"`
		Address         interface{} `json:"address"`
		Phone           string      `json:"phone"`
		PostalCode      interface{} `json:"postalCode"`
		BirthDate       interface{} `json:"birthDate"`
		Status          string      `json:"status"`
		RegDate         int64       `json:"regDate"`
		ExpirationDate  int64       `json:"expirationDate"`
		PurgeDate       interface{} `json:"purgeDate"`
		StatusDate      interface{} `json:"statusDate"`
		LibCode         string      `json:"libCode"`
		GroupCode       string      `json:"groupCode"`
		CreateBy        interface{} `json:"createBy"`
		CreateDate      interface{} `json:"createDate"`
		UpdateBy        interface{} `json:"updateBy"`
		UpdateDate      interface{} `json:"updateDate"`
		AddType         interface{} `json:"addType"`
		BatchAddId      interface{} `json:"batchAddId"`
		CollegeYear     interface{} `json:"collegeYear"`
		CollegeClass    interface{} `json:"collegeClass"`
		CollegeDept     interface{} `json:"collegeDept"`
		VendorId        interface{} `json:"vendorId"`
		EmailCheckFlag  interface{} `json:"emailCheckFlag"`
		MobileCheckFlag interface{} `json:"mobileCheckFlag"`
	} `json:"data"`
}

func GetUserInfo(cookies []*http.Cookie) (UserInfo, error) {
	var userInfo UserInfo
	client := resty.New()
	_, err := client.R().
		EnableTrace().
		SetCookies(cookies).
		SetResult(&userInfo).
		Post(library.UserInfo)
	return userInfo, err
}

func CheckCookie(cookies []*http.Cookie) bool {
	userInfo, err := GetUserInfo(cookies)
	if err != nil {
		return false
	}
	return userInfo.Success
}
