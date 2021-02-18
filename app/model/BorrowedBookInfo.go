package model

type BorrowedBookInfo struct {
	Name          string
	LibraryID     string
	LibraryPlace  string
	Time          string
	ReturnTime    string
	RenewalsTimes string
	IsExtended    string
}
