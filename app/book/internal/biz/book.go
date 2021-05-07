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
	UpdateBook(ctx context.Context, book *Book) (*Book, error)
	DeleteBook(ctx context.Context, ids []string) error
	GetBook(ctx context.Context, id string) (*Book, error)
	ListBook(ctx context.Context, pageNum, pageSize int64) ([]*Book, error)
}

type BookUseCase struct {
	repo BookRepo
	log  *log.Helper
}

func (uc *BookUseCase) Create(ctx context.Context, book *Book) (*Book, error) {
	return uc.repo.CreateBook(ctx, book)
}

func (uc *BookUseCase) Update(ctx context.Context, book *Book) (*Book, error) {
	return uc.repo.UpdateBook(ctx, book)
}

func (uc *BookUseCase) Delete(ctx context.Context, ids []string) error {
	return uc.repo.DeleteBook(ctx, ids)
}

func (uc *BookUseCase) Get(ctx context.Context, id string) (*Book, error) {
	return uc.repo.GetBook(ctx, id)
}

func (uc *BookUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*Book, error) {
	return uc.repo.ListBook(ctx, pageNum, pageSize)
}
