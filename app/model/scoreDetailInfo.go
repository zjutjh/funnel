package model

func TransformScoreDetailInfo(raw *ScoreDetailRawInfo) ScoreDetailInfo {
	var scoreInfo ScoreDetailInfo
	for _, value := range raw.Items {
		scoreInfo = append(scoreInfo,
			&ScoreDetail{
				Score: value.Xmcj,

				LessonName: value.Kcmc,
				LessonID:   value.Kch,
				ClassName:  value.Jxbmc,
				Credits:    value.Xf,

				Name: value.Xmblmc,
			})
	}

	return scoreInfo
}

type ScoreDetail struct {
	Name       string `json:"name"`
	Score      string `json:"score"`
	LessonID   string `json:"lessonID"`
	LessonName string `json:"lessonName"`
	ClassName  string `json:"className"`
	Credits    string `json:"credits"`
}

type ScoreDetailInfo []*ScoreDetail

type ScoreDetailRawInfo struct {
	Items []*struct {
		Xmcj   string
		Jsxm   string
		Kch    string
		Jxb_id string
		Jxbmc  string
		Kkbmmc string
		Xmblmc string
		Khfsmc string
		Ksxz   string
		Kcmc   string
		Xf     string
	}
}
