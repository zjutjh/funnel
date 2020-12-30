package apis

import (
	"strconv"
	"time"
)

var ZfLoginGetPublickey = "xtgl/login_getPublicKey.html?time=" + strconv.FormatInt(time.Now().Unix()*1000,10)

