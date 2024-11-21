package main

import (
	"SAT/config"
	"SAT/handlers"
	"fmt"
)

func main() {
	db := config.ConnectDatabase()
	defer db.Close()
	/*
		for {
			fmt.Println("Please log in to continue.")
			if handlers.LoginUser(db) {
				fmt.Println("Welcome to the SAT system!")
				break
			}
			fmt.Print("Login failed. Please try again.\n\n")
		}*/

	// CLI options
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Item")
		fmt.Println("2. Delete Item")
		fmt.Println("3. Update Item")
		fmt.Println("4. Add Seller")
		fmt.Println("5. Delete Seller")
		fmt.Println("6. Update Seller")
		fmt.Println("7. Total Sales Report")
		fmt.Println("8. Popular Items Report")
		fmt.Println("9. Seller Ranking Report")
		fmt.Println("10. Exit")
		fmt.Print("Your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			handlers.AddItemInteractive(db)
		case 2:
			handlers.DeleteItemInteractive(db)
		case 3:
			handlers.UpdateItemInteractive(db)
		case 4:
			handlers.AddSellerInteractive(db)
		case 5:
			handlers.DeleteSellerInteractive(db)
		// case 6:
		// 	handlers.EditSellerInteractive(db)
		case 7:
			handlers.TotalSalesReport(db)
		case 8:
			handlers.PopularItemsReport(db)
		case 9:
			handlers.SellerRankingReport(db)
		case 10:
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
	fmt.Println("\nWelcome.")
}
