package config

import (
	"os"
	"time"
)

type Config struct {
	Username      string
	Password      string
	CookiePath    string
	ActionDelayMin time.Duration
	ActionDelayMax time.Duration
}

func Load() Config {
	return Config{
		Username:       os.Getenv("LOGIN_USERNAME"),
		Password:       os.Getenv("LOGIN_PASSWORD"),
		CookiePath:     getEnv("COOKIE_PATH", "data/session.json"),
		ActionDelayMin: 500 * time.Millisecond,
		ActionDelayMax: 1500 * time.Millisecond,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
