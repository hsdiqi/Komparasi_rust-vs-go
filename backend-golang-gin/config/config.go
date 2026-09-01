package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort        string
	DatabaseURL    string
	MaxOpenConns   int
	MaxIdleConns   int
}

func Load() Config {
	_ = godotenv.Load()
	return Config{
		AppPort:      getEnv("APP_PORT", "8080"),
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/perf_db?sslmode=disable"),
		MaxOpenConns: getEnvInt("DB_MAX_OPEN_CONNS", 25),
		MaxIdleConns: getEnvInt("DB_MAX_IDLE_CONNS", 25),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" { return v }
	return fallback
}

func getEnvInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" { return fallback }
	n, err := strconv.Atoi(v)
	if err != nil { return fallback }
	return n
}
