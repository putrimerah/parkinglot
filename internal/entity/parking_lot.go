package entity

type ParkingLot struct {
	Spots [][]*ParkingSpot
}

// NewParkingLotFromRepo constructs a ParkingLot from persisted spots
func NewParkingLotFromRepo(spots []*ParkingSpot) *ParkingLot {
	rowMap := map[int][]*ParkingSpot{}
	for _, s := range spots {
		rowMap[s.Row] = append(rowMap[s.Row], s)
	}

	lot := &ParkingLot{}
	for i := 0; i < len(rowMap); i++ {
		lot.Spots = append(lot.Spots, rowMap[i])
	}
	return lot
}

// FindAvailableSpot returns the first unoccupied, active spot for the given vehicle type
func (lot *ParkingLot) FindAvailableSpot(vehicleType VehicleType) *ParkingSpot {
	for _, row := range lot.Spots {
		for _, spot := range row {
			spot.Lock()
			if spot.Active && !spot.Occupied && spot.VehicleType == vehicleType {
				spot.Occupied = true
				spot.Unlock()
				return spot
			}
			spot.Unlock()
		}
	}
	return nil
}

// GetSpot safely retrieves a spot from the lot
func (lot *ParkingLot) GetSpot(row, col int) *ParkingSpot {
	if row >= len(lot.Spots) || col >= len(lot.Spots[row]) {
		return nil
	}
	return lot.Spots[row][col]
}
