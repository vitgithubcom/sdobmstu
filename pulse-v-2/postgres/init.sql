-- ============================================
-- Создание таблиц
-- ============================================

-- Пользователи
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100),
    role VARCHAR(20) DEFAULT 'viewer',
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP
);

-- KPI определения
CREATE TABLE IF NOT EXISTS kpi_definitions (
    id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    unit VARCHAR(20),
    direction VARCHAR(10) CHECK (direction IN ('up', 'down')),
    source_system VARCHAR(50),
    is_active BOOLEAN DEFAULT TRUE
);

-- План/факт KPI
CREATE TABLE IF NOT EXISTS kpi_values (
    id SERIAL PRIMARY KEY,
    kpi_id INTEGER REFERENCES kpi_definitions(id),
    period_start TIMESTAMP NOT NULL,
    period_end TIMESTAMP NOT NULL,
    plan_value DECIMAL(15,2),
    fact_value DECIMAL(15,2),
    UNIQUE(kpi_id, period_start)
);

-- Тревоги
CREATE TABLE IF NOT EXISTS alerts (
    id SERIAL PRIMARY KEY,
    system VARCHAR(50),
    message TEXT,
    severity VARCHAR(20) CHECK (severity IN ('info', 'warning', 'critical')),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    resolved_at TIMESTAMP
);

-- Интеграции
CREATE TABLE IF NOT EXISTS integrations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    status VARCHAR(20) DEFAULT 'ok',
    last_sync TIMESTAMP,
    lag_seconds INTEGER DEFAULT 0,
    error_message TEXT
);

-- Логи аудита
CREATE TABLE IF NOT EXISTS audit_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    action VARCHAR(100),
    details JSONB,
    ip_address VARCHAR(45),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ============================================
-- Мок-данные (для разработки)
-- ============================================

-- Пользователи (пароль: admin123 → хеш bcrypt)
INSERT INTO users (username, email, full_name, role, password_hash) VALUES
('admin', 'admin@localhost', 'Администратор', 'admin', '$2b$12$OlRu4/kPaNEhyJGQJci2X.zQTbWvxoCXaQp0I7j0bFuIy.XW/Irfu'),
('analyst', 'analyst@localhost', 'Аналитик', 'analyst', '$2b$12$OlRu4/kPaNEhyJGQJci2X.zQTbWvxoCXaQp0I7j0bFuIy.XW/Irfu'),
('viewer', 'viewer@localhost', 'Наблюдатель', 'viewer', '$2b$12$OlRu4/kPaNEhyJGQJci2X.zQTbWvxoCXaQp0I7j0bFuIy.XW/Irfu'),
('manager', 'manager@localhost', 'Менеджер', 'manager', '$2b$12$OlRu4/kPaNEhyJGQJci2X.zQTbWvxoCXaQp0I7j0bFuIy.XW/Irfu')
ON CONFLICT (username) DO NOTHING;

-- KPI определения
INSERT INTO kpi_definitions (code, name, unit, direction, source_system) VALUES
('revenue', 'Выручка', 'тыс. руб', 'up', 'Mock-ERP'),
('oee', 'OEE (общая эффективность)', '%', 'up', 'Mock-MES'),
('downtime', 'Простои', 'часов', 'down', 'Mock-MES'),
('on_time', 'Заказы в срок', '%', 'up', 'Mock-CRM'),
('turnover', 'Склад (оборачиваемость)', 'дней', 'down', 'Mock-Warehouse')
ON CONFLICT (code) DO NOTHING;

-- KPI значения за последние 7 дней (для графика)
INSERT INTO kpi_values (kpi_id, period_start, period_end, plan_value, fact_value)
SELECT 
    k.id,
    dates.day::timestamp as period_start,
    (dates.day::timestamp + INTERVAL '1 day') as period_end,
    CASE k.id
        WHEN 1 THEN 15200   -- Выручка
        WHEN 2 THEN 78      -- OEE
        WHEN 3 THEN 80      -- Простои
        WHEN 4 THEN 92      -- Заказы в срок
        WHEN 5 THEN 7.2     -- Склад
    END as plan_value,
    CASE k.id
        WHEN 1 THEN 14280 + (random() * 1000 - 500)  -- Выручка
        WHEN 2 THEN 73.5 + (random() * 10 - 5)       -- OEE
        WHEN 3 THEN 124 + (random() * 20 - 10)       -- Простои
        WHEN 4 THEN 87.2 + (random() * 10 - 5)       -- Заказы в срок
        WHEN 5 THEN 8.3 + (random() * 2 - 1)         -- Склад
    END as fact_value
FROM kpi_definitions k
CROSS JOIN (
    SELECT generate_series(
        NOW() - INTERVAL '6 days',
        NOW(),
        INTERVAL '1 day'
    ) as day
) dates
ON CONFLICT (kpi_id, period_start) DO UPDATE SET
    fact_value = EXCLUDED.fact_value,
    plan_value = EXCLUDED.plan_value;

-- Тревоги
INSERT INTO alerts (system, message, severity) VALUES
('Mock-ERP', 'План продаж под угрозой (выполнение 68%)', 'critical'),
('Mock-MES', 'Станок #1042: превышение времени цикла', 'warning'),
('Mock-CRM', 'Не загружены сделки за последний час', 'critical'),
('Mock-Warehouse', 'Остатки по группе А ниже нормы', 'info')
ON CONFLICT (id) DO NOTHING;

-- Интеграции
INSERT INTO integrations (name, status, last_sync, lag_seconds) VALUES
('Mock-ERP', 'ok', NOW() - INTERVAL '12 seconds', 12),
('Mock-MES', 'ok', NOW() - INTERVAL '34 seconds', 34),
('Mock-CRM', 'warning', NOW() - INTERVAL '14 minutes', 840),
('Mock-Warehouse', 'ok', NOW() - INTERVAL '2 minutes', 120)
ON CONFLICT (name) DO NOTHING;

-- Логи аудита (тестовые записи)
INSERT INTO audit_logs (user_id, action, details, ip_address) VALUES
(1, 'login', '{"status":"success"}', '127.0.0.1'),
(1, 'view_dashboard', '{"page":"dashboard"}', '192.168.1.1'),
(1, 'view_kpi', '{"kpi_id":1}', '192.168.1.1'),
(1, 'logout', '{}', '127.0.0.1')
ON CONFLICT (id) DO NOTHING;