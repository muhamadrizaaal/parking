package test

import (
    // "parking/internal/domain"
    "parking/internal/infra"
    "parking/internal/usecases"
    "testing"
)

func TestCreateParkingLot(t *testing.T) {
    repo := &infra.ParkingLotRepository{}
    usecase := usecases.NewCreateParkingLotUseCase(repo)

    usecase.Execute(6)
    if repo.ParkingLot.Capacity != 6 {
        t.Errorf("Expected capacity 6, got %d", repo.ParkingLot.Capacity)
    }
}

func TestParkCar(t *testing.T) {
    repo := &infra.ParkingLotRepository{}
    createUseCase := usecases.NewCreateParkingLotUseCase(repo)
    createUseCase.Execute(6)

    parkUseCase := usecases.NewParkCarUseCase(repo)
    slot, err := parkUseCase.Execute("KA-01-HH-1234", "White")
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if slot != 1 {
        t.Errorf("Expected slot 1, got %d", slot)
    }
}

func TestLeaveCar(t *testing.T) {
    repo := &infra.ParkingLotRepository{}
    createUseCase := usecases.NewCreateParkingLotUseCase(repo)
    createUseCase.Execute(6)

    parkUseCase := usecases.NewParkCarUseCase(repo)
    parkUseCase.Execute("KA-01-HH-1234", "White")

    leaveUseCase := usecases.NewLeaveCarUseCase(repo)
    car, charge, err := leaveUseCase.Execute(1, 3)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    if car.RegistrationNumber != "KA-01-HH-1234" {
        t.Errorf("Expected car registration number KA-01-HH-1234, got %s", car.RegistrationNumber)
    }
    if charge != 20 {
        t.Errorf("Expected charge 20, got %d", charge)
    }
}

func TestStatus(t *testing.T) {
    repo := &infra.ParkingLotRepository{}
    createUseCase := usecases.NewCreateParkingLotUseCase(repo)
    createUseCase.Execute(6)

    parkUseCase := usecases.NewParkCarUseCase(repo)
    parkUseCase.Execute("KA-01-HH-1234", "White")
    parkUseCase.Execute("KA-01-HH-9999", "Black")

    statusUseCase := usecases.NewStatusUseCase(repo)
    slots := statusUseCase.Execute()
    if len(slots) != 6 {
        t.Errorf("Expected 6 slots, got %d", len(slots))
    }
    if slots[0].Car == nil || slots[0].Car.RegistrationNumber != "KA-01-HH-1234" {
        t.Errorf("Expected slot 1 to have car KA-01-HH-1234")
    }
    if slots[1].Car == nil || slots[1].Car.RegistrationNumber != "KA-01-HH-9999" {
        t.Errorf("Expected slot 2 to have car KA-01-HH-9999")
    }
}
