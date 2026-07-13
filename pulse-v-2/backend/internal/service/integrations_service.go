package service

import (
    "pulse-backend/internal/domain"
    "pulse-backend/internal/repository"
)

type IntegrationsService struct {
    repo *repository.IntegrationsRepository
}

func NewIntegrationsService(repo *repository.IntegrationsRepository) *IntegrationsService {
    return &IntegrationsService{repo: repo}
}

func (s *IntegrationsService) GetAll() ([]domain.Integration, error) {
    return s.repo.GetAll()
}