package zfService

import (
	"funnel/app/apis/zf"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"funnel/app/utils/fetch"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
)

func GetClassTable(stu *model.User, year string, term string) (string, error) {
	return fetchTermRelatedInfo(stu, zf.ZfClassTable(), year, term)
}
func GetExamInfo(stu *model.User, year string, term string) (string, error) {
	return fetchTermRelatedInfo(stu, zf.ZfExamInfo(), year, term)
}
func GetScoreDetail(stu *model.User, year string, term string) (string, error) {
	return fetchTermRelatedInfo(stu, zf.ZfScoreDetail(), year, term)
}
func GetScore(stu *model.User, year string, term string) (string, error) {
	return fetchTermRelatedInfo(stu, zf.ZfScore(), year, term)
}

func fetchTermRelatedInfo(stu *model.User, requestUrl, year, term string) (string, error) {
	f := fetch.Fetch{}
	f.Init()
	f.Cookie = append(f.Cookie, &stu.Session)
	requestData := genTermRelatedInfoReqData(year, term)
	s, err := f.PostForm(requestUrl, requestData)

	if len(s) == 0 {
		service.ForgetUserByUsername(service.ZFPrefix, stu.Username)
		return "", errors.ERR_Session_Expired
	}
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func GetTrainingPrograms(stu *model.User) ([]byte, error) {
	f := fetch.Fetch{}
	f.Init()
	f.Cookie = append(f.Cookie, &stu.Session)
	response, err := f.GetRaw(zf.ZfUserInfo())

	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	s, exist := doc.Find("#pyfaxx_id").Attr("value")
	if exist {
		res, _ := f.GetRaw(zf.ZfPY() + s)
		s, _ := ioutil.ReadAll(res.Body)
		return s, nil
	}
	return nil, nil
}

func GetEmptyRoomInfo(stu *model.User, year string, term string, campus string, weekday string, week string, classPeriod string) (string, error) {
	f := fetch.Fetch{}
	f.Init()
	f.Cookie = append(f.Cookie, &stu.Session)
	requestData := genEmptyRoomReqData(year, term, campus, week, weekday, classPeriod)
	s, err := f.PostForm(zf.ZfEmptyClassRoom(), requestData)

	if len(s) == 0 {
		service.ForgetUserByUsername(service.ZFPrefix, stu.Username)
		return "", errors.ERR_Session_Expired
	}
	if err != nil {
		return "", err
	}

	return string(s), nil
}

func GetUser(username string, password string) (*model.User, error) {
	user, err := service.GetUser(service.ZFPrefix, username, password)
	if err != nil {
		return login(username, password)
	}
	return user, err
}
