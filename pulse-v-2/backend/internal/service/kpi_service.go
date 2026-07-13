package service

import (
    "pulse-backend/internal/domain"
    "pulse-backend/internal/repository"
)

type KPIService struct {
    repo *repository.KPIRepository
}

func NewKPIService(repo *repository.KPIRepository) *KPIService {
    return &KPIService{repo: repo}
}

func (s *KPIService) GetAll() ([]domain.KPI, error) {
    return s.repo.GetAll()
}

func (s *KPIService) GetByID(id int) (*domain.KPI, error) {
    return s.repo.GetByID(id)
}

func (s *KPIService) GetChartData() ([]domain.ChartData, error) {
    return s.repo.GetChartData()
}