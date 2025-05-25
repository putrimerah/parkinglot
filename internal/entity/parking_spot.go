package entity

import (
	"fmt"
	"sync"
)

type ParkingSpot struct {
	Floor       int
	Row         int
	Col         int
	VehicleType VehicleType
	Occupied    bool
	Active      bool
	VehicleID   string
	mu          sync.Mutex
}

func (s *ParkingSpot) Lock()   { s.mu.Lock() }
func (s *ParkingSpot) Unlock() { s.mu.Unlock() }

func ParseSpotType(code string) (VehicleType, bool) {
	switch code {
	case "B-1":
		return Bicycle, true
	case "M-1":
		return Motorcycle, true
	case "A-1":
		return Automobile, true
	case "X-0":
		return -1, false
	default:
		return -1, false
	}
}

func FormatSpotID(floor, row, col int) string {
	return fmt.Sprintf("%d-%d-%d", floor, row, col)
}

func ParseSpotID(spotID string) (int, int, int, error) {
	var floor, row, col int
	_, err := fmt.Sscanf(spotID, "%d-%d-%d", &floor, &row, &col)
	return floor, row, col, err
}
