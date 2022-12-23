package persistence_test

import (
	"context"
	"fmt"
	"os"
	"server/api/domain/model"
	"server/api/infrastracture/persistence"
	"server/config"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var d *gorm.DB
var di *Injector

const (
	DriveName = "mysql"
	DSN       = "mock-user:password@tcp(mock-tmysql)/mock-test?charset=utf8mb4&parseTime=true"
)

type Injector struct {
	ctx context.Context
	r   persistence.RepositoryInterface
	br  persistence.BookRepositoryInterface
}

func tearUpDB() {
	var err error
	if d, err = gorm.Open(mysql.Open(DSN), &gorm.Config{}); err != nil {
		panic(err)
	}
}

func tearDownDB() {
	q := d.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped()

	// Book
	q.Table("books").Delete(&model.Book{})

}

func inject() {
	di = &Injector{
		ctx: context.TODO(),
		r:   persistence.SetConnection(d),
		br:  persistence.NewBookRepository(),
	}
}

func TestMain(m *testing.M) {
	mes := "repository test"

	fmt.Println("start ", mes)
	tearUpDB()

	fmt.Println("di")
	inject()

	code := m.Run()

	fmt.Println("end ", mes)
	tearDownDB()

	os.Exit(code)
}

func TestConnectSuccess(t *testing.T) {
	cfg := &config.ServerConfig{
		DataSource: DSN,
	}

	// TestConnect
	db, err := persistence.Connect(cfg)
	if err != nil {
		t.Errorf("dbの接続に失敗しました\n詳細: %s", err)
	}
	if db == nil {
		t.Error("gorm.DBが注入されていません")
	}
}

func TestSetConnectionSuccess(t *testing.T) {
	// TestSetConnection
	if res := persistence.SetConnection(d); res == nil {
		t.Error("コネクション注入に失敗しました")
	}
}

func TestGetDB(t *testing.T) {
	// TestGetDB
	if res := di.r.GetDB(); res == nil {
		t.Error("gormの取得に失敗しています")
	}
}
