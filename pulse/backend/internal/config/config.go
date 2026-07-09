package config

import (
    "os"
    "strconv"
)

type Config struct {
    DBHost     string
    DBPort     int
    DBUser     string
    DBPassword string
    DBName     string
    JWTSecret  string
    Port       int
}

func Load() (*Config, error) {
    port, _ := strconv.Atoi(getEnv("PORT", "8082"))
    dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

    return &Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     dbPort,
        DBUser:     getEnv("DB_USER", "pulse_user"),
        DBPassword: getEnv("DB_PASSWORD", "pulse_pass"),
        DBName:     getEnv("DB_NAME", "pulse_db"),
        JWTSecret:  getEnv("JWT_SECRET", "super-secret-key"),
        Port:       port,
    }, nil
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}