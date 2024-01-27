package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Init(databasePath string) error {
	var err error

	db, err = sql.Open("sqlite3", databasePath)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Database connected successfully")
	return nil
}

func Close() {
	if db != nil {
		db.Close()
		fmt.Println("Database connection closed")
	}
}

func GetDB() *sql.DB {
	return db
}
