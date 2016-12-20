package main

import (
	"database/sql"
	"fmt"
	// _ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)
	defer db.Close()

	fmt.Println("db", db)

}
