# 🅿️ Parking Lot CLI (Clean Architecture + SQLite)

A thread-safe, multi-floor parking lot system implemented in Go using Clean Architecture principles, SQLite for persistence, and mutexes for concurrency control.

---

## 🚀 Features

- Multiple floors, rows, and columns of parking spots
- Supports spot types:
  - `B-1` → Active Bicycle spot
  - `M-1` → Active Motorcycle spot
  - `A-1` → Active Automobile spot
  - `X-0` → Inactive spot
- Vehicle types: 🚲 Bicycle, 🏍️ Motorcycle, 🚗 Automobile
- Spot ID format: `floor-row-col` (e.g., `1-0-2`)
- Thread-safe parking via `sync.Mutex`
- SQLite-backed persistence of parked vehicles
- `sync.Map` used to store last known spot of each vehicle
- CLI interface simulating multiple concurrent gates

---

## 🛠 Requirements

- Go 1.20+
- SQLite3 (CLI optional)
- Git

---

## 📦 Installation

Clone the repo and install dependencies:

```bash
git clone https://github.com/your-username/parkinglot
cd parkinglot
go mod tidy
```

## 💻 CLI Commands

These are the available commands you can run after starting the CLI:

| Command                                   | Description                                                            |
|-------------------------------------------|------------------------------------------------------------------------|
| `park <vehicle_type> <vehicle_id>`        | Park a vehicle of a given type in the first available spot             |
|                                           | _Example_: `park bicycle B1`                                           |
| `unpark <spot_id> <vehicle_id>`           | Free a parking spot by specifying the spot ID and the matching vehicle |
|                                           | _Example_: `unpark 1-0-0 B1`                                           |
| `availableSpot <vehicle_type>`            | Show only **free and active** spots for the specified vehicle type     |
|                                           | _Example_: `availableSpot motorcycle`                                  |
| `searchVehicle <vehicle_id>`              | Display the **last known spot** of the vehicle, even if already exited |
|                                           | _Example_: `searchVehicle B1`                                          |
| `status`                                  | Show **all parking spots**, including occupied and inactive            |
| `quit`                                    | Exit the CLI application                                               |


### 🧠 Notes

- `vehicle_type` must be one of:
  - `bicycle`
  - `motorcycle`
  - `automobile`
- `vehicle_id` should be unique per vehicle.
- `spot_id` must follow the format `floor-row-col`, such as `2-0-1`.