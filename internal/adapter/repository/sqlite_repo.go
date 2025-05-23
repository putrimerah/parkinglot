package repository

import (
	"database/sql"
	"errors"
	"parkinglot/internal/entity"
)

type SQLiteRepo struct {
	db *sql.DB
}

func NewSQLiteRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{db: db}
}

func (r *SQLiteRepo) SaveVehicle(vehicle entity.Vehicle, spot *entity.ParkingSpot) error {
	_, err := r.db.Exec(
		`INSERT OR REPLACE INTO vehicle_spots 
		 (vehicle_id, row, col, vehicle_type, occupied, active)
		 VALUES (?, ?, ?, ?, ?, ?)`,
		vehicle.ID, spot.Row, spot.Col, int(spot.VehicleType), true, spot.Active,
	)
	return err
}

func (r *SQLiteRepo) RemoveVehicle(vehicleID string) error {
	result, err := r.db.Exec(`DELETE FROM vehicle_spots WHERE vehicle_id = ?`, vehicleID)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("vehicle not found")
	}
	return nil
}

func (r *SQLiteRepo) GetVehicleSpot(vehicleID string) (*entity.ParkingSpot, error) {
	row := r.db.QueryRow(`
		SELECT row, col, vehicle_type, occupied, active
		FROM vehicle_spots
		WHERE vehicle_id = ?`, vehicleID)

	var spot entity.ParkingSpot
	var vt int
	err := row.Scan(&spot.Row, &spot.Col, &vt, &spot.Occupied, &spot.Active)
	if err != nil {
		return nil, err
	}
	spot.VehicleType = entity.VehicleType(vt)
	return &spot, nil
}

func (r *SQLiteRepo) LoadAllSpots() ([]*entity.ParkingSpot, error) {
	rows, err := r.db.Query(`SELECT row, col, vehicle_type, occupied, active FROM vehicle_spots`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spots []*entity.ParkingSpot
	for rows.Next() {
		var spot entity.ParkingSpot
		var vt int
		if err := rows.Scan(&spot.Row, &spot.Col, &vt, &spot.Occupied, &spot.Active); err != nil {
			return nil, err
		}
		spot.VehicleType = entity.VehicleType(vt)
		spots = append(spots, &spot)
	}
	return spots, nil
}
