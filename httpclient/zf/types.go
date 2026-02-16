package zfClient

type CaptchaLoginResponse struct {
	Msg    string `json:"msg"`
	T      int64  `json:"t"`
	Si     string `json:"si"`
	Imtk   string `json:"imtk"`
	Mi     string `json:"mi"`
	Vs     string `json:"vs"`
	Status string `json:"status"`
}

type CaptchaVerifyResponse struct {
	Msg    string `json:"msg"`
	Vs     string `json:"vs"`
	Status string `json:"status"`
}
