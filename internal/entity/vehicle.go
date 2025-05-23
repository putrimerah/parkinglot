package entity

import "errors"

type VehicleType int

const (
	Bicycle VehicleType = iota
	Motorcycle
	Automobile
)

type Vehicle struct {
	ID   string
	Type VehicleType
}

func ParseVehicleType(s string) (VehicleType, error) {
	switch s {
	case "bicycle":
		return Bicycle, nil
	case "motorcycle":
		return Motorcycle, nil
	case "automobile":
		return Automobile, nil
	default:
		return -1, errors.New("unknown vehicle type")
	}
}

func (vt VehicleType) String() string {
	switch vt {
	case Bicycle:
		return "Bicycle"
	case Motorcycle:
		return "Motorcycle"
	case Automobile:
		return "Automobile"
	default:
		return "Unknown"
	}
}
