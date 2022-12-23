package application_test

import (
	"fmt"
	"os"
	"server/api/application"
	"server/config"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMain(m *testing.M) {
	mes := "application test..."
	fmt.Println("start ", mes)

	code := m.Run()

	fmt.Println("end ", mes)

	os.Exit(code)
}

func TestNewApplication(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	a := application.NewApplication(&application.ApplicationBundle{
		ServerConfig: &config.ServerConfig{},
	})

	if a == nil {
		t.Error("アプリケーション層の初期化に失敗しました")
	}
}
