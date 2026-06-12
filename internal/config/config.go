package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Upload   UploadConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type ServerConfig struct {
	Port string
	Env  string
}

type UploadConfig struct {
	Path        string
	MaxFileSize int64
}

// Load charge la configuration depuis les variables d'environnement
func Load() (*Config, error) {
	// Charger le fichier .env si présent (ignoré en production)
	_ = godotenv.Load()

	cfg := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "127.0.0.1"),
			Port:     getEnvAsInt("DB_PORT", 3306),
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "portefolio"),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Env:  getEnv("ENV", "development"),
		},
		Upload: UploadConfig{
			Path:        getEnv("UPLOAD_PATH", "./web/static/uploads"),
			MaxFileSize: getEnvAsInt64("MAX_UPLOAD_SIZE", 52428800), // 50MB par défaut
		},
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuration invalide: %w", err)
	}

	return cfg, nil
}

// Validate vérifie que la configuration est valide
func (c *Config) Validate() error {
	if c.Database.User == "" {
		return fmt.Errorf("DB_USER est requis")
	}
	if c.Database.Password == "" {
		return fmt.Errorf("DB_PASSWORD est requis")
	}
	if c.Database.Name == "" {
		return fmt.Errorf("DB_NAME est requis")
	}
	if c.Upload.MaxFileSize <= 0 {
		return fmt.Errorf("MAX_UPLOAD_SIZE doit être positif")
	}
	return nil
}

// DSN retourne la chaîne de connexion à la base de données
func (c *Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)
}

// getEnv récupère une variable d'environnement avec une valeur par défaut
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt récupère une variable d'environnement en tant qu'entier
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// getEnvAsInt64 récupère une variable d'environnement en tant qu'int64
func getEnvAsInt64(key string, defaultValue int64) int64 {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		return defaultValue
	}
	return value
}
