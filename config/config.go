package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// ConnectDatabase connects to database on supabase it returns pointer to database
func ConnectDatabase() *sql.DB {
	// database connection data
	db_name := "postgres"
	db_user := "postgres.wyxrurmpypoqythiftvq"
	db_pass := "PSiRcsVEJ9NsTjv8"
	db_host := "aws-0-ap-southeast-1.pooler.supabase.com"

	// opening connection
	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require", db_name, db_user, db_pass, db_host))
	if err != nil {
		panic(err)
	}

	// pinging connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Print("Database connected!\n\n")

	return db
}
