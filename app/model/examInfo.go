package model

func TransformExamInfo(raw *ExamRawInfo) ExamInfo {
	var examInfo ExamInfo
	for _, value := range raw.Items {
		examInfo = append(examInfo,
			&Exam{
				LessonID:    value.Kch,
				LessonName:  value.Kcmc,
				Campus:      value.Cdxqmc,
				LessonPlace: value.Jxdd,
				TeacherName: value.Jsxx,
				ClassName:   value.Jxbmc,
				Credits:     value.Xf,
				ExamPlace:   value.Cdmc,
				ExamTime:    value.Kssj,
			})
	}

	return examInfo
}

type Exam struct {
	LessonID    string `json:"id"`
	LessonName  string `json:"lessonName"`
	LessonPlace string `json:"lessonPlace"`
	ExamPlace   string `json:"examPlace"`
	ExamTime    string `json:"examTime"`
	Campus      string `json:"campus"`
	TeacherName string `json:"teacherName"`
	ClassName   string `json:"className"`
	Credits     string `json:"credits"`
}

type ExamInfo []*Exam

type ExamRawInfo struct {
	Items []*struct {
		Kch    string
		Jxdd   string
		Jxbmc  string
		Kcmc   string
		Kcxz   string
		Kssj   string
		Kcxszc string
		Jsxx   string
		Cdxqmc string
		Xf     string
		Cdmc   string
	}
}
