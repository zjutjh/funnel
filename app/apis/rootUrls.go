package apis

import (
	"funnel/config/config"
)

//var LIBRARY_URL = os.Getenv("LIBRARY_URL")
//var CAPTCHA_BREAKER_URL = os.Getenv("CAPTCHA_BREAKER_BK_URL")
//var CAPTCHA_NEW_BREAKER_URL = os.Getenv("CAPTCHA_BREAKER_NEW_URL")
//var ZF_URL = os.Getenv("ZF_URL")
//
//var ZF_Main_URL = os.Getenv("ZF_URL")
//var ZF_BK_URL = os.Getenv("ZF_URL_BK")
//
//var CANTEEN_URL = os.Getenv("CANTEEN_URL")

//var LIBRARY_URL string = config.Config.GetString("library")

var CAPTCHA_BREAKER_URL string = config.Config.GetString("captcha.breaker_url")
var CAPTCHA_NEW_BREAKER_URL string = config.Config.GetString("captcha.new_breaker_url")

var ZF_URL string = config.Config.GetString("zf.url")
var ZF_Main_URL string = config.Config.GetString("zf.main_url")
var ZF_BK_URL string = config.Config.GetString("zf.bk_url")

var CANTEEN_URL string = config.Config.GetString("canteen")
