package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("sua_chave_secreta_supersegura") // Troque por uma chave forte em produção

// Claims personalizados
// Profile pode ser: ADMIN, MOTORISTA, CLIENTE

type Claims struct {
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
	jwt.RegisteredClaims
}

// Gera um token JWT
func GenerateJWT(userID, name, profile string, duration time.Duration) (string, error) {
	expirationTime := time.Now().Add(duration)
	claims := &Claims{
		UserID:  userID,
		Name:    name,
		Profile: profile,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Valida e retorna os claims do token
func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("token inválido ou expirado")
	}
	return claims, nil
}
