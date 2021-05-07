package data

import (
	"context"
	"four/app/book/internal/biz"
	"four/app/book/internal/data/ent/book"
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

func (r *bookRepo) UpdateBook(ctx context.Context, book *biz.Book) (*biz.Book, error) {
	po, err := r.data.db.Book.UpdateOneID(book.ID).
		SetName(book.Name).
		SetAuthor(book.Author).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return Book2DO(po), nil
}

func (r *bookRepo) DeleteBook(ctx context.Context, ids []string) error {
	_, err := r.data.db.Book.Delete().Where(book.IDIn(ids...)).Exec(ctx)
	return err
}

func (r *bookRepo) GetBook(ctx context.Context, id string) (*biz.Book, error) {
	po, err := r.data.db.Book.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return Book2DO(po), nil
}

func (r *bookRepo) ListBook(ctx context.Context, pageNum, pageSize int64) ([]*biz.Book, error) {
	pos, err := r.data.db.Book.Query().
		Offset(int((pageNum - 1) * pageSize)).
		Limit(int(pageSize)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	dos := make([]*biz.Book, 0, len(pos))
	for _, po := range pos {
		dos = append(dos, Book2DO(po))
	}
	return dos, nil
}
