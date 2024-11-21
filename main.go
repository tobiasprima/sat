package main

import (
	"SAT/config"
	"fmt"
)

func main() {
	db := config.ConnectDatabase()
	defer db.Close()

	fmt.Println("Welcome to the SAT system!")

	for {
		fmt.Println("Please log in to continue.")
		// Login handler disini
		// if handlers.LoginUser(db){
		// 	break
		// }
		fmt.Println("Login failed. Please try again.\n")
	}
}
