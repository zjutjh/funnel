package apis

import (
	"funnel/config/config"
)

var ZF_URL string = config.Config.GetString("zf.url")
var ZF_Main_URL string = config.Config.GetString("zf.main_url")
var ZF_BK_URL string = config.Config.GetString("zf.bk_url")
