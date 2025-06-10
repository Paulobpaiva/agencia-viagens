package config

import (
	"os"
	"time"
)

// JWTConfig contém as configurações do JWT
type JWTConfig struct {
	Secret           string
	Expiration       time.Duration
	RefreshExpiration time.Duration
	Issuer           string
}

// NewJWTConfig cria uma nova configuração do JWT a partir das variáveis de ambiente
func NewJWTConfig() JWTConfig {
	expiration, _ := time.ParseDuration(getEnv("JWT_EXPIRATION", "24h"))
	refreshExpiration, _ := time.ParseDuration(getEnv("JWT_REFRESH_EXPIRATION", "720h")) // 30 dias

	return JWTConfig{
		Secret:           getEnv("JWT_SECRET", "your-secret-key"),
		Expiration:       expiration,
		RefreshExpiration: refreshExpiration,
		Issuer:           getEnv("JWT_ISSUER", "agencia-viagens"),
	}
} 