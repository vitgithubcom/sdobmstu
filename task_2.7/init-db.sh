#!/bin/sh

if [ ! -f /app/users.db ]; then
    echo "Initializing database..."
    sqlite3 /app/users.db < /app/create_db.sql
    echo "Database initialized successfully."
else
    echo "Database already exists."
fi