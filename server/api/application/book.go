package application

import (
	"context"
	"server/api/domain/model"
)

type CreateBookRequest struct {
	Name string `json:"name"`
}

func (a *application) GetBook(ctx context.Context, uid string) (*model.Book, error) {
	return a.BookRepository.FindByUUID(ctx, a.Repository, uid)
}

func (a *application) GetBooks(ctx context.Context) ([]*model.Book, error) {
	return a.BookRepository.FindAll(ctx, a.Repository)
}

func (a *application) CreateBook(ctx context.Context, req *CreateBookRequest) (*model.Book, error) {
	book := &model.Book{
		Name: req.Name,
	}

	if err := book.SetUUID(); err != nil {
		return nil, err
	}

	return a.BookRepository.Create(ctx, a.Repository, book)
}
