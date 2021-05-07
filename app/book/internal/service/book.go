package service

import (
	"context"
	"four/api/book/v1"
	"four/app/book/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

func NewBookService(buc *biz.BookUseCase, logger log.Logger) *BookService {
	return &BookService{
		buc: buc,
		log: log.NewHelper("service/book", logger),
	}
}

type BookService struct {
	v1.UnimplementedBookServer

	buc *biz.BookUseCase
	log *log.Helper
}

func (s *BookService) CreateBook(ctx context.Context, req *v1.CreateBookRequest) (*v1.CreateBookReply, error) {
	book, err := s.buc.Create(ctx, &biz.Book{
		ISBN:   req.ISBN,
		Name:   req.Name,
		Author: req.Author,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreateBookReply{Book: Book2DTO(book)}, nil
}

func (s *BookService) DeleteBook(ctx context.Context, req *v1.DeleteBookRequest) (*v1.DeleteBookReply, error) {
	err := s.buc.Delete(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteBookReply{}, nil
}

func (s *BookService) GetBook(ctx context.Context, req *v1.GetBookRequest) (*v1.GetBookReply, error) {
	book, err := s.buc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.GetBookReply{Book: Book2DTO(book)}, nil
}

func (s *BookService) ListBook(ctx context.Context, req *v1.ListBookRequest) (*v1.ListBookReply, error) {
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	books, err := s.buc.List(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}

	repBooks := make([]*v1.BookInfo, 0, len(books))
	for _, book := range books {
		repBooks = append(repBooks, Book2DTO(book))
	}

	return &v1.ListBookReply{
		Books: repBooks,
	}, nil
}

func (s *BookService) UpdateBook(ctx context.Context, req *v1.UpdateBookRequest) (*v1.UpdateBookReply, error) {
	_, err := s.buc.Update(ctx, Book2DO(req.Book))
	if err != nil {
		return nil, err
	}
	return &v1.UpdateBookReply{}, nil
}
