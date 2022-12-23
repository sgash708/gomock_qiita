package application

import (
	"context"
	"server/api/domain/model"
	"server/config"
)

type application struct {
	*ApplicationBundle
}

type ApplicationBundle struct {
	ServerConfig *config.ServerConfig
	// Repository persistence.RepositoryInterface
}

type ApplicationInterface interface {
	// Book
	GetBook(ctx context.Context, uid string) (*model.Book, error)
}

func NewApplication(bdl *ApplicationBundle) ApplicationInterface {
	return &application{bdl}
}
