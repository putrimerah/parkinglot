# 🅿️ Parking Lot CLI (Clean Architecture + SQLite)

A thread-safe, multi-floor parking lot system implemented in Go using Clean Architecture principles, SQLite for persistence, and mutexes for concurrency control.

---

## 🚀 Features

- Multiple vehicle types: 🚲 Bicycle, 🏍️ Motorcycle, 🚗 Automobile
- Thread-safe vehicle entry/exit using `sync.Mutex`
- SQLite-backed persistence of parked vehicles
- Auto-seeded parking spots on first run
- CLI interface to simulate multiple gates

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

| Command                          | Description                                                  |
|----------------------------------|--------------------------------------------------------------|
| `enter <type> <vehicle_id>`     | Park a vehicle of a given type and ID                        |
|                                  | _Example_: `enter bicycle B1`                                |
| `exit <vehicle_id>`             | Free a parked vehicle’s spot by its vehicle ID               |
|                                  | _Example_: `exit B1`                                         |
| `status`                        | Show **all parking spots**, including occupied and inactive  |
| `available`                     | Show only **free and active** spots across all vehicle types |
| `quit`                          | Exit the CLI application                                     |

### 🧠 Notes

- `vehicle_type` must be one of:
  - `bicycle`
  - `motorcycle`
  - `automobile`
- `vehicle_id` should be unique per vehicle.