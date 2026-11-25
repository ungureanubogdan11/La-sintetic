package config

import (
	"fmt"
	"net/url"
	"os"
	"time"
)

type Config struct {
	Env  string // "dev" or "prod"
	Port string

	DB DBConfig

	JWTSecret string
	JWTTTL    time.Duration
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// DSN returns a libpq-style connection URL.
func (d DBConfig) DSN() string {
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(d.User, d.Password),
		Host:     d.Host + ":" + d.Port,
		Path:     d.Name,
		RawQuery: "sslmode=" + d.SSLMode + "&TimeZone=UTC",
	}
	return u.String()
}

func (c Config) IsDev() bool { return c.Env != "prod" }

func Load() (Config, error) {
	ttl, err := time.ParseDuration(getEnv("JWT_TTL", "24h"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid JWT_TTL: %w", err)
	}

	cfg := Config{
		Env:  getEnv("ENV", "dev"),
		Port: getEnv("PORT", "8080"),
		DB: DBConfig{
			Host:     getEnv("PG_HOST", "localhost"),
			Port:     getEnv("PG_PORT", "5432"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
			Name:     os.Getenv("PG_DB"),
			SSLMode:  getEnv("PG_SSL", "disable"),
		},
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTTTL:    ttl,
	}

	for name, val := range map[string]string{
		"PG_USER":     cfg.DB.User,
		"PG_PASSWORD": cfg.DB.Password,
		"PG_DB":       cfg.DB.Name,
		"JWT_SECRET":  cfg.JWTSecret,
	} {
		if val == "" {
			return Config{}, fmt.Errorf("missing required env var %s", name)
		}
	}
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
