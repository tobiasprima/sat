package handlers

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// AddItem adds a new item to the database
func AddItem(db *sql.DB, name string, price float64, stock int) {
	name = strings.ToLower(name)

	query := `
		INSERT INTO items (name, price, stock)
		VALUES ($1, $2, $3)
	`

	_, err := db.Exec(query, name, price, stock)
	if err != nil {
		fmt.Printf("Error adding item: %v\n", err)
		return
	}

	fmt.Println("Item added successfully!")
}

// AddItemInteractive handles user input for adding a seller
func AddItemInteractive(db *sql.DB) {
	fmt.Print("Enter item name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid name format. Please try again.")
		return
	}
	name = name[:len(name)-1]

	var input string

	fmt.Print("Enter item price: ")
	fmt.Scan(&input)
	price, err := strconv.ParseFloat(input, 64)
	if err != nil || price < 0 {
		fmt.Println("Invalid price. Please try again.")
		return
	}

	fmt.Print("Enter item stock: ")
	fmt.Scan(&input)
	stock, err := strconv.ParseInt(input, 10, 64)
	if err != nil || stock < 0 {
		fmt.Println("Invalid stock amount. Please try again.")
		return
	}

	AddItem(db, name, price, int(stock))
}

// DeleteItem deletes an item from the database
func DeleteItem(db *sql.DB, name string) {
	query := `
		DELETE FROM items
		WHERE name = $1
	`

	result, err := db.Exec(query, name)
	if err != nil {
		fmt.Printf("Error deleting item: %v\n", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("No item found with the given name.")
		return
	}

	fmt.Println("Item deleted successfully!")
}

// DeleteItemInteractive handles user input for deleting an item
func DeleteItemInteractive(db *sql.DB) {
	fmt.Print("Enter the name of the item to delete: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid name format. Please try again.")
		return
	}
	name = name[:len(name)-1]
	name = strings.ToLower(name)

	DeleteItem(db, name)
}

// UpdateItem updates items from the databse
func UpdateItem(db *sql.DB, column, newValue string, name string) error {
	// Prepare the SQL query dynamically for the specified column
	query := fmt.Sprintf(`
		UPDATE items
		SET %s = $1
		WHERE name = $2
	`, column)

	// Execute the query with the provided new value and item name
	_, err := db.Exec(query, newValue, name)
	if err != nil {
		return fmt.Errorf("error updating item: %v", err)
	}

	return nil // Return nil if no error occurs
}

// UpdateItemInteractive handles user input for updating items
// Users can choose to edit the name, price, or stock, and see current values before making changes.
func UpdateItemInteractive(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin) // Create buffered reader for user input
	fmt.Print("Enter the name of the item to edit: ")
	name, _ := reader.ReadString('\n') // Read the full line of input
	name = strings.TrimSpace(name)     // Trim any extra spaces or newlines

	// Variables to hold the current item details
	var currentName string
	var currentPrice float64
	var currentStock int

	// Fetch current item details from the database
	err := db.QueryRow(`
		SELECT name, price, stock
		FROM items
		WHERE name = $1
	`, name).Scan(&currentName, &currentPrice, &currentStock)

	if err != nil {
		fmt.Println("Item not found.") // Notify if the item doesn't exist
		return
	}

	// Loop to allow multiple edits until the user chooses to exit
	for {
		// Display the current item details for reference
		fmt.Printf("\nCurrent Details:\nName: %s | Price: %.2f | Stock: %d\n", currentName, currentPrice, currentStock)

		// Display the options menu for editing
		fmt.Println("\nWhat would you like to edit?")
		fmt.Println("1. Edit name")
		fmt.Println("2. Edit price")
		fmt.Println("3. Edit stock")
		fmt.Println("4. Done")
		fmt.Print("Your choice: ")

		var choice int
		fmt.Scan(&choice) // Get user input for menu choice

		switch choice {
		case 1:
			// Handle editing the item name
			fmt.Print("Enter new name: ")
			newName, _ := reader.ReadString('\n') // Read the full line of input
			newName = strings.TrimSpace(newName)
			if newName == currentName {
				fmt.Println("No changes made to the name.") // No update if the name is unchanged
			} else {
				err := UpdateItem(db, "name", newName, currentName)
				if err != nil {
					fmt.Println(err) // Print error if update fails
					return
				}
				currentName = newName // Update the local variable to reflect the change
				fmt.Printf("Name updated successfully! New Name: %s\n", currentName)
			}
		case 2:
			// Handle editing the item price
			fmt.Print("Enter new price: ")
			var newPrice float64
			fmt.Scan(&newPrice)
			if newPrice == currentPrice {
				fmt.Println("No changes made to the price.") // No update if the price is unchanged
			} else {
				err := UpdateItem(db, "price", fmt.Sprintf("%.2f", newPrice), currentName)
				if err != nil {
					fmt.Println(err) // Print error if update fails
					return
				}
				currentPrice = newPrice // Update the local variable to reflect the change
				fmt.Printf("Price updated successfully! New Price: %.2f\n", currentPrice)
			}
		case 3:
			// Handle editing the item stock
			fmt.Print("Enter new stock: ")
			var newStock int
			fmt.Scan(&newStock)
			if newStock == currentStock {
				fmt.Println("No changes made to the stock.") // No update if the stock is unchanged
			} else {
				err := UpdateItem(db, "stock", fmt.Sprintf("%d", newStock), currentName)
				if err != nil {
					fmt.Println(err) // Print error if update fails
					return
				}
				currentStock = newStock // Update the local variable to reflect the change
				fmt.Printf("Stock updated successfully! New Stock: %d\n", currentStock)
			}
		case 4:
			// Exit the editing menu
			fmt.Println("Exiting item edit menu.")
			return
		default:
			// Handle invalid menu choices
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
