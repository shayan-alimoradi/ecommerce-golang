package config

import "os"

type DBConfig struct {
	ADDR         string
	MaxOpenConns string
	MaxIdleConns string
	MaxIdleTime  string
}

var Envs = initConfig()

func initConfig() DBConfig {
	return DBConfig{
		ADDR:         getEnv("DB_ADDR", "postgresql://admin:adminpassword@localhost:5434/ecom?sslmode=disable"),
		MaxOpenConns: getEnv("DB_MAX_OPEN_CONNS", "30"),
		MaxIdleConns: getEnv("DB_MAX_IDLE_CONNS", "30"),
		MaxIdleTime:  getEnv("DB_MAX_IDLE_TIME", "15m"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
