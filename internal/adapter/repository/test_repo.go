package repository

import "parkinglot/internal/entity"

type TestRepo struct {
	Spots []*entity.ParkingSpot
}

func NewTestRepo() *TestRepo {
	return &TestRepo{}
}

func (r *TestRepo) SeedSpots(spots []*entity.ParkingSpot) {
	r.Spots = spots
}

func (r *TestRepo) SaveVehicle(v entity.Vehicle, s *entity.ParkingSpot) error {
	return nil
}

func (r *TestRepo) RemoveVehicle(string) error { return nil }

func (r *TestRepo) GetVehicleSpot(string) (*entity.ParkingSpot, error) { return nil, nil }

func (r *TestRepo) LoadAllSpots() ([]*entity.ParkingSpot, error) {
	return r.Spots, nil
}
