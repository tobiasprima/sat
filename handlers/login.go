package handlers

import (
	"database/sql"
	"fmt"
	"log"
)

func LoginUser(db *sql.DB) bool {
	fmt.Print("Input user email:")
	var inputEmail string
	fmt.Scan(&inputEmail)

	fmt.Print("Input password:")
	var inputPassword string
	fmt.Scan(&inputPassword)

	rows, err := db.Query("SELECT email, password FROM users")
	if err != nil {
		log.Print("Error fetching records: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var email, password string
		err = rows.Scan(&email, &password)
		if err != nil {
			log.Print("Error scanning record: ", err)
		}

		if email == inputEmail {
			if inputPassword == password {
				fmt.Print("\nPassword match.\n")
				return true
			} else {
				fmt.Print("\nPassword doesn't match.\n")
				return false
			}
		}
	}
	fmt.Print("\nUser email couldn't be found.\n")
	return false
}
