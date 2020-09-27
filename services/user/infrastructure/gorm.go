package infrastructure

import (
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// NewGormConnect function
func NewGormConnect() (*gorm.DB, error) {
	mysqlConf := &mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("INSTANCE_CONNECTION_NAME"),
		DBName:               os.Getenv("DB_NAME"),
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  time.Local,
		AllowNativePasswords: true,
	}

	db, err := gorm.Open("mysql", mysqlConf.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
