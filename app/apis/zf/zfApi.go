package zf

import (
	"funnel/config/config"
	"net/url"
	"strconv"
	"time"
)

var UrlMap = config.Config.GetStringMapString("zf.url")

const defaultURL = "http://www.gdjw.zjut.edu.cn/jwglxt/"

const (
	ZFURLDefaultFlag = "default"
	ZFURLMainFlag    = "main"
	ZFURLJfFlag      = "jf"
)

// UrlToFLag 用于获取url对应的flag, 默认返回main
func UrlToFLag(rawUrl string) string {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return ZFURLMainFlag
	}
	for k, v := range UrlMap {
		if k == "default" {
			continue
		}
		u2, err := url.Parse(v)
		if err != nil {
			continue
		}
		if u.Host == u2.Host {
			return k
		}
	}
	return ZFURLMainFlag
}

// ChooseURL 根据flag 选择正方 host
func ChooseURL(flag string) string {
	if u, ok := UrlMap[flag]; ok {
		return u
	} else {
		return defaultURL
	}
}

func ZfLoginGetPublickey(flag string) string {
	return ChooseURL(flag) + "xtgl/login_getPublicKey.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
}
func ZfLoginHome(flag string) string {
	return ChooseURL(flag) + "xtgl/login_slogin.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
}
func ZfExamInfo(flag string) string {
	return ChooseURL(flag) + "kwgl/kscx_cxXsksxxIndex.html?doType=query&gnmkdm=N358105"
}
func ZfClassTable(flag string) string {
	return ChooseURL(flag) + "kbcx/xskbcx_cxXsgrkb.html?gnmkdm=N2151&su="
}
func ZfScore(flag string) string {
	return ChooseURL(flag) + "cjcx/cjcx_cxDgXscj.html?doType=query&gnmkdm=N305005"
}
func ZfMinTermScore(flag string) string {
	return ChooseURL(flag) + "design/funcData_cxFuncDataList.html?func_widget_guid=5EF567BFD3CE243EE053A11310AC1252&gnmkdm=N305013"
}

func ZfCaptchaURL(flag string) string {
	return ChooseURL(flag) + "zfcaptchaLogin?instanceId=zfcaptchaLogin"
}

func ZfEmptyClassRoom(flag string) string {
	return ChooseURL(flag) + "cdjy/cdjy_cxKxcdlb.html?doType=query"
}
