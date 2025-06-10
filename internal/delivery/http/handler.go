package http

import (
	"net/http"

	"agencia-viagens/internal/delivery/http/middleware"
	"agencia-viagens/internal/domain"
	"agencia-viagens/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	viagemUseCase    *usecase.ViagemUseCase
	veiculoUseCase   *usecase.VeiculoUseCase
	motoristaUseCase *usecase.MotoristaUseCase
}

func NewHandler(
	viagemUseCase *usecase.ViagemUseCase,
	veiculoUseCase *usecase.VeiculoUseCase,
	motoristaUseCase *usecase.MotoristaUseCase,
) *Handler {
	return &Handler{
		viagemUseCase:    viagemUseCase,
		veiculoUseCase:   veiculoUseCase,
		motoristaUseCase: motoristaUseCase,
	}
}

func (h *Handler) InitRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")

	// Rota de login
	api.POST("/auth/login", LoginHandler)

	// Rotas de Viagens
	viagens := api.Group("/viagens")
	{
		viagens.POST("", h.CriarViagem)
		viagens.GET("", h.ListarViagens)
		viagens.GET("/:id", h.BuscarViagem)
		viagens.PUT("/:id", h.AtualizarViagem)
		viagens.DELETE("/:id", h.CancelarViagem)
	}

	// Rotas de Veículos
	veiculos := api.Group("/veiculos")
	{
		veiculos.POST("", h.CriarVeiculo)
		veiculos.GET("/:id", h.BuscarVeiculo)
		veiculos.PUT("/:id", h.AtualizarVeiculo)
		veiculos.DELETE("/:id", h.RemoverVeiculo)
		veiculos.GET("/", middleware.AuthRequired(), h.ListarVeiculos)
	}

	// Rotas de Motoristas
	motoristas := api.Group("/motoristas")
	{
		motoristas.POST("", h.CriarMotorista)
		motoristas.GET("", h.ListarMotoristas)
		motoristas.GET("/:id", h.BuscarMotorista)
		motoristas.PUT("/:id", h.AtualizarMotorista)
		motoristas.DELETE("/:id", h.RemoverMotorista)
	}
}

// Handlers de Viagem
func (h *Handler) CriarViagem(c *gin.Context) {
	var viagem domain.Viagem
	if err := c.ShouldBindJSON(&viagem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.viagemUseCase.Criar(c.Request.Context(), &viagem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, viagem)
}

func (h *Handler) ListarViagens(c *gin.Context) {
	viagens, err := h.viagemUseCase.Listar(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, viagens)
}

func (h *Handler) BuscarViagem(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	viagem, err := h.viagemUseCase.BuscarPorID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Viagem não encontrada"})
		return
	}

	c.JSON(http.StatusOK, viagem)
}

func (h *Handler) AtualizarViagem(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var viagem domain.Viagem
	if err := c.ShouldBindJSON(&viagem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	viagem.ID = id
	if err := h.viagemUseCase.Atualizar(c.Request.Context(), &viagem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, viagem)
}

func (h *Handler) CancelarViagem(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.viagemUseCase.Cancelar(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// Handlers de Veículo
// @Summary      Lista todos os veículos
// @Description  Retorna a lista de veículos cadastrados
// @Tags         veiculos
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array}  domain.Veiculo
// @Failure      500 {object} map[string]string "Erro interno"
// @Router       /veiculos [get]
func (h *Handler) ListarVeiculos(c *gin.Context) {
	veiculos, err := h.veiculoUseCase.Listar(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, veiculos)
}

// @Summary      Busca um veículo pelo ID
// @Description  Retorna os detalhes de um veículo específico
// @Tags         veiculos
// @Accept       json
// @Produce      json
// @Param        id path string true "ID do veículo" format(uuid)
// @Success      200 {object} domain.Veiculo
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      404 {object} map[string]string "Veículo não encontrado"
// @Failure      500 {object} map[string]string "Erro interno"
// @Router       /veiculos/{id} [get]
func (h *Handler) BuscarVeiculo(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	veiculo, err := h.veiculoUseCase.BuscarPorID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Veículo não encontrado"})
		return
	}

	c.JSON(http.StatusOK, veiculo)
}

// @Summary      Cria um novo veículo
// @Description  Cadastra um novo veículo no sistema
// @Tags         veiculos
// @Accept       json
// @Produce      json
// @Param        veiculo body domain.Veiculo true "Dados do veículo"
// @Success      201 {object} domain.Veiculo
// @Failure      400 {object} map[string]string "Dados inválidos"
// @Failure      500 {object} map[string]string "Erro interno"
// @Router       /veiculos [post]
func (h *Handler) CriarVeiculo(c *gin.Context) {
	var veiculo domain.Veiculo
	if err := c.ShouldBindJSON(&veiculo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.veiculoUseCase.Criar(c.Request.Context(), &veiculo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, veiculo)
}

// @Summary      Atualiza um veículo
// @Description  Atualiza os dados de um veículo existente
// @Tags         veiculos
// @Accept       json
// @Produce      json
// @Param        id path string true "ID do veículo" format(uuid)
// @Param        veiculo body domain.Veiculo true "Dados do veículo"
// @Success      200 {object} domain.Veiculo
// @Failure      400 {object} map[string]string "Dados inválidos"
// @Failure      404 {object} map[string]string "Veículo não encontrado"
// @Failure      500 {object} map[string]string "Erro interno"
// @Router       /veiculos/{id} [put]
func (h *Handler) AtualizarVeiculo(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var veiculo domain.Veiculo
	if err := c.ShouldBindJSON(&veiculo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	veiculo.ID = id
	if err := h.veiculoUseCase.Atualizar(c.Request.Context(), &veiculo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, veiculo)
}

// @Summary      Remove um veículo
// @Description  Remove um veículo do sistema
// @Tags         veiculos
// @Accept       json
// @Produce      json
// @Param        id path string true "ID do veículo" format(uuid)
// @Success      204 "No Content"
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      404 {object} map[string]string "Veículo não encontrado"
// @Failure      500 {object} map[string]string "Erro interno"
// @Router       /veiculos/{id} [delete]
func (h *Handler) RemoverVeiculo(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.veiculoUseCase.Remover(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// Handlers de Motorista
func (h *Handler) CriarMotorista(c *gin.Context) {
	var motorista domain.Motorista
	if err := c.ShouldBindJSON(&motorista); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.motoristaUseCase.Criar(c.Request.Context(), &motorista); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, motorista)
}

func (h *Handler) ListarMotoristas(c *gin.Context) {
	motoristas, err := h.motoristaUseCase.Listar(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, motoristas)
}

func (h *Handler) BuscarMotorista(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	motorista, err := h.motoristaUseCase.BuscarPorID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Motorista não encontrado"})
		return
	}

	c.JSON(http.StatusOK, motorista)
}

func (h *Handler) AtualizarMotorista(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var motorista domain.Motorista
	if err := c.ShouldBindJSON(&motorista); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	motorista.ID = id
	if err := h.motoristaUseCase.Atualizar(c.Request.Context(), &motorista); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, motorista)
}

func (h *Handler) RemoverMotorista(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.motoristaUseCase.Remover(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
