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
	"golang.org/x/xerrors"
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

func TestCreateBookSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var err error
	now := time.Now()
	uid := "test-uuid"
	name := "test-name"
	appReq := &application.CreateBookRequest{
		Name: name,
	}
	req := &model.Book{
		Name: name,
		UUID: uid,
	}
	expected := &model.Book{
		ID:        1,
		Name:      name,
		UUID:      uid,
		CreatedAt: now,
		UpdatedAt: now,
	}
	ctx := context.TODO()

	r := mock_repository.NewMockRepositoryInterface(ctrl)
	bRepo := mock_repository.NewMockBookRepositoryInterface(ctrl)
	bRepo.EXPECT().Create(ctx, r, req).Return(expected, err)

	// add fakeuuid
	model.SetFakeUUID(uid)

	a := application.NewApplication(
		&application.ApplicationBundle{
			Repository:     r,
			BookRepository: bRepo,
		},
	)

	res, err := a.CreateBook(ctx, appReq)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, expected) {
		t.Error("想定するbookと一致していません")
	}
}

func TestCreateBookError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()
	name := "test"
	uid := "test-uuid"
	err := xerrors.New("create error")

	model.SetFakeUUID(uid)

	appReq := &application.CreateBookRequest{
		Name: name,
	}
	modelReq := &model.Book{
		Name: name,
		UUID: uid,
	}

	r := mock_repository.NewMockRepositoryInterface(ctrl)
	bRepo := mock_repository.NewMockBookRepositoryInterface(ctrl)
	bRepo.EXPECT().Create(ctx, r, modelReq).Return(nil, err)

	a := application.NewApplication(
		&application.ApplicationBundle{
			Repository:     r,
			BookRepository: bRepo,
		},
	)

	_, e := a.CreateBook(ctx, appReq)
	if e.Error() != err.Error() {
		t.Error(e)
	}
}
