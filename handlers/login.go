package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/term"
)

// LoginUser asks for email and password until succeed
func LoginUser(db *sql.DB) bool {
	// prompt and ask for user email
	fmt.Print("Input user email:")
	var inputEmail string
	fmt.Scan(&inputEmail)

	// prompt and ask for user password
	fmt.Print("Input password:")
	passwordByte, _ := term.ReadPassword(0)
	inputPassword := string(passwordByte)

	// fetch email and passwords from databse
	rows, err := db.Query("SELECT email, password FROM users")
	if err != nil {
		log.Print("Error fetching records: ", err)
	}
	defer rows.Close()

	// scan rows for matched user email and password
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
