package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {

	var err error

	DB, err = sql.Open("sqlite", "./monitoring.db")

	if err != nil {
		log.Fatal(err)
	}

	DB.Exec("PRAGMA journal_mode=WAL;") // allows r & w to happen at the same time

	query := `
	CREATE TABLE IF NOT EXISTS system_metrics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip_address TEXT,
		cpu_usage REAL,
		memory_usage REAL,
		disk_usage REAL,
		timestamp TEXT
	);
	`

	_, err = DB.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("SQLite database initialized")
}

