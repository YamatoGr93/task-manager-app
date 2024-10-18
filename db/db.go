package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

var DB *sql.DB // Global database connection variable

// Initialize the database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create tasks table if it doesn't exist
	sqlStmt := `CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        description TEXT,
        due_date TEXT,
        status TEXT
    );`
	if _, err = DB.Exec(sqlStmt); err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
}

// Close the database connection
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Println("Error closing database:", err)
	}
}
