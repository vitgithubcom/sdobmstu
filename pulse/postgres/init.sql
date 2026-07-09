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
('admin', 'admin@vitwebsite.ru', 'Администратор', 'admin', '$2a$10$N9qo8uLOickgx2ZMRZoMy.Mr/.3Z5QZwUz5k6Fv.3Xp2CkUf8kXUG'),
('analyst', 'analyst@vitwebsite.ru', 'Аналитик', 'analyst', '$2a$10$N9qo8uLOickgx2ZMRZoMy.Mr/.3Z5QZwUz5k6Fv.3Xp2CkUf8kXUG'),
('viewer', 'viewer@vitwebsite.ru', 'Наблюдатель', 'viewer', '$2a$10$N9qo8uLOickgx2ZMRZoMy.Mr/.3Z5QZwUz5k6Fv.3Xp2CkUf8kXUG'),
('manager', 'manager@vitwebsite.ru', 'Менеджер', 'manager', '$2a$10$N9qo8uLOickgx2ZMRZoMy.Mr/.3Z5QZwUz5k6Fv.3Xp2CkUf8kXUG');

-- KPI определения
INSERT INTO kpi_definitions (code, name, unit, direction, source_system) VALUES
('revenue', 'Выручка', 'тыс. руб', 'up', 'Mock-ERP'),
('oee', 'OEE (общая эффективность)', '%', 'up', 'Mock-MES'),
('downtime', 'Простои', 'часов', 'down', 'Mock-MES'),
('on_time', 'Заказы в срок', '%', 'up', 'Mock-CRM'),
('turnover', 'Склад (оборачиваемость)', 'дней', 'down', 'Mock-Warehouse');

-- KPI значения
INSERT INTO kpi_values (kpi_id, period_start, period_end, plan_value, fact_value) VALUES
(1, NOW() - INTERVAL '7 days', NOW(), 15200, 14280),
(2, NOW() - INTERVAL '7 days', NOW(), 78, 73.5),
(3, NOW() - INTERVAL '7 days', NOW(), 80, 124),
(4, NOW() - INTERVAL '7 days', NOW(), 92, 87.2),
(5, NOW() - INTERVAL '7 days', NOW(), 7.2, 8.3);

-- Тревоги
INSERT INTO alerts (system, message, severity) VALUES
('Mock-ERP', 'План продаж под угрозой (выполнение 68%)', 'critical'),
('Mock-MES', 'Станок #1042: превышение времени цикла', 'warning'),
('Mock-CRM', 'Не загружены сделки за последний час', 'critical'),
('Mock-Warehouse', 'Остатки по группе А ниже нормы', 'info');

-- Интеграции
INSERT INTO integrations (name, status, last_sync, lag_seconds) VALUES
('Mock-ERP', 'ok', NOW() - INTERVAL '12 seconds', 12),
('Mock-MES', 'ok', NOW() - INTERVAL '34 seconds', 34),
('Mock-CRM', 'warning', NOW() - INTERVAL '14 minutes', 840),
('Mock-Warehouse', 'ok', NOW() - INTERVAL '2 minutes', 120);