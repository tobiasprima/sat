package handlers

import (
	"SAT/utils"
	"database/sql"
	"fmt"
)

// AddSeller adds a new seller to the database
func AddSeller(db *sql.DB, name string, email string) {
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

	_, err := db.Exec(query, name, email)
	if err != nil {
		fmt.Printf("Error adding seller: %v\n", err)
		return
	}

	fmt.Println("Seller added successfully!")
}

// AddSellerInteractive handles user input for adding a seller
func AddSellerInteractive(db *sql.DB) {
	fmt.Print("Enter seller name: ")
	var name string
	fmt.Scan(&name)

	fmt.Print("Enter seller email: ")
	var email string
	fmt.Scan(&email)

	AddSeller(db, name, email)
}


// DeleteSeller deletes a seller from the database
func DeleteSeller(db *sql.DB, email string) {
	query := `
		DELETE FROM Sellers
		WHERE email = $1
	`

	result, err := db.Exec(query, email)
	if err != nil {
		fmt.Printf("Error deleting seller: %v\n", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("No seller found with the given email.")
		return
	}

	fmt.Println("Seller deleted successfully!")
}

// DeleteSellerInteractive handles user input for deleting a seller
func DeleteSellerInteractive(db *sql.DB) {
	fmt.Print("Enter the email of the seller to delete: ")
	var email string
	fmt.Scan(&email)

	DeleteSeller(db, email)
}