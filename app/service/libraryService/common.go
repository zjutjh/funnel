package libraryService

type BookInfo struct {
	LoanID         uint   `json:"loanId"`         // 借书记录的id
	Barcode        string `json:"barcode"`        //书的条形码
	LocationName   string `json:"locationName"`   //图书馆位置
	Title          string `json:"title"`          //书名
	LoanDate       string `json:"loanDate"`       //借的时间
	NormReturnDate string `json:"normReturnDate"` //还书的期限
	ReturnDate     string `json:"returnDate"`     //还书的时间
}

// Result 交由Resty实现response的自动反序列化
type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	ErrCode int    `json:"errCode"`
	Data    struct {
		SearchResult []BookInfo `json:"searchResult"`
		NumFound     int        `json:"numFound"`
	} `json:"data"`
}
