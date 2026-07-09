package repository

import (
    "database/sql"
    "pulse-backend/internal/domain"
)

type KPIRepository struct {
    db *sql.DB
}

func NewKPIRepository(db *sql.DB) *KPIRepository {
    return &KPIRepository{db: db}
}

func (r *KPIRepository) GetAll() ([]domain.KPI, error) {
    rows, err := r.db.Query(`
        SELECT k.id, k.code, k.name, k.unit, k.direction, k.source_system,
               COALESCE(v.fact_value, 0) as fact_value,
               COALESCE(v.plan_value, 0) as plan_value
        FROM kpi_definitions k
        LEFT JOIN kpi_values v ON k.id = v.kpi_id 
        WHERE k.is_active = true
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var kpis []domain.KPI
    for rows.Next() {
        var k domain.KPI
        err := rows.Scan(&k.ID, &k.Code, &k.Name, &k.Unit, &k.Direction,
            &k.Source, &k.Value, &k.Plan)
        if err != nil {
            return nil, err
        }
        k.Delta = 0
        if k.Plan > 0 {
            k.Delta = (k.Value - k.Plan) / k.Plan * 100
        }
        k.Completion = 0
        if k.Plan > 0 {
            k.Completion = (k.Value / k.Plan) * 100
        }

        if k.Delta < -5 {
            k.Status = "critical"
        } else if k.Delta < 0 {
            k.Status = "warning"
        } else {
            k.Status = "ok"
        }
        kpis = append(kpis, k)
    }
    return kpis, nil
}

func (r *KPIRepository) GetByID(id int) (*domain.KPI, error) {
    var k domain.KPI
    err := r.db.QueryRow(`
        SELECT k.id, k.code, k.name, k.unit, k.direction, k.source_system,
               COALESCE(v.fact_value, 0) as fact_value,
               COALESCE(v.plan_value, 0) as plan_value
        FROM kpi_definitions k
        LEFT JOIN kpi_values v ON k.id = v.kpi_id 
        WHERE k.id = $1
    `, id).Scan(&k.ID, &k.Code, &k.Name, &k.Unit, &k.Direction,
        &k.Source, &k.Value, &k.Plan)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    k.Delta = 0
    if k.Plan > 0 {
        k.Delta = (k.Value - k.Plan) / k.Plan * 100
    }
    k.Completion = 0
    if k.Plan > 0 {
        k.Completion = (k.Value / k.Plan) * 100
    }
    if k.Delta < -5 {
        k.Status = "critical"
    } else if k.Delta < 0 {
        k.Status = "warning"
    } else {
        k.Status = "ok"
    }

    // История
    rows, err := r.db.Query(`
        SELECT TO_CHAR(period_start, 'YYYY-MM-DD') as period,
               COALESCE(fact_value, 0) as fact, COALESCE(plan_value, 0) as plan
        FROM kpi_values WHERE kpi_id = $1 ORDER BY period_start
    `, id)
    if err == nil {
        defer rows.Close()
        for rows.Next() {
            var h domain.History
            rows.Scan(&h.Period, &h.Fact, &h.Plan)
            k.History = append(k.History, h)
        }
    }

    return &k, nil
}

func (r *KPIRepository) GetChartData() ([]domain.ChartData, error) {
    rows, err := r.db.Query(`
        SELECT 'Пн' as name, COALESCE(SUM(fact_value), 0) as fact, COALESCE(SUM(plan_value), 0) as plan
        FROM kpi_values
        UNION ALL SELECT 'Вт', COALESCE(SUM(fact_value), 0), COALESCE(SUM(plan_value), 0)
        FROM kpi_values
        UNION ALL SELECT 'Ср', COALESCE(SUM(fact_value), 0), COALESCE(SUM(plan_value), 0)
        FROM kpi_values
        UNION ALL SELECT 'Чт', COALESCE(SUM(fact_value), 0), COALESCE(SUM(plan_value), 0)
        FROM kpi_values
        UNION ALL SELECT 'Пт', COALESCE(SUM(fact_value), 0), COALESCE(SUM(plan_value), 0)
        FROM kpi_values
        UNION ALL SELECT 'Сб', COALESCE(SUM(fact_value), 0), COALESCE(SUM(plan_value), 0)
        FROM kpi_values
        UNION ALL SELECT 'Вс', COALESCE(SUM(fact_value), 0), COALESCE(SUM(plan_value), 0)
        FROM kpi_values
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var data []domain.ChartData
    for rows.Next() {
        var d domain.ChartData
        rows.Scan(&d.Name, &d.Fact, &d.Plan)
        data = append(data, d)
    }
    return data, nil
}