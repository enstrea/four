package errors

var errText map[string]string

func Tip(reason string) string {
	if tip, ok := errText[reason]; ok {
		return tip
	}
	return "未知错误"
}

func init() {
	errText = make(map[string]string)

	errText[CreateBookFail] = "添加图书信息失败"
	errText[UpdateBookFail] = "更新图书信息失败"
	errText[DeleteBookFail] = "更新图书信息失败"
	errText[GetBookFail] = "获取图书信息失败"
	errText[BookNotFound] = "图书信息未找到"

}
