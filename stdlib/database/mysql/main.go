package main

import (
	"database/sql"
	"fmt"
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

func preparingQueries(db *sql.DB) {
	var (
		id   int
		name string
	)

	log.Println("Preparing queries example ...")

	stmt, err := db.Prepare("select id, name from person where id < ?")
	fatality(err)
	defer stmt.Close()

	rows, err := stmt.Query(5)
	fatality(err)
	defer rows.Close()

	out := ""
	for rows.Next() {
		err := rows.Scan(&id, &name)
		fatality(err)
		if out == "" {
			out += fmt.Sprintf("(%02d: %s)", id, name)
		} else {
			out += fmt.Sprintf(", (%02d: %s)", id, name)
		}
	}
	log.Println(out)

	// Always check for errors, even after for rows.Next()
	fatality(rows.Err())
}

func singleRowQueries(db *sql.DB) {
	log.Println("Single-row queries example ...")

	var id int = 3
	var name string

	err := db.QueryRow("select name from person where id = ?", 1).Scan(&name)
	fatality(err)
	log.Println("Name for id", id, "is", name)
}

func preparedSingleRowQueries(db *sql.DB) {
	log.Println("Single-row prepared queries example ...")

	var id int = 5
	var name string

	stmt, err := db.Prepare("select name from person where id = ?")
	fatality(err)
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&name)
	fatality(err)

	log.Println("Name for id", id, "is", name)
}

func deleteId(db *sql.DB, id int) {
	log.Println("Delete ID example ...")

	_, err := db.Exec("DELETE FROM person WHERE id = ?", id)
	fatality(err)

	log.Println("Person with id", id, "has been deleted")
}

func insertData(db *sql.DB) {
	log.Println("Insert data example ...")

	stmt, err := db.Prepare("INSERT INTO person (name) VALUES(?)")
	fatality(err)

	res, err := stmt.Exec("Dolly")
	fatality(err)

	lastId, err := res.LastInsertId()
	fatality(err)

	rowCount, err := res.RowsAffected()
	fatality(err)

	log.Printf("id: %d, affected: %d\n", lastId, rowCount)

	deleteId(db, int(lastId))
}

func main() {
	// Connection are established in a lazy fashion so following line just set
	// up the connect without openning it
	db, err := sql.Open("mysql", "root:go-mysql@tcp(127.0.0.1:3306)/main")
	fatality(err)
	defer db.Close()

	// any method call on the db (sql.DB) will open up the connection
	err = db.Ping()
	fatality(err)

	// If we reach here it means database connection is established
	log.Println("Database connection established.")

	fetchData(db)
	preparingQueries(db)
	singleRowQueries(db)
	preparedSingleRowQueries(db)
	insertData(db)
}
