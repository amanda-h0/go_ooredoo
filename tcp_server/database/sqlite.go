package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {

	var err error

	DB, err = sql.Open("sqlite", "/home/amanda/ooredoo_26/tcp_server/monitoring.db")

	if err != nil {
		log.Fatal(err)
	}


	query := `
	CREATE TABLE IF NOT EXISTS system_metrics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip_address TEXT,
		cpu_usage REAL,
		memory_usage REAL,
		disk_usage REAL,
		timestamp INTEGER
	);
	`

	_, err = DB.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("SQLite database initialized")
}