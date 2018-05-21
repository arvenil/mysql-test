package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root@/"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connection works.")
}
