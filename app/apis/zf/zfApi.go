package zf

import (
	"funnel/app/apis"
	"funnel/config"
	"strconv"
	"strings"
	"time"
)

func ChooseURL() string {
	if config.Redis.Exists("zf_url").Val() != 1 {
		config.Redis.Set("zf_url", "bk", 0)
	}
	if strings.Compare(config.Redis.Get("zf_url").String(), "new") == 0 {
		return apis.ZF_URL
	} else {
		return apis.ZF_BK_URL
	}
}

func ZfLoginGetPublickey() string {
	return ChooseURL() + "xtgl/login_getPublicKey.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
}
func ZfLoginHome() string {
	return ChooseURL() + "xtgl/login_slogin.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
}
func ZfLoginKaptcha() string {
	return ChooseURL() + "kaptcha?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
}
func ZfExamInfo() string {
	return ChooseURL() + "kwgl/kscx_cxXsksxxIndex.html?doType=query"
}
func ZfClassTable() string {
	return ChooseURL() + "kbcx/xskbcx_cxXsgrkb.html?gnmkdm=N2151&su="
}
func ZfScore() string {
	return ChooseURL() + "cjcx/cjcx_cxDgXscj.html?doType=query&gnmkdm=N305005"
}
func ZfScoreDetail() string {
	return ChooseURL() + "cjcx/cjcx_cxXsKccjList.html?doType=query"
}
func ZfMinTermScore() string {
	return ChooseURL() + "design/funcData_cxFuncDataList.html?func_widget_guid=5EF567BFD3CE243EE053A11310AC1252&gnmkdm=N305013"
}
func ZfEmptyClassRoom() string {
	return ChooseURL() + "cdjy/cdjy_cxKxcdlb.html?doType=query"
}
func ZfUserInfo() string {
	return ChooseURL() + "xsxxxggl/xsgrxxwh_cxXsgrxx.html?gnmkdm=N100801&layout=default"
}
func ZfPY() string {
	return ChooseURL() + "pyfagl/pyfaxxck_dyPyfaxx.html?id="
}
