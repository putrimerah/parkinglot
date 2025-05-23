package usecase

import (
	"strconv"
	"sync"
	"testing"

	"parkinglot/internal/adapter/repository"
	"parkinglot/internal/entity"
)

func TestEnterVehicleRaceCondition(t *testing.T) {
	// Setup in-memory test repo with 2 Bicycle spots
	memRepo := repository.NewTestRepo()
	memRepo.SeedSpots([]*entity.ParkingSpot{
		{Row: 0, Col: 0, VehicleType: entity.Bicycle, Active: true, Occupied: false},
		{Row: 0, Col: 1, VehicleType: entity.Bicycle, Active: true, Occupied: false},
	})

	lot := entity.NewParkingLotFromRepo(memRepo.Spots)
	uc := NewUseCases(lot, memRepo)

	var wg sync.WaitGroup
	successCount := 0
	mu := sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			vehicleID := "B" + strconv.Itoa(id)
			err := uc.EnterVehicle(entity.Bicycle, vehicleID)
			if err == nil {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()

	if successCount != 2 {
		t.Errorf("Expected 2 vehicles to park successfully, got %d", successCount)
	}
}
