package app

import (
    "parking/internal/domain"
    "parking/internal/infra"
)

type ParkingLotService struct {
    Repo *infra.ParkingLotRepository
}

func NewParkingLotService(repo *infra.ParkingLotRepository) *ParkingLotService {
    return &ParkingLotService{Repo: repo}
}

func (service *ParkingLotService) CreateParkingLot(capacity int) {
    service.Repo.CreateParkingLot(capacity)
}

func (service *ParkingLotService) ParkCar(registrationNumber, color string) (int, error) {
    car := &domain.Car{RegistrationNumber: registrationNumber, Color: color}
    return service.Repo.ParkCar(car)
}

func (service *ParkingLotService) LeaveCar(slotNumber, hours int) (*domain.Car, int, error) {
    return service.Repo.LeaveCar(slotNumber, hours)
}

func (service *ParkingLotService) GetStatus() []*domain.Slot {
    return service.Repo.GetStatus()
}
