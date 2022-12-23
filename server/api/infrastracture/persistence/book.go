package persistence

import (
	"context"
	"errors"
	"server/api/domain/model"

	"gorm.io/gorm"
)

type BookRepositoryInterface interface {
	FindAll(ctx context.Context, ri RepositoryInterface) ([]*model.Book, error)
	FindByUUID(context.Context, RepositoryInterface, string) (*model.Book, error)
	Create(context.Context, RepositoryInterface, *model.Book) (*model.Book, error)
	Update(context.Context, RepositoryInterface, *model.Book) (*model.Book, error)
	Delete(context.Context, RepositoryInterface, *model.Book) error
}

type bookRepository struct{}

func NewBookRepository() BookRepositoryInterface {
	return &bookRepository{}
}

func (br *bookRepository) FindAll(ctx context.Context, ri RepositoryInterface) ([]*model.Book, error) {
	db, err := getDBConnection(ri)
	if err != nil {
		return nil, err
	}

	var book []*model.Book
	if err := db.Find(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return book, nil
}

func (br *bookRepository) FindByUUID(ctx context.Context, ri RepositoryInterface, uid string) (*model.Book, error) {
	db, err := getDBConnection(ri)
	if err != nil {
		return nil, err
	}

	var book model.Book
	if err := db.First(&book, "uuid = ?", uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &book, nil
}

func (br *bookRepository) Create(ctx context.Context, ri RepositoryInterface, book *model.Book) (*model.Book, error) {
	db, err := getDBConnection(ri)
	if err != nil {
		return nil, err
	}

	if err := db.Create(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (br *bookRepository) Update(ctx context.Context, ri RepositoryInterface, book *model.Book) (*model.Book, error) {
	db, err := getDBConnection(ri)
	if err != nil {
		return nil, err
	}

	if err := db.Where(book.ID).Updates(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (br *bookRepository) Delete(ctx context.Context, ri RepositoryInterface, book *model.Book) error {
	db, err := getDBConnection(ri)
	if err != nil {
		return err
	}

	return db.Delete(book).Error
}
