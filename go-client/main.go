package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	resp, err := http.Get("http://go-postgres-api:8080/")
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	defer resp.Body.Close()

	ConnectToDatabase()
}

func ConnectToDatabase() {
	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@postgres:5432/postgres?sslmode=disable")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Ping()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("db connected")
}
