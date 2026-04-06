package zfClient

type captchaLoginResponse struct {
	Msg    string `json:"msg"`
	T      int64  `json:"t"`
	Si     string `json:"si"`
	Imtk   string `json:"imtk"`
	Mi     string `json:"mi"`
	Vs     string `json:"vs"`
	Status string `json:"status"`
}

type captchaVerifyResponse struct {
	Msg    string `json:"msg"`
	Vs     string `json:"vs"`
	Status string `json:"status"`
}

type loginPublicKeyResponse struct {
	Modulus  string `json:"modulus"`
	Exponent string `json:"exponent"`
}

type SchoolTermInfo struct {
	Year string `json:"year"` // 学年
	Term string `json:"term"` // 学期 3->上学期, 12->下学期, 16->短学期
}
