package application

import (
	"context"
	"server/api/domain/model"
)

func (a *application) GetBook(ctx context.Context, uid string) (*model.Book, error) {
	return a.BookRepository.FindByUUID(ctx, a.Repository, uid)
}

func (a *application) GetBooks(ctx context.Context) ([]*model.Book, error) {
	return a.BookRepository.FindAll(ctx, a.Repository)
}
