package handler_test

import (
	"net/http"
	"net/http/httptest"
	"server/api/client/i18n"
	"server/api/handler"
	mock_i18n "server/mock/client/i18n"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

type bindTest struct {
	Password string `validate:"required"`
}

func TestNewErrorResponseDefault(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Vars
	errName := "err_test"
	jaErrName := "テストエラー"
	err := xerrors.New(errName)

	// SetUp
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// MockClient
	i18nm := mock_i18n.NewMockI18nClientInterface(ctrl)
	i18nm.EXPECT().T(errName).Return(jaErrName)

	// TestNewErrorResponse
	h := handler.NewHandler(nil, i18nm)
	_ = h.NewErrorResponse(c, err)

	if rec.Body == nil {
		t.Error("エラーメッセージの生成ができていません")
	}
	if !strings.Contains(rec.Body.String(), jaErrName) {
		t.Error("期待するエラーメッセージ存在しません")
	}
}

func TestNewErrorResponseWithValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Vars
	errName := echo.ErrValidatorNotRegistered.Error()
	jaErrName := i18n.LocaleMessages[0].Message

	// SetUp
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Bind
	bindReq := bindTest{}
	err := c.Validate(&bindReq)

	// MockClient
	i18nm := mock_i18n.NewMockI18nClientInterface(ctrl)
	i18nm.EXPECT().T(errName).Return(jaErrName)

	// TestNewErrorResponse
	h := handler.NewHandler(nil, i18nm)
	_ = h.NewErrorResponse(c, err)

	if rec.Body == nil {
		t.Error("エラーメッセージの生成ができていません")
	}
	if !strings.Contains(rec.Body.String(), jaErrName) {
		t.Error("期待するエラーメッセージ存在しません")
	}
}
