package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"parkinglot/internal/adapter/repository"
	"parkinglot/internal/entity"
	"parkinglot/internal/usecase"
)

func Run() {
	db, err := repository.InitDB("parkinglot.db")
	if err != nil {
		fmt.Println("Failed to initialize database:", err)
		return
	}

	repo := repository.NewSQLiteRepo(db)
	config := map[string]string{
		"0-0-0": "B-1",
		"0-0-1": "M-1",
		"1-0-0": "A-1",
		"1-0-1": "M-1",
		"2-0-0": "A-1",
		"2-0-1": "M-1",
		"0-1-1": "X-0",
	}

	lot := entity.NewParkingLotFromConfig(config)
	usecases := usecase.NewUseCases(lot, repo)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("vehicle types: bicycle, motorcycle, automobile")
	fmt.Println("CLI started. Commands:")
	fmt.Println("park <vehicle_type> <vehicle_id>")
	fmt.Println("unpark <spot_id> <vehicle_id>")
	fmt.Println("availableSpot <vehicle_type>")
	fmt.Println("searchVehicle <vehicle_id>")
	fmt.Println("status")
	fmt.Println("quit")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "park":
			if len(args) != 3 {
				fmt.Println("Usage: park <vehicle_type> <vehicle_id>")
				continue
			}
			vehicleType, err := entity.ParseVehicleType(args[1])
			if err != nil {
				fmt.Println("Invalid vehicle type.")
				continue
			}
			spotID, err := usecases.Park(vehicleType, args[2]) // no floor
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Parked at spot:", spotID)
			}

		case "unpark":
			if len(args) != 3 {
				fmt.Println("Usage: unpark <spot_id> <vehicle_id>")
				continue
			}
			err := usecases.Unpark(args[1], args[2])
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Unparked successfully.")
			}
		case "availableSpot":
			if len(args) != 2 {
				fmt.Println("Usage: availableSpot <vehicle_type>")
				continue
			}
			vehicleType, err := entity.ParseVehicleType(args[1])
			if err != nil {
				fmt.Println("Invalid vehicle type.")
				continue
			}
			usecases.ShowAvailable(vehicleType)
		case "searchVehicle":
			if len(args) != 2 {
				fmt.Println("Usage: searchVehicle <vehicle_id>")
				continue
			}
			usecases.SearchVehicle(args[1])
		case "status":
			usecases.ShowStatus()
		case "quit":
			fmt.Println("Exiting.")
			return

		default:
			fmt.Println("Unknown command.")
		}
	}
}
