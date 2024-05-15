package usecases

import (
    "parking/internal/domain"
    "parking/internal/infra"
)

type ParkCarUseCase struct {
    Repo *infra.ParkingLotRepository
}

func NewParkCarUseCase(repo *infra.ParkingLotRepository) *ParkCarUseCase {
    return &ParkCarUseCase{Repo: repo}
}

func (usecase *ParkCarUseCase) Execute(registrationNumber, color string) (int, error) {
    car := &domain.Car{RegistrationNumber: registrationNumber, Color: color}
    return usecase.Repo.ParkCar(car)
}
