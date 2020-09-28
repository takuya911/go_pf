package infrastructure

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// NewGormConnect function
func NewGormConnect() (*gorm.DB, error) {

	// get env
	// user := os.Getenv("DB_USER")
	// pass := os.Getenv("DB_PASS")
	// dbName := os.Getenv("DB_NAME")

	// protocol := "tcp(" + os.Getenv("INSTANCE_CONNECTION_NAME") + ")"
	// connect := user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	// db, err := gorm.Open("mysql", connect)
	var (
		dbUser                 = os.Getenv("DB_USER")                  // e.g. 'my-db-user'
		dbPwd                  = os.Getenv("DB_PASS")                  // e.g. 'my-db-password'
		instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
		dbName                 = "go_pf"                               // e.g. 'my-database'
	)

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	var dbURI string
	dbURI = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)

	// dbPool is the pool of database connections.
	db, err := gorm.Open("mysql", dbURI)

	if err != nil {
		return nil, err
	}
	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
