package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func fatality(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fetchData(db *sql.DB) {
	var (
		id        int
		name      string
		rowsCount int = 0
	)

	log.Println("Trying to fetch person data ...")

	rows, err := db.Query("select * from person")
	fatality(err)

	defer rows.Close()

	for rows.Next() {
		rowsCount++
		err := rows.Scan(&id, &name)
		fatality(err)

		log.Println(id, name)
	}

	log.Println("Records count:", rowsCount)
	fatality(rows.Err())
}

func main() {
	db, err := sql.Open("mysql", "root:go-mysql@tcp(127.0.0.1:3306)/main")
	fatality(err)

	err = db.Ping()
	fatality(err)
	log.Println("Database connection established.")

	fetchData(db)
	defer db.Close()
}
