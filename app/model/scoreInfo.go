package model

func TransformScoreInfo(raw *ScoreRawInfo) ScoreInfo {
	var scoreInfo ScoreInfo
	for _, value := range raw.Items {
		scoreInfo = append(scoreInfo,
			&Score{
				Score:       value.Cj,
				ScorePoint:  value.Jd,
				LessonName:  value.Kcmc,
				LessonID:    value.Kch,
				TeacherName: value.Jsxm,
				ClassName:   value.Jxbmc,
				Credits:     value.Xf,
				SubmitTime:  value.Tjsj,
				SubmitName:  value.Tjrxm,
				LessonType:  value.Kcxzmc,
				ExamType:    value.Ksxz,
				SchoolTerm:  value.Xqmmc,
				SchoolYear:  value.Xnmmc,
				Key:         value.Key,
			})
	}

	return scoreInfo
}

type Score struct {
	Score       string `json:"score"`
	ScorePoint  string `json:"scorePoint"`
	TeacherName string `json:"teacherName"`
	LessonID    string `json:"lessonID"`
	LessonName  string `json:"lessonName"`
	ClassName   string `json:"className"`
	Credits     string `json:"credits"`
	LessonType  string `json:"lessonType"`
	ExamType    string `json:"examType"`
	SubmitTime  string `json:"submitTime"`
	SubmitName  string `json:"submitName"`
	SchoolTerm  string `json:"schoolTerm"`
	SchoolYear  string `json:"schoolYear"`
	Key         string `json:"key"`
}

type ScoreInfo []*Score

type ScoreRawInfo struct {
	Items []*struct {
		Cj     string
		Jsxm   string
		Kch    string
		Jxb_id string
		Jxbmc  string
		Kkbmmc string
		Kcxzmc string
		Khfsmc string
		Ksxz   string
		Kcmc   string
		Jd     string
		Xf     string
		Tjrxm  string
		Tjsj   string
		Key    string
		Xnmmc  string
		Xqmmc  string
	}
}
