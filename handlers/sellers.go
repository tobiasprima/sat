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

// UpdateSeller updates a seller from the database
func UpdateSeller(db *sql.DB, column, newValue, email string) error {
	// Validate column to prevent SQL injection.
	validColumns := map[string]bool{
		"name":  true,
		"email": true,
	}
	if !validColumns[column] { // Check if the specified column is allowed.
		return fmt.Errorf("invalid column specified: %s", column)
	}

	// Dynamically prepare the SQL query with the validated column name.
	query := fmt.Sprintf(`UPDATE sellers SET %s = $1 WHERE email = $2`, column)

	// Execute the query, binding the new value and email parameters to prevent injection.
	_, err := db.Exec(query, newValue, email)
	if err != nil {
		return fmt.Errorf("error updating seller: %v", err) // Return error with detailed message.
	}

	return nil // Return nil if no error occurs.
}

// UpdateSellerInteractive handles user input for updating a seller
// This function provides a menu to update either the seller's name or email.
func UpdateSellerInteractive(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin) // Use bufio.Reader for multi-word input.

	// Prompt user for the seller's email to identify the record.
	fmt.Print("Enter the email of the seller to edit: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err) // Handle input errors.
		return
	}
	email = strings.TrimSpace(email) // Remove any leading/trailing whitespace.

	// Variables to hold the current seller details.
	var currentName, currentEmail string

	// Fetch the seller's current details from the database.
	err = db.QueryRow(`
		SELECT name, email
		FROM sellers
		WHERE email = $1
	`, email).Scan(&currentName, &currentEmail)

	if err != nil {
		fmt.Printf("Seller with email '%s' not found.\n", email)
		return
	}

	// Allow the user to make multiple updates in a loop.
	for {
		// Display the current details of the seller for reference.
		fmt.Printf("\nCurrent Details:\nName: %s | Email: %s\n", currentName, currentEmail)

		// Display the menu for updating seller details.
		fmt.Println("\nWhat would you like to edit?")
		fmt.Println("1. Edit name")
		fmt.Println("2. Edit email")
		fmt.Println("3. Done")
		fmt.Print("Your choice: ")

		var choice int
		fmt.Scan(&choice) // Get user's choice.

		switch choice {
		case 1:
			// Handle updating the seller's name.
			fmt.Print("Enter new name: ")
			newName, _ := reader.ReadString('\n') // Read multi-word input.
			newName = strings.TrimSpace(newName)  // Remove extra spaces.
			newName = utils.CapitalizeName(newName)

			// Check if the new name is the same as the current name.
			if strings.EqualFold(newName, currentName) {
				fmt.Println("No changes made to the name.")
			} else {
				// Perform the database update.
				err := UpdateSeller(db, "name", newName, currentEmail)
				if err != nil {
					fmt.Println(err) // Print error if the update fails.
					return
				}
				currentName = newName // Update local variable to reflect the change.
				fmt.Printf("Name updated successfully! New Name: %s\n", currentName)
			}
		case 2:
			// Handle updating the seller's email.
			fmt.Print("Enter new email: ")
			newEmail, _ := reader.ReadString('\n') // Read multi-word input.
			newEmail = strings.TrimSpace(newEmail) // Remove extra spaces.
			if !utils.ValidateEmail(newEmail) {
				fmt.Println("Invalid email format. Please try again.")
				return
			}

			// Check if the new email is the same as the current email.
			if strings.EqualFold(newEmail, currentEmail) {
				fmt.Println("No changes made to the email.")
			} else {
				// Perform the database update.
				err := UpdateSeller(db, "email", newEmail, currentEmail)
				if err != nil {
					fmt.Println(err) // Print error if the update fails.
					return
				}
				currentEmail = newEmail // Update local variable to reflect the change.
				fmt.Printf("Email updated successfully! New Email: %s\n", currentEmail)
			}
		case 3:
			// Exit the update menu.
			fmt.Println("Exiting seller update menu.")
			return
		default:
			// Handle invalid menu choices.
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
