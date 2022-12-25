package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"server/api/application"
	"server/api/domain/model"
	"server/api/handler"
	mock_application "server/mock/application"
	mock_i18n "server/mock/client/i18n"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

func TestGetBookSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Vars
	id := 1
	name := "scenario_name"
	uuid := "test_handler_uuid"
	ctx := context.TODO()

	// SetUp
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(uuid)

	// IO
	appRes := &model.Book{
		ID:   id,
		Name: name,
		UUID: uuid,
	}

	// MockApplication
	app := mock_application.NewMockApplicationInterface(ctrl)
	app.EXPECT().GetBook(ctx, uuid).Return(appRes, nil)

	// TestGetBook
	ch := handler.NewHandler(app, nil)
	if err := ch.GetBook(c); err != nil {
		t.Error(err)
	}

	// Status
	expCode := http.StatusOK
	recCode := rec.Code
	recBody := rec.Body

	// Check
	if expCode != recCode {
		t.Errorf("expected: %v \n real: %v", expCode, recCode)
	}
	if recBody == nil {
		t.Errorf("bodyの取得に失敗しています")
	}
	if !strings.Contains(recBody.String(), name) {
		t.Error("期待するNameが存在しません")
	}
}

func TestGetBookError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Vars
	uuid := "test_uuid"
	errName := "err_test"
	jaErrName := "テストエラー"
	err := xerrors.New(errName)
	ctx := context.TODO()

	// SetUp
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(uuid)

	// MockApplication
	app := mock_application.NewMockApplicationInterface(ctrl)
	app.EXPECT().GetBook(ctx, uuid).Return(nil, err)

	// MockClient
	i18nm := mock_i18n.NewMockI18nClientInterface(ctrl)
	i18nm.EXPECT().T(errName).Return(jaErrName)

	// TestGetBook
	ch := handler.NewHandler(app, i18nm)
	if err := ch.GetBook(c); err != nil {
		t.Error(err)
	}

	// Status
	expCode := http.StatusBadRequest
	recCode := rec.Code
	recBody := rec.Body

	// Check
	if expCode != recCode {
		t.Errorf("expected: %v \n real: %v", expCode, recCode)
	}
	if recBody == nil {
		t.Errorf("bodyの取得に失敗しています")
	}
	if !strings.Contains(recBody.String(), jaErrName) {
		t.Error("期待するエラーが存在しません")
	}
}

func TestGetBooksSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Vars
	id := 1
	name := "scenario_name"
	uuid := "test_handler_uuid"
	ctx := context.TODO()

	// SetUp
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books")

	// IO
	appRes := []*model.Book{
		{
			ID:   id,
			Name: name,
			UUID: uuid,
		},
	}

	// MockApplication
	app := mock_application.NewMockApplicationInterface(ctrl)
	app.EXPECT().GetBooks(ctx).Return(appRes, nil)

	// TestGetBook
	ch := handler.NewHandler(app, nil)
	if err := ch.GetBooks(c); err != nil {
		t.Error(err)
	}

	// Status
	expCode := http.StatusOK
	recCode := rec.Code
	recBody := rec.Body

	// Check
	if expCode != recCode {
		t.Errorf("expected: %v \n real: %v", expCode, recCode)
	}
	if recBody == nil {
		t.Errorf("bodyの取得に失敗しています")
	}
	if !strings.Contains(recBody.String(), name) {
		t.Error("期待するNameが存在しません")
	}
}

