package entity

type ParkingRepository interface {
	SaveVehicle(vehicle Vehicle, spot *ParkingSpot) error
	RemoveVehicle(vehicleID string) error
	GetVehicleSpot(vehicleID string) (*ParkingSpot, error)
	LoadAllSpots() ([]*ParkingSpot, error)
}
