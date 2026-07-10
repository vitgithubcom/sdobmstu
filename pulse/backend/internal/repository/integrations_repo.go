package repository

import (
    "database/sql"
    "pulse-backend/internal/domain"
)

type IntegrationsRepository struct {
    db *sql.DB
}

func NewIntegrationsRepository(db *sql.DB) *IntegrationsRepository {
    return &IntegrationsRepository{db: db}
}

func (r *IntegrationsRepository) GetAll() ([]domain.Integration, error) {
    rows, err := r.db.Query(`
        SELECT id, name, status, last_sync, lag_seconds, error_message
        FROM integrations ORDER BY id
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var integrations []domain.Integration
    for rows.Next() {
        var i domain.Integration
        var errorMsg sql.NullString
        err := rows.Scan(&i.ID, &i.Name, &i.Status, &i.LastSync,
            &i.LagSeconds, &errorMsg)
        if err != nil {
            return nil, err
        }
        if errorMsg.Valid {
            i.ErrorMessage = &errorMsg.String
        }
        integrations = append(integrations, i)
    }
    return integrations, nil
}