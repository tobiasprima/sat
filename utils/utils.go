package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// CapitalizeName ensures the first letter of each word is capitalized
func CapitalizeName(name string) string {
	words := strings.Fields(name)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

// ValidateEmail checks if the input is a valid email format
func ValidateEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// PrintTable prints data in a tabular format with headers and rows
func PrintTable(headers []string, rows [][]string) {
	// Calculate column widths based on headers and rows
	columnWidths := make([]int, len(headers)) // Create a slice to store the maximum width of each column
	for i, header := range headers {
		columnWidths[i] = len(header) // Initially set the width to the length of the header
	}
	for _, row := range rows { // Iterate through all rows to check if any cell exceeds the current width
		for i, cell := range row {
			if len(cell) > columnWidths[i] {
				columnWidths[i] = len(cell) // Update the width if the cell content is longer
			}
		}
	}

	// Print headers with calculated widths
	for i, header := range headers {
		fmt.Printf("%-*s", columnWidths[i]+2, header) // %-*s aligns text to the left, with columnWidths[i]+2 for spacing
	}
	fmt.Println() // Print a newline after headers

	// Print a horizontal separator line
	fmt.Println(strings.Repeat("-", sumColumnWidths(columnWidths)+len(columnWidths)*2)) // Separator length includes column widths and spaces

	// Print each row in the table
	for _, row := range rows {
		for i, cell := range row {
			fmt.Printf("%-*s", columnWidths[i]+2, cell) // Print each cell with proper alignment
		}
		fmt.Println() // Print a newline after each row
	}
}

// sumColumnWidths calculates the sum of the column widths for the table
func sumColumnWidths(nums []int) int {
	total := 0 // Initialize the total sum to 0
	for _, num := range nums {
		total += num // Add each number in the slice to the total
	}
	return total // Return the sum of the numbers
}

// EnterToContinue asks user to press enter to continue the program
func EnterTocontinue() {
	var dummy string
	fmt.Print("\nPress enter to go to main menu...")
	fmt.Scanln(&dummy)
}

// PrintTableTitle shows table title
func PrintTableTitle(title, startDate, endDate string) {
	fmt.Printf("\n%s: ", title)
	if startDate == endDate {
		fmt.Printf("%s\n\n", startDate)
	} else {
		fmt.Printf("%s to %s\n\n", startDate, endDate)
	}

}
