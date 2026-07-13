package service

import (
    "pulse-backend/internal/domain"
    "pulse-backend/internal/repository"
)

type AlertsService struct {
    repo *repository.AlertsRepository
}

func NewAlertsService(repo *repository.AlertsRepository) *AlertsService {
    return &AlertsService{repo: repo}
}

func (s *AlertsService) GetAll() ([]domain.Alert, error) {
    return s.repo.GetAll()
}