package application

import (
	"context"
	"server/api/domain/model"
	"server/api/infrastracture/persistence"
	"server/config"
)

type application struct {
	*ApplicationBundle
}

type ApplicationBundle struct {
	ServerConfig   *config.ServerConfig
	Repository     persistence.RepositoryInterface
	BookRepository persistence.BookRepositoryInterface
}

type ApplicationInterface interface {
	// Book
	GetBook(ctx context.Context, uid string) (*model.Book, error)
	GetBooks(ctx context.Context) ([]*model.Book, error)
}

func NewApplication(bdl *ApplicationBundle) ApplicationInterface {
	return &application{bdl}
}
