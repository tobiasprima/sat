package handlers

import (
	"SAT/utils"
	"database/sql"
	"fmt"
	"time"
)

// formatDateRange calculates the start and end dates for the selected time period.
// The input is a period string ("daily", "weekly", "monthly"), and it returns the start and end date strings.
func formatDateRange(period string) (string, string, error) {
	now := time.Now()
	switch period {
	case "daily":
		// Use the current date for both start and end
		return now.Format("2006-01-02"), now.Format("2006-01-02"), nil
	case "weekly":
		// Calculate the start of the week (Sunday)
		startOfWeek := now.AddDate(0, 0, -int(now.Weekday()))
		return startOfWeek.Format("2006-01-02"), now.Format("2006-01-02"), nil
	case "monthly":
		// Calculate the first day of the current month
		startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
		return startOfMonth.Format("2006-01-02"), now.Format("2006-01-02"), nil
	default:
		// Return an error if the input period is invalid
		return "", "", fmt.Errorf("invalid time period")
	}
}

// promptForPeriod prompts the user to select a time period (daily, weekly, monthly).
// Returns the chosen period as a string.
func promptForPeriod() string {
	for {
		fmt.Println("Choose a time period:")
		fmt.Println("1. Daily")
		fmt.Println("2. Weekly")
		fmt.Println("3. Monthly")
		fmt.Print("Your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			return "daily"
		case 2:
			return "weekly"
		case 3:
			return "monthly"
		default:
			// Invalid choice loops the prompt again
			fmt.Print("Invalid choice. Please try again.\n\n")
		}
	}
}

// TotalSalesReport generates a report of total sales and items sold for the selected time period.
func TotalSalesReport(db *sql.DB) {
	// Prompt user to select a reporting period
	period := promptForPeriod()
	startDate, endDate, err := formatDateRange(period)
	if err != nil {
		fmt.Println("Error determining date range:", err)
		return
	}

	// Query to fetch total quantity sold and total sales amount per item
	rows, err := db.Query(`
		SELECT i.name, SUM(od.quantity) AS total_quantity, SUM(od.price) AS total_sales
		FROM items i
		JOIN order_details od ON i.item_id = od.item_id
		JOIN orders o ON od.order_id = o.order_id
		WHERE o.date BETWEEN $1 AND $2
		GROUP BY i.name
		ORDER BY total_sales DESC
	`, startDate, endDate)
	if err != nil {
		fmt.Println("Error fetching total sales report:", err)
		return
	}
	defer rows.Close()

	// Initialize overall totals
	var overallTotalSales float64
	var overallTotalItems int
	data := [][]string{}

	// Process each row from the query result
	for rows.Next() {
		var name string
		var totalQuantity int
		var totalSales float64
		if err := rows.Scan(&name, &totalQuantity, &totalSales); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		// Append data for display in a table
		data = append(data, []string{name, fmt.Sprintf("%d", totalQuantity), fmt.Sprintf("%.2f", totalSales)})
		overallTotalSales += totalSales
		overallTotalItems += totalQuantity
	}

	// Check if there are no sales during the selected period
	if len(data) == 0 {
		fmt.Printf("No sales made during the %s period.\n", period)
		return
	}

	// Display the data in a table
	utils.PrintTableTitle("Total Sales Report", startDate, endDate)
	utils.PrintTable([]string{"Item Name", "Total Quantity Sold", "Total Sales (Rp.)"}, data)
	fmt.Printf("\nOverall Total Sales: Rp. %.2f\n", overallTotalSales)
	fmt.Printf("Overall Total Items Sold: %d items\n", overallTotalItems)
	utils.EnterTocontinue()
}

// PopularItemsReport generates a report of popular items for the selected time period.
func PopularItemsReport(db *sql.DB) {
	// Prompt user to select a reporting period
	period := promptForPeriod()
	startDate, endDate, err := formatDateRange(period)
	if err != nil {
		fmt.Println("Error determining date range:", err)
		return
	}

	// Query to fetch item popularity based on the count of items sold
	rows, err := db.Query(`
		SELECT i.name, COUNT(od.item_id) AS popularity
		FROM items i
		JOIN order_details od ON i.item_id = od.item_id
		JOIN orders o ON od.order_id = o.order_id
		WHERE o.date BETWEEN $1 AND $2
		GROUP BY i.name
		ORDER BY popularity DESC
	`, startDate, endDate)
	if err != nil {
		fmt.Println("Error fetching popular items report:", err)
		return
	}
	defer rows.Close()

	data := [][]string{}

	// Process each row from the query result
	for rows.Next() {
		var name string
		var timesSold int
		if err := rows.Scan(&name, &timesSold); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		// Append data for display in a table
		data = append(data, []string{name, fmt.Sprintf("%d", timesSold)})
	}

	// Check if no items were sold during the selected period
	if len(data) == 0 {
		fmt.Printf("No items sold during the %s period.\n", period)
		return
	}

	// Display the data in a table
	utils.PrintTableTitle("Popular Items Report", startDate, endDate)
	utils.PrintTable([]string{"Item Name", "Times Sold"}, data)
	utils.EnterTocontinue()
}

// SellerRankingReport generates a report of sellers ranked by the number of items sold during the selected time period.
func SellerRankingReport(db *sql.DB) {
	// Prompt user to select a reporting period
	period := promptForPeriod()
	startDate, endDate, err := formatDateRange(period)
	if err != nil {
		fmt.Println("Error determining date range:", err)
		return
	}

	// Query to fetch seller rankings based on the number of items sold
	rows, err := db.Query(`
		SELECT s.name, COUNT(od.order_id) AS items_sold
		FROM sellers s
		JOIN orders o ON s.seller_id = o.seller_id
		JOIN order_details od ON o.order_id = od.order_id
		WHERE o.date BETWEEN $1 AND $2
		GROUP BY s.name
		ORDER BY items_sold DESC
	`, startDate, endDate)
	if err != nil {
		fmt.Println("Error fetching seller ranking report:", err)
		return
	}
	defer rows.Close()

	data := [][]string{}

	// Process each row from the query result
	for rows.Next() {
		var name string
		var itemsSold int
		if err := rows.Scan(&name, &itemsSold); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		// Append data for display in a table
		data = append(data, []string{name, fmt.Sprintf("%d", itemsSold)})
	}

	// Check if no sellers made sales during the selected period
	if len(data) == 0 {
		fmt.Printf("No seller activity during the %s period.\n", period)
		return
	}

	// Display the data in a table
	utils.PrintTableTitle("Seller Ranking Report", startDate, endDate)
	utils.PrintTable([]string{"Seller Name", "Items Sold"}, data)
	utils.EnterTocontinue()
}
