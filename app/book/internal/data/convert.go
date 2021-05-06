package data

import (
	"four/app/book/internal/biz"
	"four/app/book/internal/data/ent"
)

func Book2DO(book *ent.Book) *biz.Book {
	return &biz.Book{
		ID:     book.ID,
		ISBN:   book.ISBN,
		Name:   book.Name,
		Author: book.Author,
	}
}
