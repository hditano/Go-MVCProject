package initializers

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const create_database = `CREATE TABLE IF NOT EXISTS customers (
id INTEGER PRIMARY KEY AUTOINCREMENT,
firstname TEXT NOT NULL,
lastname TEXT NOT NULL,
address TEXT NOT NULL,
email TEXT NOT NULL
);`

func ConnectToDatabase() {
	os.Remove("./database.db")
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(create_database)
	if err != nil {
		log.Fatal(err)
	}

}
