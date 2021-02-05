package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/go_db?parseTime=true")
	if err != nil {
		log.Printf("db init error: %v", err)
	}
	DB = db
}
