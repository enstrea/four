package service

import (
	"context"
	"four/api/book/v1"
)

type BookService struct {
	v1.UnimplementedBookServer

}

func (s *BookService) CreateBook(ctx context.Context, request *v1.CreateBookRequest) (*v1.CreateBookReply, error) {
	panic("implement me")
}

func (s *BookService) DeleteBook(ctx context.Context, request *v1.DeleteBookRequest) (*v1.DeleteBookReply, error) {
	panic("implement me")
}

func (s *BookService) GetBook(ctx context.Context, request *v1.GetBookRequest) (*v1.GetBookReply, error) {
	panic("implement me")
}

func (s *BookService) ListBook(ctx context.Context, request *v1.ListBookRequest) (*v1.ListBookReply, error) {
	panic("implement me")
}

func (s *BookService) UpdateBook(ctx context.Context, request *v1.UpdateBookRequest) (*v1.UpdateBookReply, error) {
	panic("implement me")
}