func TestGetBooksError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Vars
	errName := "err_test"
	jaErrName := "テストエラー"
	err := xerrors.New(errName)
	ctx := context.TODO()

	// SetUp
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books")

	// MockApplication
	app := mock_application.NewMockApplicationInterface(ctrl)
	app.EXPECT().GetBooks(ctx).Return(nil, err)

	// MockClient
	i18nm := mock_i18n.NewMockI18nClientInterface(ctrl)
	i18nm.EXPECT().T(errName).Return(jaErrName)

	// TestGetBook
	ch := handler.NewHandler(app, i18nm)
	if err := ch.GetBooks(c); err != nil {
		t.Error(err)
	}

	// Status
	expCode := http.StatusBadRequest
	recCode := rec.Code
	recBody := rec.Body

	// Check
	if expCode != recCode {
		t.Errorf("expected: %v \n real: %v", expCode, recCode)
	}
	if recBody == nil {
		t.Errorf("bodyの取得に失敗しています")
	}
	if !strings.Contains(recBody.String(), jaErrName) {
		t.Error("期待するエラーが存在しません")
	}
}

func TestCreateBookSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Vars
	id := 1
	name := "scenario_name"
	uuid := "test_handler_uuid"
	ctx := context.TODO()
	reqStrJSON := fmt.Sprintf(
		`{"name": "%s" }`,
		name,
	)

	// SetUp
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqStrJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books")

	// IO
	appReq := &application.CreateBookRequest{
		Name: name,
	}
	appRes := &model.Book{
		ID:   id,
		Name: name,
		UUID: uuid,
	}

	// MockApplication
	app := mock_application.NewMockApplicationInterface(ctrl)
	app.EXPECT().CreateBook(ctx, appReq).Return(appRes, nil)

	// TestCreateBook
	ch := handler.NewHandler(app, nil)
	if err := ch.CreateBook(c); err != nil {
		t.Error(err)
	}

	// Status
	expCode := http.StatusCreated
	recCode := rec.Code
	recBody := rec.Body

	// Check
	if expCode != recCode {
		t.Errorf("expected: %v \n real: %v", expCode, recCode)
	}
	if recBody == nil {
		t.Errorf("bodyの取得に失敗しています")
	}
	if !strings.Contains(recBody.String(), name) {
		t.Error("期待するNameが存在しません")
	}
}

func TestCreateBookError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Vars
	name := "err_test"
	errName := "err_test"
	jaErrName := "テストエラー"
	err := xerrors.New(errName)
	ctx := context.TODO()
	reqStrJSON := fmt.Sprintf(
		`{"name": "%s" }`,
		name,
	)
	appReq := &application.CreateBookRequest{
		Name: name,
	}

	// SetUp
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqStrJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books")

	// MockApplication
	app := mock_application.NewMockApplicationInterface(ctrl)
	app.EXPECT().CreateBook(ctx, appReq).Return(nil, err)

	// MockClient
	i18nm := mock_i18n.NewMockI18nClientInterface(ctrl)
	i18nm.EXPECT().T(errName).Return(jaErrName)

	// TestCreateBook
	ch := handler.NewHandler(app, i18nm)
	if err := ch.CreateBook(c); err != nil {
		t.Error(err)
	}

	// Status
	expCode := http.StatusBadRequest
	recCode := rec.Code
	recBody := rec.Body

	// Check
	if expCode != recCode {
		t.Errorf("expected: %v \n real: %v", expCode, recCode)
	}
	if recBody == nil {
		t.Errorf("bodyの取得に失敗しています")
	}
	if !strings.Contains(recBody.String(), jaErrName) {
		t.Error("期待するエラーが存在しません")
	}
}

func TestUpdateBookSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Vars
	id := 1
	name := "scenario_name"
	uuid := "test_handler_uuid"
	ctx := context.TODO()
	reqStrJSON := fmt.Sprintf(
		`{"name": "%s" }`,
		name,
	)

	// SetUp
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(reqStrJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(uuid)

	// IO
	appReq := &application.UpdateBookRequest{
		UUID: uuid,
		Name: name,
	}
	appRes := &model.Book{
		ID:   id,
		UUID: uuid,
		Name: name,
	}

	// MockApplication
	app := mock_application.NewMockApplicationInterface(ctrl)
	app.EXPECT().UpdateBook(ctx, appReq).Return(appRes, nil)

	// TestUpdateBook
	ch := handler.NewHandler(app, nil)
	if err := ch.UpdateBook(c); err != nil {
		t.Error(err)
	}

	// Status
	expCode := http.StatusOK
	recCode := rec.Code
	recBody := rec.Body

	// Check
	if expCode != recCode {
		t.Errorf("expected: %v \n real: %v", expCode, recCode)
	}
	if recBody == nil {
		t.Errorf("bodyの取得に失敗しています")
	}
	if !strings.Contains(recBody.String(), name) {
		t.Error("期待するNameが存在しません")
	}
}

