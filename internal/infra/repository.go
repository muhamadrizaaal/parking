package infra

import (
	"fmt"
	"parking/internal/domain"
)

type ParkingLotRepository struct {
    ParkingLot *domain.ParkingLot
}

func (repo *ParkingLotRepository) CreateParkingLot(capacity int) {
    slots := make([]*domain.Slot, capacity)
    for i := 0; i < capacity; i++ {
        slots[i] = &domain.Slot{Number: i + 1}
    }
    repo.ParkingLot = &domain.ParkingLot{Capacity: capacity, Slots: slots}
}

func (repo *ParkingLotRepository) ParkCar(car *domain.Car) (int, error) {
    for _, slot := range repo.ParkingLot.Slots {
        if slot.Car == nil {
            slot.Car = car
            return slot.Number, nil
        }
    }
    return 0, fmt.Errorf("parking lot is full")
}

func (repo *ParkingLotRepository) LeaveCar(slotNumber int, hours int) (*domain.Car, int, error) {
    if slotNumber <= 0 || slotNumber > repo.ParkingLot.Capacity {
        return nil, 0, fmt.Errorf("slot number out of range")
    }
    slot := repo.ParkingLot.Slots[slotNumber-1]
    if slot.Car == nil {
        return nil, 0, fmt.Errorf("no car in slot")
    }
    car := slot.Car
    slot.Car = nil
    charge := 10
    if hours > 2 {
        charge += (hours - 2) * 10
    }
    return car, charge, nil
}

func (repo *ParkingLotRepository) GetStatus() []*domain.Slot {
    return repo.ParkingLot.Slots
}
