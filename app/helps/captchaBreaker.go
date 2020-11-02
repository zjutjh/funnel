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

func BreakCaptcha(s []byte) string {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	ymzImageBase64 := base64.StdEncoding.EncodeToString(s)
	url3 := "http://172.16.32.50/yzm"
	response, _ := client.PostForm(url3, url.Values{"img_base64": {"data:image/jpeg;base64," + ymzImageBase64}})
	yzm, _ := ioutil.ReadAll(response.Body)
	yzmRes := &captchaBreakerServerResponse{}
	_ = json.Unmarshal(yzm, yzmRes)
	return yzmRes.Data
}
