package utils

import (
	"encoding/base64"
	"encoding/json"
	"funnel/app/apis"
	"funnel/app/utils/fetch"
	"net/url"
)

type captchaBreakerServerResponse struct {
	Code string `json:"code"`
	Data string `json:"data"`
}

func BreakCaptcha(s []byte) (string, error) {

	f := fetch.Fetch{}
	f.Init()
	captchaImageBase64 := base64.StdEncoding.EncodeToString(s)
	captcha, err := f.PostForm(apis.CAPTCHA_BREAKER_URL, url.Values{"img_base64": {"data:image/jpeg;base64," + captchaImageBase64}})

	if err != nil {
		return "", err
	}

	captchaRes := &captchaBreakerServerResponse{}
	err = json.Unmarshal(captcha, captchaRes)
	return captchaRes.Data, err
}
