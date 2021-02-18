package zf

import (
	"funnel/app/apis/zf"
	"funnel/app/model"
	"funnel/app/service"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetClassTable(stu *model.User, year string, term string) (string, error) {
	return fetchTermRelatedInfo(stu, zf.ZfClassTable, year, term)
}
func GetExamInfo(stu *model.User, year string, term string) (string, error) {
	return fetchTermRelatedInfo(stu, zf.ZfExamInfo, year, term)
}
func GetScoreDetail(stu *model.User, year string, term string) (string, error) {
	return fetchTermRelatedInfo(stu, zf.ZfScoreDetail, year, term)
}
func GetScore(stu *model.User, year string, term string) (string, error) {
	return fetchTermRelatedInfo(stu, zf.ZfScore, year, term)
}

func fetchTermRelatedInfo(stu *model.User, requestUrl, year, term string) (string, error) {

	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
	requestData := genTermRelatedInfoReqData(year, term)
	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(requestData.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.AddCookie(&stu.Session)
	response, err := client.Do(request)

	if err != nil {
		return "", err
	}

	s, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(s), nil
}

func GetTrainingPrograms(stu *model.User) ([]byte, error) {

	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
	request, _ := http.NewRequest("GET", zf.ZfUserInfo, nil)
	request.AddCookie(&stu.Session)
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	s, exist := doc.Find("#pyfaxx_id").Attr("value")
	if exist {
		request, _ := http.NewRequest("GET", zf.ZfPY+s, nil)
		request.AddCookie(&stu.Session)
		res, _ := client.Do(request)
		s, _ := ioutil.ReadAll(res.Body)
		return s, nil
	}
	return nil, nil
}

func GetEmptyRoomInfo(stu *model.User, year string, term string, campus string, weekday string, week string, classPeriod string) (string, error) {

	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
	requestData := genEmptyRoomReqData(year, term, campus, week, weekday, classPeriod)
	request, _ := http.NewRequest("POST", zf.ZfEmptyClassRoom, strings.NewReader(requestData.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.AddCookie(&stu.Session)
	response, err := client.Do(request)

	if err != nil {
		return "", err
	}

	s, err := ioutil.ReadAll(response.Body)
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
