package handlers_test

import (
	"SAT/handlers"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddItem(t *testing.T) {
	// Initialize mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Mock the expected query and result
	mock.ExpectExec("INSERT INTO items").
		WithArgs("device", 100000.0, 100).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call AddItem with test inputs
	handlers.AddItem(db, "device", 100000.0, 100)

	// Assert expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}

func TestDeleteItem(t *testing.T) {
	// Initialize mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %v", err)
	}
	defer db.Close()

	// Mock the expected query and result
	mock.ExpectExec("DELETE FROM items").
		WithArgs("device").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call DeleteItem with test input
	handlers.DeleteItem(db, "device")

	// Assert expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}
