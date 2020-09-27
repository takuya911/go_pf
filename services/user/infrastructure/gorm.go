package infrastructure

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"google.golang.org/appengine"
)

// NewGormConnect function
func NewGormConnect() (*gorm.DB, error) {

	// get env
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	var connect string
	if appengine.IsDevAppServer() {
		instance := "go-pf-290702:asia-northeast1:go-pf-mysql"
		connect = user + ":" + pass + "@unix(/cloudsql/" + instance + ")/" + dbName

	} else {
		protocol := "tcp(" + os.Getenv("INSTANCE_CONNECTION_NAME") + ")"
		connect = user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	}

	db, err := gorm.Open("mysql", connect)
	if err != nil {
		return nil, err
	}
	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
