package data

import (
	"context"
	"database/sql"
	"four/app/book/internal/biz"
	_err "four/app/book/internal/pkg/errors"
	kraerr "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

func NewBookRepo2(logger log.Logger) biz.BookRepo {
	return &bookRepoM{
		log: log.NewHelper("data/bookRepoM", logger),
	}
}

type bookRepoM struct {
	log *log.Helper
}

func (r *bookRepoM) CreateBook(ctx context.Context, book *biz.Book) (*biz.Book, error) {
	err := errors.New("CreateBook")
	return nil, errors.Wrapf(kraerr.Internal(_err.GetDetail(_err.CreateBookFail)), "create book failed: %v", err)
}

func (r *bookRepoM) UpdateBook(ctx context.Context, book *biz.Book) (*biz.Book, error) {
	err := errors.New("UpdateBook")
	return nil, errors.Wrapf(kraerr.Internal(_err.GetDetail(_err.UpdateBookFail)), "update book failed: %v", err)
}

func (r *bookRepoM) DeleteBook(ctx context.Context, ids []string) error {
	err := errors.New("DeleteBook")
	return errors.Wrapf(kraerr.Internal(_err.GetDetail(_err.DeleteBookFail)), "update book failed: %v", err)
}

func (r *bookRepoM) GetBook(ctx context.Context, id string) (*biz.Book, error) {
	err := errors.New("GetBook")
	if err == sql.ErrNoRows {
		return nil, errors.Wrapf(kraerr.NotFound(_err.GetDetail(_err.BookNotFound)), "get book failed: %v", err)
	} else if err != nil {
		return nil, errors.Wrapf(kraerr.Unknown(_err.GetDetail(_err.GetBookFail)), "get book failed: %v", err)
	}
	return &biz.Book{}, nil
}

func (r *bookRepoM) ListBook(ctx context.Context, pageNum, pageSize int64) ([]*biz.Book, error) {
	err := errors.New("ListBook")
	if err == sql.ErrNoRows {
		return nil, errors.Wrapf(kraerr.NotFound(_err.GetDetail(_err.BookNotFound)), "get book list failed: %v", err)
	} else if err != nil {
		return nil, errors.Wrapf(kraerr.Unknown(_err.GetDetail(_err.GetBookFail)), "get book list failed: %v", err)
	}
	return []*biz.Book{}, nil
}
