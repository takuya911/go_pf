package infrastructure

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

// NewGormConnect function
func NewGormConnect() (*gorm.DB, error) {
	// mysqlConf := &mysql.Config{
	// 	User:                 os.Getenv("DB_USER"),
	// 	Passwd:               os.Getenv("DB_PASS"),
	// 	Net:                  "tcp",
	// 	Addr:                 os.Getenv("INSTANCE_CONNECTION_NAME"),
	// 	DBName:               os.Getenv("DB_NAME"),
	// 	ParseTime:            true,
	// 	Collation:            "utf8mb4_unicode_ci",
	// 	Loc:                  time.Local,
	// 	AllowNativePasswords: true,
	// }

	// db, err := gorm.Open("mysql", mysqlConf.FormatDSN())
	// if err != nil {
	// 	return nil, err
	// }
	// if err := db.DB().Ping(); err != nil {
	// 	return nil, err
	// }

	// return db, nil
	var (
		dbUser                 = os.Getenv("DB_USER")
		dbPwd                  = os.Getenv("DB_PASS")
		instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
		dbName                 = os.Getenv("DB_NAME")
	)

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}
	dbURI := fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)
	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	return db, nil

}
