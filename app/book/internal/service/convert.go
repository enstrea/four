package service

import (
	v1 "four/api/book/v1"
	"four/app/book/internal/biz"
)

func Book2DO(info *v1.BookInfo) *biz.Book {
	return &biz.Book{
		ID:     info.Id,
		ISBN:   info.ISBN,
		Name:   info.Name,
		Author: info.Author,
	}
}

func Book2DTO(book *biz.Book) *v1.BookInfo {
	return &v1.BookInfo{
		Id:     book.ID,
		ISBN:   book.ISBN,
		Name:   book.Name,
		Author: book.Author,
	}
}
