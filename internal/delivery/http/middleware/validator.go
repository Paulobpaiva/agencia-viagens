package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paulopaiva/agencia-viagens/internal/validator"
)

// Validator é uma interface para validar objetos
type Validator interface {
	Validate() error
}

// ValidateRequest é um middleware que valida o corpo da requisição
func ValidateRequest(v Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(v); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Erro ao decodificar requisição: " + err.Error(),
			})
			c.Abort()
			return
		}

		if err := v.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Erro de validação: " + err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("validated_request", v)
		c.Next()
	}
}

// ValidateQueryParams é um middleware que valida os parâmetros de query
func ValidateQueryParams(v Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindQuery(v); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Erro ao validar parâmetros de query: " + err.Error(),
			})
			c.Abort()
			return
		}

		if err := v.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Erro de validação: " + err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("validated_query", v)
		c.Next()
	}
}

// ValidatePathParams é um middleware que valida os parâmetros de path
func ValidatePathParams(v Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindUri(v); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Erro ao validar parâmetros de path: " + err.Error(),
			})
			c.Abort()
			return
		}

		if err := v.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Erro de validação: " + err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("validated_path", v)
		c.Next()
	}
}

// ErrorResponse representa uma resposta de erro padronizada
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewErrorResponse cria uma nova resposta de erro
func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error: err.Error(),
	}
}

// RespondWithError envia uma resposta de erro
func RespondWithError(c *gin.Context, status int, err error) {
	c.JSON(status, NewErrorResponse(err))
}

// RespondWithValidationError envia uma resposta de erro de validação
func RespondWithValidationError(c *gin.Context, err error) {
	RespondWithError(c, http.StatusBadRequest, err)
}

// RespondWithJSON envia uma resposta JSON
func RespondWithJSON(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

// GetValidatedRequest retorna o objeto validado da requisição
func GetValidatedRequest(c *gin.Context, v Validator) error {
	validated, exists := c.Get("validated_request")
	if !exists {
		return validator.ErrDataInvalida
	}

	data, err := json.Marshal(validated)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

// GetValidatedQuery retorna o objeto validado dos parâmetros de query
func GetValidatedQuery(c *gin.Context, v Validator) error {
	validated, exists := c.Get("validated_query")
	if !exists {
		return validator.ErrDataInvalida
	}

	data, err := json.Marshal(validated)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

// GetValidatedPath retorna o objeto validado dos parâmetros de path
func GetValidatedPath(c *gin.Context, v Validator) error {
	validated, exists := c.Get("validated_path")
	if !exists {
		return validator.ErrDataInvalida
	}

	data, err := json.Marshal(validated)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
