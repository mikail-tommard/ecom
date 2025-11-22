package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBSSLMode     string
	JWTExpiration int64
	JWTSecret     string
}

var Envs = initEnv()

func initEnv() Env {
	godotenv.Load()

	return Env{
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "5432"),
		DBUser:        getEnv("DB_USER", "postgres"),
		DBPassword:    getEnv("DB_PASSWORD", "1234"),
		DBName:        getEnv("DB_NAME", "ecom"),
		DBSSLMode:     getEnv("DB_SSL_MODE", "disable"),
		JWTExpiration: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:     getEnv("JWT_SECRET", "not-secret-secter-anymore?"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
