package handlers_test

import (
	"SAT/handlers"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestTotalSalesReport(t *testing.T) {
	// Initialize mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Mock the expected query and result
	mock.ExpectQuery("SELECT i.name, SUM\\(od.quantity\\) AS total_quantity, SUM\\(od.price\\) AS total_sales").
		WithArgs("2024-11-01", "2024-11-07"). // Example startDate and endDate for weekly report
		WillReturnRows(sqlmock.NewRows([]string{"name", "total_quantity", "total_sales"}).
			AddRow("Item A", 10, 100.0).
			AddRow("Item B", 5, 50.0))

	// Simulate user choice for period
	handlers.TotalSalesReport(db)

	// Assert expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}

func TestPopularItemsReport(t *testing.T) {
	// Initialize mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Mock the expected query and result
	mock.ExpectQuery("SELECT i.name, COUNT\\(od.item_id\\) AS popularity").
		WithArgs("2024-11-01", "2024-11-07"). // Example startDate and endDate for weekly report
		WillReturnRows(sqlmock.NewRows([]string{"name", "popularity"}).
			AddRow("Item A", 15).
			AddRow("Item B", 10))

	// Simulate user choice for period
	handlers.PopularItemsReport(db)

	// Assert expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}

func TestSellerRankingReport(t *testing.T) {
	// Initialize mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Mock the expected query and result
	mock.ExpectQuery("SELECT s.name, COUNT\\(od.order_id\\) AS items_sold").
		WithArgs("2024-11-01", "2024-11-07"). // Example startDate and endDate for weekly report
		WillReturnRows(sqlmock.NewRows([]string{"name", "items_sold"}).
			AddRow("Seller A", 20).
			AddRow("Seller B", 15))

	// Simulate user choice for period
	handlers.SellerRankingReport(db)

	// Assert expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}
