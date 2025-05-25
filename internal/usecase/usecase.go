package usecase

import (
	"errors"
	"fmt"
	"sync"

	"parkinglot/internal/entity"
)

type UseCases struct {
	lot             *entity.ParkingLot
	repo            entity.ParkingRepository
	vehicleRegistry sync.Map
}

func NewUseCases(lot *entity.ParkingLot, repo entity.ParkingRepository) *UseCases {
	return &UseCases{lot: lot, repo: repo}
}

func (uc *UseCases) Park(vehicleType entity.VehicleType, vehicleID string) (string, error) {
	vehicle := entity.Vehicle{ID: vehicleID, Type: vehicleType}

	for spotID, spot := range uc.lot.Spots {
		if spot.VehicleType == vehicleType && spot.Active {
			spot.Lock()
			if !spot.Occupied {
				spot.Occupied = true
				spot.VehicleID = vehicleID
				spot.Unlock()

				// Save to DB if needed
				_ = uc.repo.SaveVehicle(vehicle, spot)

				// Track last known location
				uc.vehicleRegistry.Store(vehicleID, spotID)

				return spotID, nil
			}
			spot.Unlock()
		}
	}
	return "", errors.New("no available spot")
}

func (uc *UseCases) Unpark(spotID, vehicleID string) error {
	spot, ok := uc.lot.Spots[spotID]
	if !ok {
		return fmt.Errorf("spot %s not found", spotID)
	}

	spot.Lock()
	defer spot.Unlock()

	if !spot.Occupied || spot.VehicleID != vehicleID {
		return errors.New("vehicle not found in this spot")
	}

	// Mark as free
	spot.Occupied = false
	spot.VehicleID = ""

	// Update DB if needed
	_ = uc.repo.RemoveVehicle(vehicleID)

	// Track last known spot
	uc.vehicleRegistry.Store(vehicleID, spotID)

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

func (uc *UseCases) SearchVehicle(vehicleID string) {
	if val, ok := uc.vehicleRegistry.Load(vehicleID); ok {
		fmt.Printf("Vehicle %s was last seen at spot: %s\n", vehicleID, val.(string))
	} else {
		fmt.Printf("Vehicle %s not found.\n", vehicleID)
	}
}

func (uc *UseCases) ShowAvailable(vehicleType entity.VehicleType) {
	fmt.Printf("Available spots for %s:\n", vehicleType.String())
	for id, spot := range uc.lot.Spots {
		if spot.VehicleType == vehicleType && spot.Active && !spot.Occupied {
			fmt.Println("-", id)
		}
	}
}
