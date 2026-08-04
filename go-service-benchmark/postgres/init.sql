-- Создаем таблицу пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    bio TEXT,
    avatar TEXT
);

-- Индекс для быстрого поиска по id (уже есть PRIMARY KEY)
-- Но добавим дополнительный индекс для email, если понадобится
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Генерируем 10 000 тестовых пользователей
INSERT INTO users (name, email, bio, avatar)
SELECT 
    'User ' || generate_series,
    'user' || generate_series || '@example.com',
    'Bio of user ' || generate_series || '. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
    'https://avatar.example.com/' || generate_series || '.jpg'
FROM generate_series(1, 10000);

-- Создаем расширение для сбора статистики
CREATE EXTENSION IF NOT EXISTS pg_stat_statements;