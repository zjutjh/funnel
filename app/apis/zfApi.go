package apis

import (
	"strconv"
	"time"
)

var ZfLoginGetPublickey = ZF_URL + "xtgl/login_getPublicKey.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
var ZfLoginHome = ZF_URL + "xtgl/login_slogin.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
var ZfLoginKaptcha = ZF_URL + "kaptcha?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
var ZfExamInfo = ZF_URL + "/kwgl/kscx_cxXsksxxIndex.html?doType=query"
var ZfClassTable = ZF_URL + "kbcx/xskbcx_cxXsKb.html?doType=query"
var ZfScore = ZF_URL + "cjcx/cjcx_cxDgXscj.html?doType=query"
var ZfScoreDetail = ZF_URL + "cjcx/cjcx_cxXsKccjList.html?doType=query"

var ZfEmptyClassRoom = ZF_URL + "cdjy/cdjy_cxKxcdlb.html?doType=query"
var ZfUserInfo = ZF_URL + "/xsxxxggl/xsgrxxwh_cxXsgrxx.html?gnmkdm=N100801&layout=default"
var ZfPY = ZF_URL + "/pyfagl/pyfaxxck_dyPyfaxx.html?id="