func TestUpdateBookError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Vars
	name := "err_test"
	uuid := "err_uuid"
	errName := "err_test"
	jaErrName := "テストエラー"
	err := xerrors.New(errName)
	ctx := context.TODO()
	reqStrJSON := fmt.Sprintf(
		`{"name": "%s" }`,
		name,
	)
	appReq := &application.UpdateBookRequest{
		UUID: uuid,
		Name: name,
	}

	// SetUp
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(reqStrJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(uuid)

	// MockApplication
	app := mock_application.NewMockApplicationInterface(ctrl)
	app.EXPECT().UpdateBook(ctx, appReq).Return(nil, err)

	// MockClient
	i18nm := mock_i18n.NewMockI18nClientInterface(ctrl)
	i18nm.EXPECT().T(errName).Return(jaErrName)

	// TestUpdateBook
	ch := handler.NewHandler(app, i18nm)
	if err := ch.UpdateBook(c); err != nil {
		t.Error(err)
	}

	// Status
	expCode := http.StatusBadRequest
	recCode := rec.Code
	recBody := rec.Body

	// Check
	if expCode != recCode {
		t.Errorf("expected: %v \n real: %v", expCode, recCode)
	}
	if recBody == nil {
		t.Errorf("bodyの取得に失敗しています")
	}
	if !strings.Contains(recBody.String(), jaErrName) {
		t.Error("期待するエラーが存在しません")
	}
}

func TestDeleteBookSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Vars
	uuid := "test_handler_uuid"
	ctx := context.TODO()

	// SetUp
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(uuid)

	// MockApplication
	app := mock_application.NewMockApplicationInterface(ctrl)
	app.EXPECT().DeleteBook(ctx, uuid).Return(nil)

	// TestDeleteBook
	ch := handler.NewHandler(app, nil)
	if err := ch.DeleteBook(c); err != nil {
		t.Error(err)
	}

	// Status
	expCode := http.StatusNoContent
	recCode := rec.Code

	// Check
	if expCode != recCode {
		t.Errorf("expected: %v \n real: %v", expCode, recCode)
	}
}

func TestDeleteBookError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Vars
	uuid := "err_uuid"
	errName := "err_test"
	jaErrName := "テストエラー"
	err := xerrors.New(errName)
	ctx := context.TODO()

	// SetUp
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(uuid)

	// MockApplication
	app := mock_application.NewMockApplicationInterface(ctrl)
	app.EXPECT().DeleteBook(ctx, uuid).Return(err)

	// MockClient
	i18nm := mock_i18n.NewMockI18nClientInterface(ctrl)
	i18nm.EXPECT().T(errName).Return(jaErrName)

	// TestDeleteBook
	ch := handler.NewHandler(app, i18nm)
	if err := ch.DeleteBook(c); err != nil {
		t.Error(err)
	}

	// Status
	expCode := http.StatusBadRequest
	recCode := rec.Code
	recBody := rec.Body

	// Check
	if expCode != recCode {
		t.Errorf("expected: %v \n real: %v", expCode, recCode)
	}
	if recBody == nil {
		t.Errorf("bodyの取得に失敗しています")
	}
	if !strings.Contains(recBody.String(), jaErrName) {
		t.Error("期待するエラーが存在しません")
	}
}
