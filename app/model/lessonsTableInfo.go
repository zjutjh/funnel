package model

func TransformLessonTable(raw *LessonsTableRawInfo) LessonsTableInfo {
	var lessonTable LessonsTableInfo

	lessonTable.Info.Name = raw.Xsxx.XM
	lessonTable.Info.ClassName = raw.Xsxx.BJMC
	for _, value := range raw.KbList {
		lessonTable.LessonsTable = append(lessonTable.LessonsTable,
			&Lesson{
				ID:          value.Kch_id,
				Sections:    value.Jcs,
				LessonName:  value.Kcmc,
				Campus:      value.Xqmc,
				LessonPlace: value.Cdmc,
				PlaceID:     value.Cd_id,
				TeacherName: value.Xm,
				ClassName:   value.Jxbmc,
				ClassID:     value.Jxb_id,
				Weekday:     value.Xqj,
				Week:        value.Zcd,
				LessonHours: value.Zxs,
				Credits:     value.Xf,
				Type:        value.Kcxz,
			})
	}

	for _, value := range raw.SjkList {
		lessonTable.PracticeLessons = append(lessonTable.PracticeLessons,
			&PracticeLesson{
				LessonName:  value.Kcmc,
				TeacherName: value.Jsxm,
				ClassName:   value.Qsjsz,
				Credits:     value.Xf,
			})
	}

	return lessonTable
}

type Lesson struct {
	ID          string `json:"id"`
	Sections    string `json:"sections"`
	LessonName  string `json:"lessonName"`
	Campus      string `json:"campus"`
	LessonPlace string `json:"lessonPlace"`
	PlaceID     string `json:"placeID"`
	TeacherName string `json:"teacherName"`
	ClassName   string `json:"className"`
	ClassID     string `json:"classID"`
	Weekday     string `json:"weekday"`
	Week        string `json:"week"`
	LessonHours string `json:"lessonHours"`
	Credits     string `json:"credits"`
	Type        string `json:"type"`
}

type PracticeLesson struct {
	LessonName  string `json:"lessonName"`
	TeacherName string `json:"teacherName"`
	ClassName   string `json:"className"`
	Credits     string `json:"credits"`
}

type LessonsTableInfo struct {
	LessonsTable    []*Lesson         `json:"lessonsTable"`
	PracticeLessons []*PracticeLesson `json:"practiceLessons"`
	Info            struct {
		Name      string
		ClassName string
	} `json:"info"`
}
type LessonsTableRawInfo struct {
	Xsxx struct {
		XM   string
		BJMC string
	}
	SjkList []*struct {
		Jsxm  string
		Kcmc  string
		Qsjsz string
		Xf    string
	}
	KbList []*struct {
		Kch_id string
		Jcs    string
		Kcmc   string
		Xqmc   string
		Cdmc   string
		Cd_id  string
		Xm     string
		Jxbmc  string
		Jxb_id string
		Xqj    string
		Zcd    string
		Zxs    string
		Xf     string
		Kcxz   string
	}
}
