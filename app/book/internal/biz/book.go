package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

func NewBookUseCase(repo BookRepo, logger log.Logger) *BookUseCase {
	return &BookUseCase{
		repo: repo,
		log:  log.NewHelper("useCase/book", logger),
	}
}

type Book struct {
	ID     string
	ISBN   string
	Name   string
	Author string
}

type BookRepo interface {
	CreateBook(ctx context.Context, book *Book) (*Book, error)
	UpdateBook(ctx context.Context, book *Book) error
	DeleteBook(ctx context.Context, ids []string) error
	GetBook(ctx context.Context, id string) (*Book, error)
	ListBook(ctx context.Context, ids []string) ([]*Book, error)
}

type BookUseCase struct {
	repo BookRepo
	log  *log.Helper
}
