package usecases

import (
    "parking/internal/infra"
)

type CreateParkingLotUseCase struct {
    Repo *infra.ParkingLotRepository
}

func NewCreateParkingLotUseCase(repo *infra.ParkingLotRepository) *CreateParkingLotUseCase {
    return &CreateParkingLotUseCase{Repo: repo}
}

func (usecase *CreateParkingLotUseCase) Execute(capacity int) {
    usecase.Repo.CreateParkingLot(capacity)
}
