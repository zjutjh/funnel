package helps

import (
	"encoding/base64"
	"encoding/json"
	"funnel/app/apis"
	"io/ioutil"
	"net/http"
	"net/url"
)

type captchaBreakerServerResponse struct {
	Code string `json:"code"`
	Data string `json:"data"`
}

func BreakCaptcha(s []byte) (string, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	captchaImageBase64 := base64.StdEncoding.EncodeToString(s)
	response, err := client.PostForm(apis.CAPTCHA_BREAKER_URL,
		url.Values{"img_base64": {"data:image/jpeg;base64," + captchaImageBase64}})

	if err != nil {
		return "", err
	}

	captcha, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	captchaRes := &captchaBreakerServerResponse{}
	err = json.Unmarshal(captcha, captchaRes)
	return captchaRes.Data, err
}
