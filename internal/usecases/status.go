package usecases

import (
    "parking/internal/domain"
    "parking/internal/infra"
)

type StatusUseCase struct {
    Repo *infra.ParkingLotRepository
}

func NewStatusUseCase(repo *infra.ParkingLotRepository) *StatusUseCase {
    return &StatusUseCase{Repo: repo}
}

func (usecase *StatusUseCase) Execute() []*domain.Slot {
    return usecase.Repo.GetStatus()
}
