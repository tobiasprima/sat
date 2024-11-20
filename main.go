package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db_name := "postgres"
	db_user := "postgres.wyxrurmpypoqythiftvq"
	db_pass := "PSiRcsVEJ9NsTjv8"
	db_host := "aws-0-ap-southeast-1.pooler.supabase.com"

	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require", db_name, db_user, db_pass, db_host))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	fmt.Print("Input user name:")
	var username string
	fmt.Scan(&username)

	fmt.Print("Input password:")
	var password string
	fmt.Scan(&password)

}
