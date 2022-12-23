package application_test

import (
	"context"
	"reflect"
	"server/api/application"
	"server/api/domain/model"
	mock_repository "server/mock/infrastracture/persistence"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestGetBookSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var err error
	uid := "test-uuid"
	now := time.Now()
	expected := &model.Book{
		ID:        1,
		Name:      "test",
		UUID:      uid,
		CreatedAt: now,
		UpdatedAt: now,
	}
	ctx := context.TODO()
	r := mock_repository.NewMockRepositoryInterface(ctrl)
	bRepo := mock_repository.NewMockBookRepositoryInterface(ctrl)
	bRepo.EXPECT().FindByUUID(ctx, r, uid).Return(expected, err)

	a := application.NewApplication(
		&application.ApplicationBundle{
			Repository:     r,
			BookRepository: bRepo,
		},
	)

	res, err := a.GetBook(ctx, uid)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, expected) {
		t.Error("想定するbookと一致していません")
	}
}

func TestGetBookError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	uid := "test-uuid"

	r := mock_repository.NewMockRepositoryInterface(ctrl)
	bRepo := mock_repository.NewMockBookRepositoryInterface(ctrl)
	bRepo.EXPECT().FindByUUID(ctx, r, uid).Return(nil, nil)

	a := application.NewApplication(
		&application.ApplicationBundle{
			Repository:     r,
			BookRepository: bRepo,
		},
	)

	res, err := a.GetBook(ctx, uid)
	if err != nil {
		t.Error(err)
	}
	if res != nil {
		t.Error("bookが取得できています")
	}
}

func TestGetBooksSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var err error
	now := time.Now()
	expected := []*model.Book{
		{
			ID:        1,
			Name:      "test",
			UUID:      "test-uuid",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	ctx := context.TODO()

	r := mock_repository.NewMockRepositoryInterface(ctrl)
	bRepo := mock_repository.NewMockBookRepositoryInterface(ctrl)
	bRepo.EXPECT().FindAll(ctx, r).Return(expected, err)

	a := application.NewApplication(
		&application.ApplicationBundle{
			Repository:     r,
			BookRepository: bRepo,
		},
	)

	res, err := a.GetBooks(ctx)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, expected) {
		t.Error("想定するbookと一致していません")
	}
}

func TestGetBooksError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	r := mock_repository.NewMockRepositoryInterface(ctrl)
	bRepo := mock_repository.NewMockBookRepositoryInterface(ctrl)
	bRepo.EXPECT().FindAll(ctx, r).Return(nil, nil)

	a := application.NewApplication(
		&application.ApplicationBundle{
			Repository:     r,
			BookRepository: bRepo,
		},
	)

	res, err := a.GetBooks(ctx)
	if err != nil {
		t.Error(err)
	}
	if res != nil {
		t.Error("bookが取得できています")
	}
}
