package service

import (
    "errors"
    "pulse-backend/internal/domain"
    "pulse-backend/internal/repository"

    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]domain.User, error) {
    return s.repo.GetAll()
}

func (s *UserService) GetByID(id int) (*domain.User, error) {
    return s.repo.FindByID(id)
}

func (s *UserService) Create(user *domain.User, password string) error {
    // Проверка существования
    existing, _ := s.repo.FindByUsername(user.Username)
    if existing != nil {
        return errors.New("пользователь с таким логином уже существует")
    }
    existing, _ = s.repo.FindByEmail(user.Email)
    if existing != nil {
        return errors.New("пользователь с таким email уже существует")
    }

    // Хеширование пароля
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.PasswordHash = string(hash)

    return s.repo.Create(user)
}

func (s *UserService) Update(user *domain.User) error {
    return s.repo.Update(user)
}

func (s *UserService) UpdatePassword(id int, oldPassword, newPassword string) error {
    user, err := s.repo.FindByID(id)
    if err != nil || user == nil {
        return errors.New("пользователь не найден")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword)); err != nil {
        return errors.New("неверный текущий пароль")
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    return s.repo.UpdatePassword(id, string(hash))
}

func (s *UserService) ToggleActive(id int, isActive bool) error {
    return s.repo.ToggleActive(id, isActive)
}