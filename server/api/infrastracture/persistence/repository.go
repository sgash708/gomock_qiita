package persistence

import (
	"server/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RepositoryInterface interface {
	GetDB() *gorm.DB
}

type persistenceInfo struct {
	db *gorm.DB
}

func Connect(cfg *config.ServerConfig) (RepositoryInterface, error) {
	conn, err := gorm.Open(mysql.Open(cfg.DataSource), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return SetConnection(conn), nil
}

func SetConnection(db *gorm.DB) RepositoryInterface {
	return &persistenceInfo{db: db}
}

func (re *persistenceInfo) GetDB() *gorm.DB {
	return re.db
}

func getDBConnection(ri RepositoryInterface) (*gorm.DB, error) {
	return ri.GetDB(), nil
}
