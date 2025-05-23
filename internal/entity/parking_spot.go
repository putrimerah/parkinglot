package entity

import "sync"

type ParkingSpot struct {
	Row         int
	Col         int
	VehicleType VehicleType
	Occupied    bool
	Active      bool
	mu          sync.Mutex
}

func (s *ParkingSpot) Lock() {
	s.mu.Lock()
}

func (s *ParkingSpot) Unlock() {
	s.mu.Unlock()
}
