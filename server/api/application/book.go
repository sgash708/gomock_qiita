package application

import (
	"context"
	"server/api/domain/model"

	"golang.org/x/xerrors"
)

type CreateBookRequest struct {
	Name string `json:"name"`
}

type UpdateBookRequest struct {
	UUID string `json:"uuid"`
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

func (a *application) UpdateBook(ctx context.Context, req *UpdateBookRequest) (*model.Book, error) {
	book, err := a.BookRepository.FindByUUID(ctx, a.Repository, req.UUID)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, xerrors.New(model.NotFoundUUIDMsg)
	}

	book.UpdateBookName(req.Name)

	return a.BookRepository.Update(ctx, a.Repository, book)
}
