package driver

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Connect -> db connection
func Connect() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_CONN"))
	logFatal(err)

	err = db.Ping()
	logFatal(err)
	return db
}