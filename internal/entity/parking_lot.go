package entity

type ParkingLot struct {
	Spots map[string]*ParkingSpot // key: "floor-row-col"
}

func NewParkingLotFromConfig(config map[string]string) *ParkingLot {
	spots := make(map[string]*ParkingSpot)

	for id, spotType := range config {
		floor, row, col, err := ParseSpotID(id)
		if err != nil {
			continue // or log error
		}

		vType, active := ParseSpotType(spotType)
		if vType == -1 {
			continue // invalid type
		}

		spots[id] = &ParkingSpot{
			Floor:       floor,
			Row:         row,
			Col:         col,
			VehicleType: vType,
			Active:      active,
			Occupied:    false,
		}
	}

	return &ParkingLot{Spots: spots}
}

// // FindAvailableSpot returns the first unoccupied, active spot for the given vehicle type
// func (lot *ParkingLot) FindAvailableSpot(vehicleType VehicleType) *ParkingSpot {
// 	for _, row := range lot.Spots {
// 		for _, spot := range row {
// 			spot.Lock()
// 			if spot.Active && !spot.Occupied && spot.VehicleType == vehicleType {
// 				spot.Occupied = true
// 				spot.Unlock()
// 				return spot
// 			}
// 			spot.Unlock()
// 		}
// 	}
// 	return nil
// }

// GetSpot safely retrieves a spot from the lot
// func (lot *ParkingLot) GetSpot(row, col int) *ParkingSpot {
// 	if row >= len(lot.Spots) || col >= len(lot.Spots[row]) {
// 		return nil
// 	}
// 	return lot.Spots[row][col]
// }
