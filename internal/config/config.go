package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config contém todas as configurações da aplicação
type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Email    EmailConfig
	Maps     MapsConfig
	JWT      JWTConfig
}

// NewConfig cria uma nova configuração da aplicação
func NewConfig() *Config {
	return &Config{
		App:      NewAppConfig(),
		Database: NewDatabaseConfig(),
	}
}

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type EmailConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	From     string
}

type MapsConfig struct {
	APIKey string
}

type JWTConfig struct {
	Secret     string
	Expiration string
}

func Load() (*Config, error) {
	// Configurações do Servidor
	serverConfig := ServerConfig{
		Host: getEnv("APP_HOST", "0.0.0.0"),
		Port: getEnv("APP_PORT", "8080"),
	}

	// Configurações do Banco de Dados
	dbConfig := DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "agencia_viagens"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
	}

	// Configurações de E-mail
	emailPort, _ := strconv.Atoi(getEnv("SMTP_PORT", "587"))
	emailConfig := EmailConfig{
		Host:     getEnv("SMTP_HOST", "smtp.gmail.com"),
		Port:     emailPort,
		User:     getEnv("SMTP_USER", ""),
		Password: getEnv("SMTP_PASSWORD", ""),
		From:     getEnv("SMTP_FROM", ""),
	}

	// Configurações do Google Maps
	mapsConfig := MapsConfig{
		APIKey: getEnv("GOOGLE_MAPS_API_KEY", ""),
	}

	// Configurações JWT
	jwtConfig := JWTConfig{
		Secret:     getEnv("JWT_SECRET", "your-secret-key"),
		Expiration: getEnv("JWT_EXPIRATION", "24h"),
	}

	return &Config{
		Server:   serverConfig,
		Database: dbConfig,
		Email:    emailConfig,
		Maps:     mapsConfig,
		JWT:      jwtConfig,
	}, nil
}

// getEnv retorna o valor da variável de ambiente ou um valor padrão
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetDSN retorna a string de conexão do PostgreSQL
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
} 