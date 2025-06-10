package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/paulopaiva/agencia-viagens/internal/auth"
)

// Middleware para autenticação JWT
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token não informado"})
			return
		}
		tokenStr := strings.TrimPrefix(header, "Bearer ")
		claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou expirado"})
			return
		}
		// Disponibiliza os claims no contexto
		c.Set("user_id", claims.UserID)
		c.Set("user_name", claims.Name)
		c.Set("user_profile", claims.Profile)
		c.Next()
	}
}

// Middleware para autorização por perfil
func Authorize(profiles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		profile, ok := c.Get("user_profile")
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Perfil não encontrado"})
			return
		}
		for _, p := range profiles {
			if profile == p {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Acesso não autorizado"})
	}
}
