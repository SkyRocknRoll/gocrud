package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var (
		age          int
		fname, lname string
	)
	// Intializing the db
	db, err := sql.Open("mysql", "go:password@tcp(127.0.0.1:3306)/godb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Checking database connectivity
	err = db.Ping()
	if err != nil {
		log.Fatal("Connection failed !!!!")

	}
	log.Println("Successfully connected to Database!!")

	rows, err := db.Query("select * from users where age=?", 38)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(&fname, &lname, &age)
		if err != nil {
			log.Fatal("Row Fetching failed!!!")
		}
		log.Println(fname, lname, age)
	}

}
