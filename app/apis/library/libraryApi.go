package library

// OAuth
const (
	OAuthChaoXingInit = "https://tyrzfw.chaoxing.com/auth3/zjut/cas/init?isLogout=0&refer=https://opac.lib.zjut.edu.cn:8013/find/sso/login/zjut/0"
	OAuthBaseUrl      = "https://oauth.zjut.edu.cn/cas"
	OAuthPublicKey    = OAuthBaseUrl + "/v2/getPubKey"
	OAuthLogin        = OAuthBaseUrl + "/login"
)

// Library
const (
	LoginFromOAuth = OAuthLogin + "?service=http://tyrzfw.chaoxing.com/auth3/zjut/cas/index"

	BaseUrl           = "https://opac.lib.zjut.edu.cn:8013"
	CurrentLoanList   = BaseUrl + "/find/loanInfo/loanList"
	BorrowHistoryList = BaseUrl + "/find/loanInfo/loanHistoryList"
	UserInfo          = BaseUrl + "/oga/userinfo"
)
