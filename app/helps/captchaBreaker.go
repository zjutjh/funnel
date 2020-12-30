package helps

import (
	"encoding/base64"
	"encoding/json"
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
	url3 := "http://172.16.32.50/yzm"
	response, err := client.PostForm(url3,
		url.Values{"img_base64": {"data:image/jpeg;base64," + captchaImageBase64}})

	if err != nil{
		return "", err
	}

	captcha, err := ioutil.ReadAll(response.Body)

	if err != nil{
		return "", err
	}
	
	captchaRes := &captchaBreakerServerResponse{}
	err = json.Unmarshal(captcha, captchaRes)
	return captchaRes.Data, err
}
