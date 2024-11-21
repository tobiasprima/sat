package main

import (
	"SAT/config"
	"SAT/handlers"
	"fmt"
)

func main() {
	db := config.ConnectDatabase()
	defer db.Close()

	fmt.Println("Welcome to the SAT system!")

	// for {
	// 	fmt.Println("Please log in to continue.")
	// 	// Login handler disini
	// 	// if handlers.LoginUser(db){
	// 	// 	break
	// 	// }
	// 	fmt.Println("Login failed. Please try again.\n")
	// }

	// CLI options
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Item")
		fmt.Println("2. Delete Item")
		fmt.Println("3. Add Seller")
		fmt.Println("4. Delete Seller")
		fmt.Println("5. Total Sales Report")
		fmt.Println("6. List of Items(Most Popular)")
		fmt.Println("7. List of Sellers")
		fmt.Println("8. Exit")
		fmt.Print("Your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		// case 1:
		// 	handlers.AddItem(db)
		// case 2:
		// 	handlers.DeleteItem(db)
		case 3:
			handlers.AddSeller(db)
		// case 4:
		// 	handlers.DeleteSeller(db)
		// case 5:
		// 	handlers.TotalSalesReport(db)
		// case 6:
		// 	handlers.ListPopularItems(db)
		// case 7:
		// 	handlers.ListSellers(db)
		case 8:
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
