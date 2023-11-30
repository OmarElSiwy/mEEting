package backend

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// SimpleFunction takes a string and returns a modified string
func SimpleFunction(name string) string {
	return fmt.Sprintf("Hello, %s from Go!", name)
}

func database() *sql.DB {
	connStr := "user=username password=password dbname=yourdbname sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func ExecuteSQL(db *sql.DB, value1 string, value2 string) {
	_, err := db.Exec("INSERT INTO yourtable (column1, column2) VALUES ($1, $2)", value1, value2)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT column1, column2 FROM yourtable WHERE column1 = $1", value1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	HandleData(rows)
}

func HandleData(rows *sql.Rows) {
	for rows.Next() {
		var column1, column2 string
		err := rows.Scan(&column1, &column2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(column1, column2)
	}
}
