package handlers

import (
	"SAT/utils"
	"bufio"
	"database/sql"
	"fmt"
	"os"
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
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid name format. Please try again.")
		return
	}
	name = name[:len(name)-1]

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

// UpdateSeller updates seller from the databse
func UpdateSeller(db *sql.DB, column, newValue string, email string) error {
	// Prepare the SQL query dynamically for the specified column
	query := fmt.Sprintf(`
		UPDATE sellers
		SET %s = $1
		WHERE email = $2
	`, column)

	// Execute the query with the provided new value and item name
	_, err := db.Exec(query, newValue, email)
	if err != nil {
		return fmt.Errorf("error updating item: %v", err)
	}

	return nil // Return nil if no error occurs
}

// UpdateSellerInteractive handles user input for updating sellers
// Users can choose to edit the name, or email, and see current values before making changes.
func UpdateSellerInteractive(db *sql.DB) {
	fmt.Print("Enter the email of the seller to edit: ")
	var email string
	fmt.Scan(&email)

	// Variables to hold the current item details
	var currentName, currentEmail string

	// Fetch current item details from the database
	err := db.QueryRow(`
		SELECT name, email
		FROM sellers
		WHERE email = $1
	`, email).Scan(&currentName, &currentEmail)

	if err != nil {
		fmt.Println("Seller not found.") // Notify if the item doesn't exist
		return
	}

	// Loop to allow multiple edits until the user chooses to exit
	for {
		// Display the current user details for reference
		fmt.Printf("\nCurrent Details:\nName: %s | Email: %s\n", currentName, currentEmail)

		// Display the options menu for editing
		fmt.Println("\nWhat would you like to edit?")
		fmt.Println("1. Edit name")
		fmt.Println("2. Edit email")
		fmt.Println("3. Done")
		fmt.Print("Your choice: ")

		var choice int
		fmt.Scan(&choice) // Get user input for menu choice

		switch choice {
		case 1:
			// Handle editing the item name
			fmt.Print("Enter new name: ")
			reader := bufio.NewReader(os.Stdin)
			newName, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Invalid name format. Please try again.")
				return
			}
			newName = newName[:len(newName)-1]
			newName = utils.CapitalizeName(newName)

			if newName == currentName {
				fmt.Println("No changes made to the name.") // No update if the name is unchanged
			} else {
				err := UpdateSeller(db, "name", newName, email)
				if err != nil {
					fmt.Println(err) // Print error if update fails
					return
				}
				currentName = newName // Update the local variable to reflect the change
				fmt.Println("Name updated successfully!")
			}
		case 2:
			// Handle editing the item price
			fmt.Print("Enter new email: ")
			var newEmail string
			fmt.Scan(&newEmail)

			if !utils.ValidateEmail(newEmail) {
				fmt.Println("Invalid email format. Please try again.")
				return
			}

			if newEmail == currentEmail {
				fmt.Println("No changes made to the email.") // No update if the price is unchanged
			} else {
				err := UpdateSeller(db, "email", newEmail, email)
				if err != nil {
					fmt.Println(err) // Print error if update fails
					return
				}
				currentEmail = newEmail // Update the local variable to reflect the change
				fmt.Println("Email updated successfully!")
			}
		case 3:
			// Exit the editing menu
			fmt.Println("Exiting item edit menu.")
			return
		default:
			// Handle invalid menu choices
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
