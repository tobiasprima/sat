package main

import (
	"SAT/config"
	"fmt"
)

func main() {
	db := config.ConnectDatabase()
	defer db.Close()

	fmt.Print("Input user name:")
	var username string
	fmt.Scan(&username)

	fmt.Print("Input password:")
	var password string
	fmt.Scan(&password)

}
