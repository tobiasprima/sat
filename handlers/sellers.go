package handlers

import (
	"database/sql"
	"fmt"
)

func AddSeller(db *sql.DB){
	fmt.Print("Enter seller name: ")
	var name string
	fmt.Scan(&name)

	fmt.Print("Enter seller email: ")
	var email string
	fmt.Scan(&email)

	query := `
	INSERT INTO Sellers (name, email)
	VALUES ($1, $2)
	`

	_, err := db.Exec(query, name ,email)
	if err != nil {
		fmt.Printf("Error adding seller: %v\n", err)
		return
	}

	fmt.Println("Seller added successfully!")
}