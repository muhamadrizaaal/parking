package usecases

import (
    "parking/internal/domain"
    "parking/internal/infra"
)

type LeaveCarUseCase struct {
    Repo *infra.ParkingLotRepository
}

func NewLeaveCarUseCase(repo *infra.ParkingLotRepository) *LeaveCarUseCase {
    return &LeaveCarUseCase{Repo: repo}
}

func (usecase *LeaveCarUseCase) Execute(slotNumber, hours int) (*domain.Car, int, error) {
    return usecase.Repo.LeaveCar(slotNumber, hours)
}
