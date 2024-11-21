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
