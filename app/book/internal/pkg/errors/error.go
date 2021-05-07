package errors

func GetDetail(reason string) (string, string) {
	return reason, Tip(reason)
}

const (
	CreateBookFail = "CreateBookFail"
	UpdateBookFail = "UpdateBookFail"
	DeleteBookFail = "DeleteBookFail"
	GetBookFail    = "GetBookFail"
	BookNotFound   = "BookNotFound"
)
