package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func Connect() {
	var err error
	dataSourceName := "root:12345678@tcp(127.0.0.1:3306)/userdb"
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")
}
