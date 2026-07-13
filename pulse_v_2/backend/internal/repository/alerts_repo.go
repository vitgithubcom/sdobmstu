package repository

import (
    "database/sql"
    "pulse-backend/internal/domain"
)

type AlertsRepository struct {
    db *sql.DB
}

func NewAlertsRepository(db *sql.DB) *AlertsRepository {
    return &AlertsRepository{db: db}
}

func (r *AlertsRepository) GetAll() ([]domain.Alert, error) {
    rows, err := r.db.Query(`
        SELECT id, system, message, severity, is_active, created_at, resolved_at
        FROM alerts ORDER BY created_at DESC LIMIT 20
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var alerts []domain.Alert
    for rows.Next() {
        var a domain.Alert
        err := rows.Scan(&a.ID, &a.System, &a.Message, &a.Severity,
            &a.IsActive, &a.CreatedAt, &a.ResolvedAt)
        if err != nil {
            return nil, err
        }
        alerts = append(alerts, a)
    }
    return alerts, nil
}