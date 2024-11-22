package handlers

import (
	"SAT/utils"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
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

// UpdateSeller updates a specific column of a seller in the database with a new value
func UpdateSeller(db *sql.DB, column, newValue, email string) error {
	// Validate column to prevent SQL injection
	validColumns := map[string]bool{
		"name":  true,
		"email": true,
	}
	if !validColumns[column] {
		return fmt.Errorf("invalid column specified: %s", column)
	}

	// Prepare the query safely
	query := fmt.Sprintf(`
		UPDATE sellers
		SET %s = $1
		WHERE email = $2
	`, column)

	_, err := db.Exec(query, newValue, email)
	if err != nil {
		return fmt.Errorf("error updating seller: %v", err)
	}

	return nil
}

// UpdateSellerInteractive allows users to interactively update a seller's details
func UpdateSellerInteractive(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for the seller's email
	fmt.Print("Enter the email of the seller to edit: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	// Variables to hold the current seller details
	var currentName, currentEmail string

	// Fetch the current details of the seller from the database
	err := db.QueryRow(`
		SELECT name, email
		FROM sellers
		WHERE email = $1
	`, email).Scan(&currentName, &currentEmail)

	if err != nil {
		fmt.Printf("Seller with email '%s' not found.\n", email)
		return
	}

	// Loop to allow multiple edits until the user chooses to exit
	for {
		// Display the current details of the seller
		fmt.Printf("\nCurrent Details:\nName: %s | Email: %s\n", currentName, currentEmail)

		// Menu options for updating seller details
		fmt.Println("\nWhat would you like to edit?")
		fmt.Println("1. Edit name")
		fmt.Println("2. Edit email")
		fmt.Println("3. Done")
		fmt.Print("Your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Handle editing the seller name
			fmt.Print("Enter new name: ")
			newName, _ := reader.ReadString('\n')
			newName = strings.TrimSpace(newName)

			if newName == currentName {
				fmt.Println("No changes made to the name.")
			} else {
				err := UpdateSeller(db, "name", newName, currentEmail)
				if err != nil {
					fmt.Println(err)
					return
				}
				currentName = newName // Update the local variable to reflect the change
				fmt.Printf("Name updated successfully! New Name: %s\n", currentName)
			}
		case 2:
			// Handle editing the seller email
			fmt.Print("Enter new email: ")
			newEmail, _ := reader.ReadString('\n')
			newEmail = strings.TrimSpace(newEmail)

			if newEmail == currentEmail {
				fmt.Println("No changes made to the email.")
			} else {
				err := UpdateSeller(db, "email", newEmail, currentEmail)
				if err != nil {
					fmt.Println(err)
					return
				}
				currentEmail = newEmail // Update the local variable to reflect the change
				fmt.Printf("Email updated successfully! New Email: %s\n", currentEmail)
			}
		case 3:
			// Exit the update menu
			fmt.Println("Exiting seller update menu.")
			return
		default:
			// Handle invalid menu options
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
