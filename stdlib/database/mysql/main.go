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

/*
** Fetch data from the database
 */
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

	// Always check for errors, even after for rows.Next()
	fatality(rows.Err())
	log.Println("Records count:", rowsCount)
}

func main() {
	// Connection are established in a lazy fashion so following line just set
	// up the connect without openning it
	db, err := sql.Open("mysql", "root:go-mysql@tcp(127.0.0.1:3306)/main")
	fatality(err)

	// any method call on the db (sql.DB) will open up the connection
	err = db.Ping()
	fatality(err)

	// If we reach here it means database connection is established
	log.Println("Database connection established.")

	fetchData(db)
	defer db.Close()
}
