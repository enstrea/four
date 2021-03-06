// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package errors

import (
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

const (
	Errors_Unknown        = "Book_Unknown"
	Errors_CreateBookFail = "Book_CreateBookFail"
	Errors_UpdateBookFail = "Book_UpdateBookFail"
	Errors_DeleteBookFail = "Book_DeleteBookFail"
	Errors_GetBookFail    = "Book_GetBookFail"
	Errors_BookNotFound   = "Book_BookNotFound"
)

func IsUnknown(err error) bool {
	return errors.Reason(err) == Errors_Unknown
}

func IsCreateBookFail(err error) bool {
	return errors.Reason(err) == Errors_CreateBookFail
}

func IsUpdateBookFail(err error) bool {
	return errors.Reason(err) == Errors_UpdateBookFail
}

func IsDeleteBookFail(err error) bool {
	return errors.Reason(err) == Errors_DeleteBookFail
}

func IsGetBookFail(err error) bool {
	return errors.Reason(err) == Errors_GetBookFail
}

func IsBookNotFound(err error) bool {
	return errors.Reason(err) == Errors_BookNotFound
}
