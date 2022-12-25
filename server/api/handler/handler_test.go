package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"server/api/handler"
	mock_application "server/mock/application"
	mock_i18n "server/mock/client/i18n"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func inject() {
	e = echo.New()
}

func getPath() []string {
	return []string{
		"books",
	}
}

func contains(paths []string, target string) bool {
	for _, v := range paths {
		if strings.Contains(target, v) {
			return true
		}
	}
	return false
}

func TestMain(m *testing.M) {
	mes := "handler test..."
	fmt.Println("start ", mes)

	fmt.Println("di")
	inject()

	code := m.Run()

	fmt.Println("end ", mes)

	os.Exit(code)
}

func TestNewHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// MockApplication
	am := mock_application.NewMockApplicationInterface(ctrl)
	// MockClient
	i18nm := mock_i18n.NewMockI18nClientInterface(ctrl)

	res := handler.NewHandler(am, i18nm)
	if res == nil {
		t.Error("ハンドラの初期化に失敗しました")
	}
}

func TestAssignRoutes(t *testing.T) {
	h := handler.NewHandler(nil, nil)

	// TestAssignRoutes
	h.AssignRoutes(e)
	routes := e.Routes()

	for _, v := range routes {
		paths := getPath()
		if !contains(paths, v.Path) {
			t.Errorf("存在しないルーティングです。\n Path: %v, Method: %v", v.Path, v.Method)
		}
	}
}

func TestGetCtx(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := handler.NewHandler(nil, nil)

	// TestGetCtx
	ctx := h.GetCtx(c)
	if ctx == nil {
		t.Error("echo.Contextをcontext.Contextに変換できませんでした")
	}
}
