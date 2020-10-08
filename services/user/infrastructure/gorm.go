package infrastructure

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// NewGormConnect function
func NewGormConnect() (*gorm.DB, error) {

	var connect string
	env := os.Getenv("ENV")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	if "dev" == env {
		dbHostName := os.Getenv("DB_HOST")
		connect = user + ":" + pass + "@tcp(" + dbHostName + ":" + os.Getenv("DB_PORT") + ")/" + dbName + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	} else {
		instanceConnectionName := os.Getenv("INSTANCE_CONNECTION_NAME")
		connect = user + ":" + pass + "@unix(/cloudsql/" + instanceConnectionName + ")/" + dbName + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	}

	db, err := gorm.Open("mysql", connect)
	if err != nil {
		return nil, err
	}

	return db, nil

}
