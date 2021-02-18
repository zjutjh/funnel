package zf

import (
	"funnel/app/apis"
	"strconv"
	"time"
)

var ZfLoginGetPublickey = apis.ZF_URL + "xtgl/login_getPublicKey.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
var ZfLoginHome = apis.ZF_URL + "xtgl/login_slogin.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
var ZfLoginKaptcha = apis.ZF_URL + "kaptcha?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
var ZfExamInfo = apis.ZF_URL + "/kwgl/kscx_cxXsksxxIndex.html?doType=query"
var ZfClassTable = apis.ZF_URL + "kbcx/xskbcx_cxXsKb.html?doType=query"
var ZfScore = apis.ZF_URL + "cjcx/cjcx_cxDgXscj.html?doType=query"
var ZfScoreDetail = apis.ZF_URL + "cjcx/cjcx_cxXsKccjList.html?doType=query"

var ZfEmptyClassRoom = apis.ZF_URL + "cdjy/cdjy_cxKxcdlb.html?doType=query"
var ZfUserInfo = apis.ZF_URL + "/xsxxxggl/xsgrxxwh_cxXsgrxx.html?gnmkdm=N100801&layout=default"
var ZfPY = apis.ZF_URL + "/pyfagl/pyfaxxck_dyPyfaxx.html?id="

func ZFServerChange() {
	if apis.ZF_URL == apis.ZF_Main_URL {
		apis.ZF_URL = apis.ZF_BK_URL
	} else {
		apis.ZF_URL = apis.ZF_Main_URL
	}
}
