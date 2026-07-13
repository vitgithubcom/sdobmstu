package repository

import (
    "database/sql"
    "pulse-backend/internal/domain"
)

type AuditRepository struct {
    db *sql.DB
}

func NewAuditRepository(db *sql.DB) *AuditRepository {
    return &AuditRepository{db: db}
}

func (r *AuditRepository) GetAll() ([]domain.AuditLog, error) {
    rows, err := r.db.Query(`
        SELECT a.id, a.user_id, u.username, u.full_name, a.action, a.details, a.ip_address, a.created_at
        FROM audit_logs a
        LEFT JOIN users u ON a.user_id = u.id
        ORDER BY a.created_at DESC LIMIT 100
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var logs []domain.AuditLog
    for rows.Next() {
        var l domain.AuditLog
        var details sql.NullString
        err := rows.Scan(&l.ID, &l.UserID, &l.Username, &l.FullName,
            &l.Action, &details, &l.IPAddress, &l.CreatedAt)
        if err != nil {
            return nil, err
        }
        if details.Valid {
            l.Details = details.String
        }
        logs = append(logs, l)
    }
    return logs, nil
}

func (r *AuditRepository) Create(userID int, action, details, ip string) error {
    _, err := r.db.Exec(`
        INSERT INTO audit_logs (user_id, action, details, ip_address)
        VALUES ($1, $2, $3, $4)
    `, userID, action, details, ip)
    return err
}