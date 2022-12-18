package model

import "strings"

func TransformMidTermScoreInfo(raw *MidTermScoreRawInfo) MidTermScoreInfo {
	var midTermScoreInfo MidTermScoreInfo
	for _, value := range raw.Items {
		value.Jsxx = strings.Split(value.Jsxx, "/")[1]
		midTermScoreInfo = append(midTermScoreInfo,
			&MidTermScore{
				Score:       value.Xmcj,
				LessonName:  value.Kcmc,
				LessonID:    value.Kch,
				TeacherName: value.Jsxx,
				ClassName:   value.Jxbmc,
				Credits:     value.Xf,
			})
	}

	return midTermScoreInfo
}

type MidTermScore struct {
	Score       string `json:"score"`
	TeacherName string `json:"teacherName"`
	LessonID    string `json:"lessonID"`
	LessonName  string `json:"lessonName"`
	ClassName   string `json:"className"`
	Credits     string `json:"credits"`
}

type MidTermScoreInfo []*MidTermScore

type MidTermScoreRawInfo struct {
	Items []*struct {
		Xmcj   string
		Jsxx   string
		Kch    string
		Kch_id string
		Jxbmc  string
		Xsxy   string
		Kcmc   string
		Xf     string
	}
}
