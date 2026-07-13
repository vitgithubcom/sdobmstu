package repository

import (
    "database/sql"
    "pulse-backend/internal/domain"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) FindByUsername(username string) (*domain.User, error) {
    var user domain.User
    var lastLogin sql.NullTime
    err := r.db.QueryRow(`
        SELECT id, username, email, password_hash, full_name, role, is_active, created_at, updated_at, last_login
        FROM users WHERE username = $1
    `, username).Scan(
        &user.ID, &user.Username, &user.Email, &user.PasswordHash,
        &user.FullName, &user.Role, &user.IsActive,
        &user.CreatedAt, &user.UpdatedAt, &lastLogin,
    )
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    if lastLogin.Valid {
        user.LastLogin = &lastLogin.Time
    }
    return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
    var user domain.User
    var lastLogin sql.NullTime
    err := r.db.QueryRow(`
        SELECT id, username, email, password_hash, full_name, role, is_active, created_at, updated_at, last_login
        FROM users WHERE email = $1
    `, email).Scan(
        &user.ID, &user.Username, &user.Email, &user.PasswordHash,
        &user.FullName, &user.Role, &user.IsActive,
        &user.CreatedAt, &user.UpdatedAt, &lastLogin,
    )
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    if lastLogin.Valid {
        user.LastLogin = &lastLogin.Time
    }
    return &user, nil
}

func (r *UserRepository) FindByID(id int) (*domain.User, error) {
    var user domain.User
    var lastLogin sql.NullTime
    err := r.db.QueryRow(`
        SELECT id, username, email, password_hash, full_name, role, is_active, created_at, updated_at, last_login
        FROM users WHERE id = $1
    `, id).Scan(
        &user.ID, &user.Username, &user.Email, &user.PasswordHash,
        &user.FullName, &user.Role, &user.IsActive,
        &user.CreatedAt, &user.UpdatedAt, &lastLogin,
    )
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    if lastLogin.Valid {
        user.LastLogin = &lastLogin.Time
    }
    return &user, nil
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
    rows, err := r.db.Query(`
        SELECT id, username, email, full_name, role, is_active, created_at, last_login
        FROM users ORDER BY id
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []domain.User
    for rows.Next() {
        var u domain.User
        var lastLogin sql.NullTime
        err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.FullName,
            &u.Role, &u.IsActive, &u.CreatedAt, &lastLogin)
        if err != nil {
            return nil, err
        }
        if lastLogin.Valid {
            u.LastLogin = &lastLogin.Time
        }
        users = append(users, u)
    }
    return users, nil
}

// ===== CREATE — с возвратом id и created_at =====
func (r *UserRepository) Create(user *domain.User) error {
    err := r.db.QueryRow(`
        INSERT INTO users (username, email, password_hash, full_name, role)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at
    `, user.Username, user.Email, user.PasswordHash, user.FullName, user.Role).
        Scan(&user.ID, &user.CreatedAt)
    return err
}

func (r *UserRepository) Update(user *domain.User) error {
    _, err := r.db.Exec(`
        UPDATE users SET email=$1, full_name=$2, role=$3, updated_at=CURRENT_TIMESTAMP
        WHERE id=$4
    `, user.Email, user.FullName, user.Role, user.ID)
    return err
}

func (r *UserRepository) UpdatePassword(id int, hash string) error {
    _, err := r.db.Exec(`
        UPDATE users SET password_hash=$1, updated_at=CURRENT_TIMESTAMP
        WHERE id=$2
    `, hash, id)
    return err
}

func (r *UserRepository) ToggleActive(id int, isActive bool) error {
    _, err := r.db.Exec(`
        UPDATE users SET is_active=$1, updated_at=CURRENT_TIMESTAMP
        WHERE id=$2
    `, isActive, id)
    return err
}

func (r *UserRepository) UpdateLastLogin(id int) error {
    _, err := r.db.Exec(`
        UPDATE users SET last_login=CURRENT_TIMESTAMP WHERE id=$1
    `, id)
    return err
}