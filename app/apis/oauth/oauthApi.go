package oauth

func OauthLoginGetPublickey() string {
	return "https://oauth.zjut.edu.cn/cas/v2/getPubKey"
}

func OauthLoginHome() string {
	return "https://oauth.zjut.edu.cn/cas/login?service=http%3A%2F%2Fwww.gdjw.zjut.edu.cn%2Fsso%2Fzfiotlogin"
}
