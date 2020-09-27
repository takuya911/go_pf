package infrastructure

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// NewGormConnect function
func NewGormConnect() (*gorm.DB, error) {

	// get env
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	instance := "go-pf-290702:asia-northeast1:go-pf-mysql"
	connect := user + ":" + pass + "@unix(/cloudsql/" + instance + ")/" + dbName

	db, err := gorm.Open("mysql", connect)
	if err != nil {
		return nil, err
	}
	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
