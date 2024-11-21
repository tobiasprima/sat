package handlers

import (
	"SAT/utils"
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

	// Format the name and validate the email
	name = utils.CapitalizeName(name)
	if !utils.ValidateEmail(email) {
		fmt.Println("Invalid email format. Please try again.")
		return
	}

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

func DeleteSeller(db *sql.DB){
	fmt.Print("Enter the email of the seller to delete: ")
	var sellerEmail string
	fmt.Scan(&sellerEmail)

	query := `
	DELETE FROM Sellers
	WHERE email = $1
	`
	result, err := db.Exec(query, sellerEmail)
	if err != nil {
		fmt.Printf("Error deleting seller: %v\n", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("No seller found with the given email.")
		return
	}

	fmt.Println("Seller deleted successfully")
}