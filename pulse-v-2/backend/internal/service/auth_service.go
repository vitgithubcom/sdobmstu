package service

import (
    "errors"
    "fmt"
    "pulse-backend/internal/domain"
    "pulse-backend/internal/repository"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    userRepo  *repository.UserRepository
    auditRepo *repository.AuditRepository
    jwtSecret string
}

func NewAuthService(userRepo *repository.UserRepository, auditRepo *repository.AuditRepository, jwtSecret string) *AuthService {
    return &AuthService{
        userRepo:  userRepo,
        auditRepo: auditRepo,
        jwtSecret: jwtSecret,
    }
}

func (s *AuthService) Login(req domain.LoginRequest, ip string) (*domain.LoginResponse, error) {
    // Поиск по username или email
    user, err := s.userRepo.FindByUsername(req.Login)
    if err != nil {
        return nil, err
    }
    if user == nil {
        user, err = s.userRepo.FindByEmail(req.Login)
        if err != nil || user == nil {
            return nil, errors.New("неверный логин или пароль")
        }
    }

    // Проверка активности
    if !user.IsActive {
        return nil, errors.New("учётная запись заблокирована")
    }

    // Проверка пароля
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        return nil, errors.New("неверный логин или пароль")
    }

    // Обновляем время входа
    s.userRepo.UpdateLastLogin(user.ID)

    // Логируем вход
    s.auditRepo.Create(user.ID, "login", "Успешный вход в систему", ip)

    // Генерируем JWT
    token, err := s.generateJWT(user)
    if err != nil {
        return nil, err
    }

    return &domain.LoginResponse{
        Token: token,
        User:  *user,
    }, nil
}

func (s *AuthService) ValidateToken(tokenString string) (*domain.User, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(s.jwtSecret), nil
    })
    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userId := int(claims["user_id"].(float64))
        user, err := s.userRepo.FindByID(userId)
        if err != nil || user == nil {
            return nil, errors.New("user not found")
        }
        if !user.IsActive {
            return nil, errors.New("user is not active")
        }
        return user, nil
    }

    return nil, errors.New("invalid token")
}

func (s *AuthService) generateJWT(user *domain.User) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id":    user.ID,
        "username":   user.Username,
        "role":       user.Role,
        "exp":        time.Now().Add(time.Hour * 24).Unix(),
        "iat":        time.Now().Unix(),
        "iss":        "pulse-backend",
        "sub":        "auth",
    })
    return token.SignedString([]byte(s.jwtSecret))
}