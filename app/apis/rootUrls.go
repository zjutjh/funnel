package apis

import "os"

var LIBRARY_URL = os.Getenv("LIBRARY_URL")
var CAPTCHA_BREAKER_URL = os.Getenv("CAPTCHA_BREAKER_URL")
var ZF_URL = os.Getenv("ZF_URL")
