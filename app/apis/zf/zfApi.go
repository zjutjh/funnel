package zf

import (
	"funnel/app/apis"
	"strconv"
	"time"
)

func ZfLoginGetPublickey() string {
	return apis.ZF_URL + "xtgl/login_getPublicKey.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
}
func ZfLoginHome() string {
	return apis.ZF_URL + "xtgl/login_slogin.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
}
func ZfLoginKaptcha() string {
	return apis.ZF_URL + "kaptcha?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
}
func ZfExamInfo() string {
	return apis.ZF_URL + "/kwgl/kscx_cxXsksxxIndex.html?doType=query"
}
func ZfClassTable() string {
	return apis.ZF_URL + "kbcx/xskbcx_cxXsKb.html?doType=query"
}
func ZfScore() string {
	return apis.ZF_URL + "cjcx/cjcx_cxDgXscj.html?doType=query"
}
func ZfScoreDetail() string {
	return apis.ZF_URL + "cjcx/cjcx_cxXsKccjList.html?doType=query"
}

func ZfEmptyClassRoom() string {
	return apis.ZF_URL + "cdjy/cdjy_cxKxcdlb.html?doType=query"
}
func ZfUserInfo() string {
	return apis.ZF_URL + "/xsxxxggl/xsgrxxwh_cxXsgrxx.html?gnmkdm=N100801&layout=default"
}
func ZfPY() string {
	return apis.ZF_URL + "/pyfagl/pyfaxxck_dyPyfaxx.html?id="
}
