package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() error {
	var err error

	dsn := "root:@tcp(localhost:3306)/momon"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}
