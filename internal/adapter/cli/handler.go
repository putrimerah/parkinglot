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
	spots, err := repo.LoadAllSpots()
	if err != nil {
		fmt.Println("Failed to load spots:", err)
		return
	}

	lot := entity.NewParkingLotFromRepo(spots)
	usecases := usecase.NewUseCases(lot, repo)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("vehicle types: bicycle, motorcycle, automobile")
	fmt.Println("CLI started. Commands:")
	fmt.Println("enter <vehicle_type> <vehicle_id>")
	fmt.Println("exit <vehicle_id>")
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
		case "enter":
			if len(args) != 3 {
				fmt.Println("Usage: enter <vehicle_type> <vehicle_id>")
				continue
			}
			vehicleType, err := entity.ParseVehicleType(args[1])
			if err != nil {
				fmt.Println("Invalid vehicle type.")
				continue
			}
			err = usecases.EnterVehicle(vehicleType, args[2])
			if err != nil {
				fmt.Println("Error:", err)
			}

		case "exit":
			if len(args) != 2 {
				fmt.Println("Usage: exit <vehicle_id>")
				continue
			}
			err := usecases.ExitVehicle(args[1])
			if err != nil {
				fmt.Println("Error:", err)
			}

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
