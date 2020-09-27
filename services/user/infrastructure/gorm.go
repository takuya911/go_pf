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

	protocol := "tcp(" + os.Getenv("INSTANCE_CONNECTION_NAME") + ")"
	connect := user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db, err := gorm.Open("mysql", connect)
	if err != nil {
		return nil, err
	}
	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
