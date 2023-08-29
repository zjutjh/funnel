package zfService

import (
	"encoding/json"
	"funnel/app/apis/zf"
	"funnel/app/controller"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"funnel/app/utils/fetch"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/url"
	"sort"
)

func GetLessonsTable(stu *model.User, year string, term string) (interface{}, error) {
	res, err := fetchTermRelatedInfo(stu, zf.ZfClassTable()+stu.Username, year, term, -1)
	if err != nil {
		return nil, err
	}
	var f model.LessonsTableRawInfo
	err = json.Unmarshal([]byte(res), &f)
	return model.TransformLessonTable(&f), err
}
func GetExamInfo(stu *model.User, year string, term string) (interface{}, error) {
	var result model.ExamInfo
	resultMap := make(map[string]*model.Exam)
	for i := 0; i < 7; i++ {
		res, err := fetchTermRelatedInfo(stu, zf.ZfExamInfo(), year, term, i)
		if err != nil {
			return nil, err
		}
		var f model.ExamRawInfo
		err = json.Unmarshal([]byte(res), &f)
		if err != nil {
			continue
			//return nil, err
		}
		examInfo := model.TransformExamInfo(&f)
		for _, v := range examInfo {
			resultMap[v.ExamTime] = v
		}
	}
	for _, v := range resultMap {
		result = append(result, v)
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].ExamTime > result[j].ExamTime
	})
	return result, nil
}
func GetScoreDetail(stu *model.User, year string, term string) (interface{}, error) {
	res, err := fetchTermRelatedInfo(stu, zf.ZfScoreDetail(), year, term, -1)
	if err != nil {
		return nil, err
	}
	var f model.ScoreDetailRawInfo
	err = json.Unmarshal([]byte(res), &f)
	return model.TransformScoreDetailInfo(&f), err
}
func GetScore(stu *model.User, year string, term string) (interface{}, error) {
	res, err := fetchTermRelatedInfo(stu, zf.ZfScore(), year, term, -1)
	if err != nil {
		return nil, err
	}
	var f model.ScoreRawInfo
	err = json.Unmarshal([]byte(res), &f)
	return model.TransformScoreInfo(&f), err
}
func GetMidTermScore(stu *model.User, year string, term string) (interface{}, error) {
	res, err := fetchTermRelatedInfo(stu, zf.ZfMinTermScore(), year, term, -1)
	if err != nil {
		return nil, err
	}
	var f model.MidTermScoreRawInfo
	err = json.Unmarshal([]byte(res), &f)
	return model.TransformMidTermScoreInfo(&f), err
}

func fetchTermRelatedInfo(stu *model.User, requestUrl, year, term string, examIndex int) (string, error) {
	f := fetch.Fetch{}
	f.Init()
	f.Cookie = append(f.Cookie, &stu.Session)
	if term == "上" {
		term = "3"
	} else if term == "下" {
		term = "12"
	} else if term == "短" {
		term = "16"
	}
	var requestData url.Values
	if examIndex != -1 {
		// 因正方考试信息查询需要一个单独的参数来判断考试类型，所以需要examIndex来进行标识，为-1表示此次查询不是查考试信息
		requestData = genTermExamInfoReqData(year, term, examIndex)
	} else {
		requestData = genTermRelatedInfoReqData(year, term)
	}
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
	if term == "上" {
		term = "3"
	} else if term == "下" {
		term = "12"
	} else if term == "短" {
		term = "16"
	}
	if campus == "朝晖" {
		campus = "01"
	} else if campus == "屏峰" {
		campus = "02"
	} else if campus == "莫干山" {
		campus = "A61400B98155D41AE0550113465EF1CF"
	}
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

func GetUser(username, password string, loginType controller.LoginType, typeFlag bool) (*model.User, error) {
	user, err := service.GetUser(service.ZFPrefix, username, password)
	if err != nil || typeFlag {
		switch loginType {
		case controller.ZF:
			return login(username, password)
		case controller.OAUTH:
			return loginByOauth(username, password)
		}
	}
	return user, err
}
