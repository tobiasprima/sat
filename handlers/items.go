package handlers

import (
	"database/sql"
	"fmt"
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
	var name string
	fmt.Scan(&name)

	fmt.Print("Enter item price: ")
	var price float64
	fmt.Scan(&price)

	fmt.Print("Enter item stock: ")
	var stock int
	fmt.Scan(&stock)

	AddItem(db, name, price, stock)
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
	var name string
	fmt.Scan(&name)

	DeleteItem(db, name)
}
