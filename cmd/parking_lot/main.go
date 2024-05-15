package main

import (
    "bufio"
    "fmt"
    "os"
    "parking/internal/app"
    "parking/internal/infra"
    "strconv"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: parking_lot <input_file>")
        return
    }

    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    repo := &infra.ParkingLotRepository{}
    service := app.NewParkingLotService(repo)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        command := scanner.Text()
        parts := strings.Fields(command)
        switch parts[0] {
        case "create_parking_lot":
            capacity, _ := strconv.Atoi(parts[1])
            service.CreateParkingLot(capacity)
            fmt.Printf("Created parking lot with %d slots\n", capacity)
        case "park":
            slot, err := service.ParkCar(parts[1], parts[2])
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Printf("Allocated slot number: %d\n", slot)
            }
        case "leave":
            slotNumber, _ := strconv.Atoi(parts[1])
            hours, _ := strconv.Atoi(parts[2])
            car, charge, err := service.LeaveCar(slotNumber, hours)
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Printf("Registration number %s with Slot Number %d is free with Charge %d\n", car.RegistrationNumber, slotNumber, charge)
            }
        case "status":
            slots := service.GetStatus()
            fmt.Println("Slot No. Registration No.")
            for _, slot := range slots {
                if slot.Car != nil {
                    fmt.Printf("%d %s\n", slot.Number, slot.Car.RegistrationNumber)
                }
            }
        default:
            fmt.Println("Unknown command:", command)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}
