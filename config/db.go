package config

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var dbOnce = &sync.Once{}
var client *sql.DB

func NewDBClient() *sql.DB {
	dbOnce.Do(func() {
		DRIVER := os.Getenv("DB_DRIVER")
		URL := os.Getenv("DB_URL")
		db, err := sql.Open(DRIVER, URL)
		if err != nil {
			panic(err)
		}
		if err := db.Ping(); err != nil {
			panic(err)
		}
		client = db
	})
	return client
}
