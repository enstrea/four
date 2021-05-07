package errors

import "net/http"

var (
	NotFound = New(http.StatusNotFound, "")
)
