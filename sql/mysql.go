package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"username:password@tcp(127.0.0.1:3306)/pos")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")
	defer db.Close()

	fmt.Scanln()
}
