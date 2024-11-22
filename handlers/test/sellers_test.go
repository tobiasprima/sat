package handlers_test

import (
	"SAT/handlers"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddSeller(t *testing.T) {
	// Initialize mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Mock the expected query and result
	mock.ExpectExec("INSERT INTO Sellers").
		WithArgs("John Doe", "john.doe@example.com").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call AddSeller with test inputs
	handlers.AddSeller(db, "John Doe", "john.doe@example.com")

	// Assert expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}

func TestDeleteSeller(t *testing.T) {
	// Initialize mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Mock the expected query and result
	mock.ExpectExec("DELETE FROM Sellers").
		WithArgs("john.doe@example.com").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call DeleteSeller with test input
	handlers.DeleteSeller(db, "john.doe@example.com")

	// Assert expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}
