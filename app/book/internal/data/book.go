package data

import (
	"context"
	"four/app/book/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/satori/go.uuid"
)

func NewBookRepo(data *Data, logger log.Logger) biz.BookRepo {
	return &bookRepo{
		data: data,
		log:  log.NewHelper("data/bookRepo", logger),
	}
}

type bookRepo struct {
	data *Data
	log  *log.Helper
}

func (r *bookRepo) CreateBook(ctx context.Context, book *biz.Book) (*biz.Book, error) {
	po, err := r.data.db.Book.Create().
		SetID(uuid.NewV4().String()).
		SetISBN(book.ISBN).
		SetName(book.Name).
		SetAuthor(book.Author).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return Book2DO(po), nil
}

func (r *bookRepo) UpdateBook(ctx context.Context, book *biz.Book) error {
	panic("implement me")
}

func (r *bookRepo) DeleteBook(ctx context.Context, ids []string) error {
	panic("implement me")
}

func (r *bookRepo) GetBook(ctx context.Context, id string) (*biz.Book, error) {
	po, err := r.data.db.Book.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return Book2DO(po), nil
}

func (r *bookRepo) ListBook(ctx context.Context, ids []string) ([]*biz.Book, error) {
	panic("implement me")
}
