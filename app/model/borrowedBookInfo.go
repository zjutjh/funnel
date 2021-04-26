package model

type BorrowedBookInfo struct {
	Name          string `json:"name"`
	LibraryID     string `json:"libraryID"`
	LibraryPlace  string `json:"libraryPlace"`
	Time          string `json:"time"`
	ReturnTime    string `json:"returnTime"`
	RenewalsTimes string `json:"renewalsTimes"`
	OverdueTime   string `json:"overdueTime"`
}
