package main

import (
	    "context"
	    "fmt"
	    "log"
	    "time"

	    "github.com/jackc/pgx/v5"
)

type QueryResult struct {
    Timestamp time.Time
	Version string
}

func fatality(err error, message string) {
	if err != nil {
		log.Fatalf("Fatal error [%s]: %v", message, err)
	}
}

func readQuery(conn *pgx.Conn) QueryResult {
	var r QueryResult

	err := conn.QueryRow(context.Background(), 
		"SELECT now() AS timestamp, version() AS version",
	).Scan(&r.Timestamp, &r.Version)
	fatality(err, "Query failed: %v\n")

	return r
}

func readQueryString(conn *pgx.Conn, query string) string {
	var r string

	err := conn.QueryRow(context.Background(), query).Scan(&r)
	fatality(err, "Query failed: %v\n")
	return r
}

func main() {
    conn, err := pgx.Connect(context.Background(),
		"postgres://echoapp:echoapp@localhost:5432/echoapp")
	fatality(err, "Unable to connect to database: %v\n")
    defer conn.Close(context.Background())

	res := readQuery(conn)
	fmt.Printf("Timestamp: %v\n", res.Timestamp)
	fmt.Printf("Version: %s\n", res.Version)

	sum := readQueryString(conn, "select 2+2")
	fmt.Printf("Sum: %s\n", sum)
}

