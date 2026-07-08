CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

INSERT OR IGNORE INTO users (email, password) VALUES 
    ('user1@example.com', 'pass123'),
    ('user2@example.com', 'qwerty456'),
    ('user3@example.com', 'letmein789'),
    ('user4@example.com', '123456789'),
    ('user5@example.com', 'password'),
    ('admin@example.com', 'admin123'),
    ('test@example.com', 'test123'),
    ('demo@example.com', 'demo456'),
    ('hello@example.com', 'world789'),
    ('golang@example.com', 'gopher123');