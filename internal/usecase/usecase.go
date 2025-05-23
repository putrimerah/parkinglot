package usecase

import (
	"errors"
	"fmt"

	"parkinglot/internal/entity"
)

type UseCases struct {
	lot  *entity.ParkingLot
	repo entity.ParkingRepository
}

func NewUseCases(lot *entity.ParkingLot, repo entity.ParkingRepository) *UseCases {
	return &UseCases{lot: lot, repo: repo}
}

// EnterVehicle assigns an available spot and saves it to the repo
func (uc *UseCases) EnterVehicle(vehicleType entity.VehicleType, vehicleID string) error {
	vehicle := entity.Vehicle{ID: vehicleID, Type: vehicleType}

	spot := uc.lot.FindAvailableSpot(vehicleType)
	if spot == nil {
		return errors.New("no available spot")
	}

	// No need to check if already occupied — FindAvailableSpot() guaranteed it
	err := uc.repo.SaveVehicle(vehicle, spot)
	if err != nil {
		// rollback in memory
		spot.Lock()
		spot.Occupied = false
		spot.Unlock()
		return fmt.Errorf("failed to persist vehicle: %w", err)
	}

	fmt.Printf("Vehicle %s parked at [%d,%d]\n", vehicle.ID, spot.Row, spot.Col)
	return nil
}

// ExitVehicle releases a spot and removes the vehicle from the repo
func (uc *UseCases) ExitVehicle(vehicleID string) error {
	spot, err := uc.repo.GetVehicleSpot(vehicleID)
	if err != nil {
		return fmt.Errorf("could not find vehicle: %w", err)
	}

	memSpot := uc.lot.GetSpot(spot.Row, spot.Col)
	if memSpot == nil {
		return errors.New("spot not found in memory")
	}

	memSpot.Lock()
	defer memSpot.Unlock()

	memSpot.Occupied = false
	err = uc.repo.RemoveVehicle(vehicleID)
	if err != nil {
		return fmt.Errorf("failed to remove vehicle: %w", err)
	}

	fmt.Printf("Vehicle %s exited. Spot [%d,%d] is now free.\n", vehicleID, spot.Row, spot.Col)
	return nil
}

// ShowStatus lists all parking spots from the repo
func (uc *UseCases) ShowStatus() {
	spots, err := uc.repo.LoadAllSpots()
	if err != nil {
		fmt.Println("Failed to load spots:", err)
		return
	}

	fmt.Println("Current Parking Lot Status:")
	for _, spot := range spots {
		status := "Free"
		if !spot.Active {
			status = "Inactive"
		} else if spot.Occupied {
			status = "Occupied"
		}
		fmt.Printf("Spot [%d,%d] - %s (%s)\n", spot.Row, spot.Col, status, spot.VehicleType.String())
	}
}
