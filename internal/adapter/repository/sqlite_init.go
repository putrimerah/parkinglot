package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the SQLite database and ensures the schema and seed data are set.
func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	schema := `
	CREATE TABLE IF NOT EXISTS vehicle_spots (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		vehicle_id TEXT,
		row INTEGER,
		col INTEGER,
		vehicle_type INTEGER,
		occupied BOOLEAN,
		active BOOLEAN
	);`
	if _, err = db.Exec(schema); err != nil {
		return nil, fmt.Errorf("failed to create schema: %w", err)
	}

	if err := seedInitialSpots(db); err != nil {
		return nil, fmt.Errorf("failed to seed spots: %w", err)
	}

	return db, nil
}

// seedInitialSpots inserts initial free and active spots if the table is empty.
func seedInitialSpots(db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM vehicle_spots").Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil // Already seeded
	}

	fmt.Println("Seeding initial parking spots...")

	stmt := `
	INSERT INTO vehicle_spots (
		vehicle_id, row, col, vehicle_type, occupied, active
	) VALUES (?, ?, ?, ?, ?, ?)`

	spots := []struct {
		row, col, vehicleType int
	}{
		{0, 0, 0}, {0, 1, 0}, // Bicycle (0)
		{1, 0, 1}, {1, 1, 1}, // Motorcycle (1)
		{2, 0, 2}, {2, 1, 2}, // Automobile (2)
	}

	for _, s := range spots {
		if _, err := db.Exec(stmt, nil, s.row, s.col, s.vehicleType, false, true); err != nil {
			return err
		}
	}

	return nil
}
