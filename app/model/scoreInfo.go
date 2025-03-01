package model

func TransformScoreInfo(raw *ScoreRawInfo) ScoreInfo {
	var scoreInfo ScoreInfo
	for _, value := range raw.Items {
		scoreInfo = append(scoreInfo,
			&Score{
				Score:       value.Cj,     // 成绩
				ScorePoint:  value.Jd,     // 绩点
				LessonName:  value.Kcmc,   // 课程名称
				LessonID:    value.Kch,    // 课程id
				TeacherName: value.Jsxm,   // 教师姓名
				ClassName:   value.Jxbmc,  // 课程名字
				Credits:     value.Xf,     // 学分
				SubmitTime:  value.Tjsj,   // 提交时间
				SubmitName:  value.Tjrxm,  // 提交名称
				LessonType:  value.Kcxzmc, // 课程类型
				ExamType:    value.Ksxz,   // 考试类型
				SchoolTerm:  value.Xqmmc,  // 学期
				SchoolYear:  value.Xnmmc,  // 学年
				Key:         value.Key,    // id
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
