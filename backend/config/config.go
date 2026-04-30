// Package config loads and validates all application configuration from
// environment variables. It uses viper for env binding and exposes a single
// typed Config struct consumed by the rest of the application.
package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config is the root configuration struct. Every sub-domain has its own
// nested struct so callers only reference what they need.
type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Rabbit   RabbitConfig
	JWT      JWTConfig
	Log      LogConfig
	CORS     CORSConfig
	Events   EventsConfig
}

type AppConfig struct {
	Env  string `mapstructure:"APP_ENV"`
	Port int    `mapstructure:"APP_PORT"`
	Name string `mapstructure:"APP_NAME"`
}

type DatabaseConfig struct {
	URL             string        `mapstructure:"DATABASE_URL"`
	MaxConns        int32         `mapstructure:"DATABASE_MAX_CONNS"`
	MinConns        int32         `mapstructure:"DATABASE_MIN_CONNS"`
	MaxConnLifetime time.Duration `mapstructure:"DATABASE_MAX_CONN_LIFETIME"`
	MaxConnIdleTime time.Duration `mapstructure:"DATABASE_MAX_CONN_IDLE_TIME"`
}

type RedisConfig struct {
	URL      string `mapstructure:"REDIS_URL"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DB"`
}

type RabbitConfig struct {
	URL      string `mapstructure:"RABBITMQ_URL"`
	Exchange string `mapstructure:"RABBITMQ_EXCHANGE"`
}

type JWTConfig struct {
	Secret        string        `mapstructure:"JWT_SECRET"`
	Expiry        time.Duration `mapstructure:"JWT_EXPIRY"`
	RefreshExpiry time.Duration `mapstructure:"JWT_REFRESH_EXPIRY"`
}

type LogConfig struct {
	Level  string `mapstructure:"LOG_LEVEL"`
	Format string `mapstructure:"LOG_FORMAT"`
}

type CORSConfig struct {
	Origins []string
	MaxAge  int `mapstructure:"CORS_MAX_AGE"`
}

type EventsConfig struct {
	Broker string `mapstructure:"EVENT_BROKER"`
}

// Load reads configuration from environment variables (and optionally a .env
// file) and returns a validated Config. Returns an error if any required value
// is missing or invalid.
func Load() (*Config, error) {
	v := viper.New()

	// Read from environment
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Attempt to read a .env file; ignore if absent (Docker already injects env).
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	_ = v.ReadInConfig()

	setDefaults(v)

	cfg := &Config{}

	// App
	cfg.App.Env = v.GetString("APP_ENV")
	cfg.App.Port = v.GetInt("APP_PORT")
	cfg.App.Name = v.GetString("APP_NAME")

	// Database
	cfg.Database.URL = v.GetString("DATABASE_URL")
	cfg.Database.MaxConns = int32(v.GetInt("DATABASE_MAX_CONNS"))
	cfg.Database.MinConns = int32(v.GetInt("DATABASE_MIN_CONNS"))
	cfg.Database.MaxConnLifetime = v.GetDuration("DATABASE_MAX_CONN_LIFETIME")
	cfg.Database.MaxConnIdleTime = v.GetDuration("DATABASE_MAX_CONN_IDLE_TIME")

	// Redis
	cfg.Redis.URL = v.GetString("REDIS_URL")
	cfg.Redis.Password = v.GetString("REDIS_PASSWORD")
	cfg.Redis.DB = v.GetInt("REDIS_DB")

	// RabbitMQ
	cfg.Rabbit.URL = v.GetString("RABBITMQ_URL")
	cfg.Rabbit.Exchange = v.GetString("RABBITMQ_EXCHANGE")

	// JWT
	cfg.JWT.Secret = v.GetString("JWT_SECRET")
	cfg.JWT.Expiry = v.GetDuration("JWT_EXPIRY")
	cfg.JWT.RefreshExpiry = v.GetDuration("JWT_REFRESH_EXPIRY")

	// Logging
	cfg.Log.Level = v.GetString("LOG_LEVEL")
	cfg.Log.Format = v.GetString("LOG_FORMAT")

	// CORS — split comma-separated origins
	rawOrigins := v.GetString("CORS_ORIGINS")
	for _, o := range strings.Split(rawOrigins, ",") {
		if t := strings.TrimSpace(o); t != "" {
			cfg.CORS.Origins = append(cfg.CORS.Origins, t)
		}
	}
	cfg.CORS.MaxAge = v.GetInt("CORS_MAX_AGE")

	// Events
	cfg.Events.Broker = v.GetString("EVENT_BROKER")

	if err := validate(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// IsDevelopment returns true when running in dev mode.
func (c *Config) IsDevelopment() bool {
	return c.App.Env == "development"
}

// IsProduction returns true when running in production mode.
func (c *Config) IsProduction() bool {
	return c.App.Env == "production"
}

// Addr returns the listen address for the HTTP server.
func (c *Config) Addr() string {
	return fmt.Sprintf(":%d", c.App.Port)
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("APP_ENV", "development")
	v.SetDefault("APP_PORT", 8080)
	v.SetDefault("APP_NAME", "platform")

	v.SetDefault("DATABASE_MAX_CONNS", 25)
	v.SetDefault("DATABASE_MIN_CONNS", 5)
	v.SetDefault("DATABASE_MAX_CONN_LIFETIME", "1h")
	v.SetDefault("DATABASE_MAX_CONN_IDLE_TIME", "30m")

	v.SetDefault("REDIS_DB", 0)

	v.SetDefault("RABBITMQ_EXCHANGE", "platform.events")

	v.SetDefault("JWT_EXPIRY", "15m")
	v.SetDefault("JWT_REFRESH_EXPIRY", "168h")

	v.SetDefault("LOG_LEVEL", "info")
	v.SetDefault("LOG_FORMAT", "json")

	v.SetDefault("CORS_ORIGINS", "http://localhost:3000")
	v.SetDefault("CORS_MAX_AGE", 86400)

	v.SetDefault("EVENT_BROKER", "inmemory")
}

func validate(cfg *Config) error {
	if cfg.Database.URL == "" {
		return fmt.Errorf("config: DATABASE_URL is required")
	}
	if cfg.JWT.Secret == "" {
		return fmt.Errorf("config: JWT_SECRET is required")
	}
	if len(cfg.JWT.Secret) < 32 {
		return fmt.Errorf("config: JWT_SECRET must be at least 32 characters")
	}
	return nil
}
