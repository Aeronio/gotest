package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Notice how it's capitalized DB,we will be using this DB in other packages
var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
