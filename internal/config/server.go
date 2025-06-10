package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// ServerConfig contém as configurações do servidor
type ServerConfig struct {
	Host            string
	Port            int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
	CorsOrigins     []string
	CorsMethods     []string
	CorsHeaders     []string
}

// NewServerConfig cria uma nova configuração do servidor a partir das variáveis de ambiente
func NewServerConfig() ServerConfig {
	port, _ := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	readTimeout, _ := time.ParseDuration(getEnv("SERVER_READ_TIMEOUT", "5s"))
	writeTimeout, _ := time.ParseDuration(getEnv("SERVER_WRITE_TIMEOUT", "10s"))
	shutdownTimeout, _ := time.ParseDuration(getEnv("SERVER_SHUTDOWN_TIMEOUT", "15s"))

	return ServerConfig{
		Host:            getEnv("SERVER_HOST", "0.0.0.0"),
		Port:            port,
		ReadTimeout:     readTimeout,
		WriteTimeout:    writeTimeout,
		ShutdownTimeout: shutdownTimeout,
		CorsOrigins:     getEnvSlice("CORS_ORIGINS", []string{"*"}),
		CorsMethods:     getEnvSlice("CORS_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		CorsHeaders:     getEnvSlice("CORS_HEADERS", []string{"Content-Type", "Authorization"}),
	}
}

// getEnvSlice retorna um slice de strings a partir de uma variável de ambiente
func getEnvSlice(key string, defaultValue []string) []string {
	if value, exists := os.LookupEnv(key); exists {
		return strings.Split(value, ",")
	}
	return defaultValue
} 