package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func createTable(db *sql.DB) {
	query := `create table users (
		id int auto_increment,
		username text not null,
		password text not null,
		created_at datetime,
		primary key(id)
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("The users table has been created successfully")
	}
}

func insertUser(db *sql.DB) {
	query := `insert into users (username, password, created_at) values (?, ?, ?)`

	username := "adam"
	password := "pass"
	createdAt := time.Now()

	res, err := db.Exec(query, username, password, createdAt)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("A new record has been inserted into users table")
	}

	userId, err := res.LastInsertId()
	log.Println("New record's ID:", userId)
}

func getUserById(db *sql.DB, userId int) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	query := `select * from users where id = ?`
	err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully read fields for id:", id)

		// Display the fields
		fmt.Println("User ID:\t", id)
		fmt.Println("Username:\t", username)
		fmt.Println("Password:\t", password)
		fmt.Println("The date:\t", createdAt)
	}
}

func getAllUsers(db *sql.DB) {
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query("select * from users")
	defer rows.Close()

	var users []user
	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			var u user
			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Println("Error getting user record")
			} else {
				users = append(users, u)
			}
		}
	}

	fmt.Println("User records:")
	for index, user := range users {
		fmt.Println("Index:", index)
		fmt.Println("Username:", user.username)
		fmt.Println("Password:", user.password)
		fmt.Println("The Date:", user.createdAt)
		fmt.Println()
	}
}

func main() {
	// Open a connection to database
	db, err := sql.Open("mysql", "goweb:goweb@(127.0.0.1:3306)/goweb?parseTime=true")
	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %s", err))
	} else {
		fmt.Println("Database connection is successful")
	}

	fmt.Println(db.Ping())
	createTable(db)
	insertUser(db)
	getUserById(db, 1)
	getAllUsers(db)
}
