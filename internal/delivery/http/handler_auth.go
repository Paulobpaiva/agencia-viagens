package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paulopaiva/agencia-viagens/internal/auth"
)

// LoginRequest representa os dados necessários para autenticação
type LoginRequest struct {
	CPF    string `json:"cpf" binding:"required" example:"12345678900"`                            // CPF do usuário
	Senha  string `json:"senha" binding:"required" example:"senha123"`                             // Senha do usuário
	Perfil string `json:"perfil" binding:"required" example:"MOTORISTA" enums:"MOTORISTA,CLIENTE"` // Perfil do usuário
}

// LoginResponse representa a resposta do endpoint de login
type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."` // Token JWT
}

// @Summary      Autentica um usuário
// @Description  Realiza a autenticação de um usuário (motorista ou cliente) e retorna um token JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "Dados de login"
// @Success      200 {object} LoginResponse
// @Failure      400 {object} map[string]string "Dados inválidos"
// @Failure      401 {object} map[string]string "Credenciais inválidas"
// @Failure      500 {object} map[string]string "Erro interno"
// @Router       /auth/login [post]
func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Simulação: aceita qualquer senha para CPF válido
	if req.Perfil != "MOTORISTA" && req.Perfil != "CLIENTE" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Perfil inválido"})
		return
	}

	// Aqui você faria a validação real no banco de dados
	// Exemplo: buscar motorista/cliente pelo CPF e comparar senha (hash)
	userID := req.CPF
	name := "Usuário Exemplo"
	profile := req.Perfil

	token, err := auth.GenerateJWT(userID, name, profile, 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}
