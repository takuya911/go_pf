package infrastructure

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// NewGormConnect function
func NewGormConnect() (*sql.DB, error) {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	protocol := "tcp(" + os.Getenv("INSTANCE_CONNECTION_NAME") + ")"
	connect := user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := sql.Open("mysql", connect)
	if err != nil {
		return nil, err
	}

	return db, nil

}
