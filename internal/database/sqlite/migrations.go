package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func MigrateDatabase(db *sql.DB) error {
	db, err := sql.Open("sqlite3", "brag.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS days (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			date DATE NOT NULL
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create days table: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			text TEXT NOT NULL,
			day_id INTEGER,
			FOREIGN KEY (day_id) REFERENCES days(id)
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create tasks table: %v", err)
	}

	fmt.Println("Database Tables created successfully")
	return nil
}
