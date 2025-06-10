package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/paulopaiva/agencia-viagens/internal/domain"
	"github.com/paulopaiva/agencia-viagens/internal/validator"
)

// CreateViagemRequest representa a requisição de criação de viagem
type CreateViagemRequest struct {
	VeiculoID   uuid.UUID           `json:"veiculo_id" binding:"required"`
	MotoristaID uuid.UUID           `json:"motorista_id" binding:"required"`
	ClienteID   uuid.UUID           `json:"cliente_id" binding:"required"`
	Origem      string              `json:"origem" binding:"required"`
	Destino     string              `json:"destino" binding:"required"`
	DataInicio  time.Time           `json:"data_inicio" binding:"required"`
	DataFim     time.Time           `json:"data_fim" binding:"required"`
	Valor       float64             `json:"valor" binding:"required"`
	Status      domain.StatusViagem `json:"status" binding:"required"`
	Observacoes string              `json:"observacoes"`
}

// Validate implementa a interface Validator
func (r *CreateViagemRequest) Validate() error {
	if err := validator.ValidarPeriodo(r.DataInicio, r.DataFim); err != nil {
		return err
	}

	if err := validator.ValidarStatusViagem(r.Status); err != nil {
		return err
	}

	if err := validator.ValidarValor(r.Valor); err != nil {
		return err
	}

	return nil
}

// UpdateViagemRequest representa a requisição de atualização de viagem
type UpdateViagemRequest struct {
	Origem      string              `json:"origem"`
	Destino     string              `json:"destino"`
	DataInicio  time.Time           `json:"data_inicio"`
	DataFim     time.Time           `json:"data_fim"`
	Valor       float64             `json:"valor"`
	Status      domain.StatusViagem `json:"status"`
	Observacoes string              `json:"observacoes"`
}

// Validate implementa a interface Validator
func (r *UpdateViagemRequest) Validate() error {
	if !r.DataInicio.IsZero() && !r.DataFim.IsZero() {
		if err := validator.ValidarPeriodo(r.DataInicio, r.DataFim); err != nil {
			return err
		}
	}

	if r.Status != "" {
		if err := validator.ValidarStatusViagem(r.Status); err != nil {
			return err
		}
	}

	if r.Valor != 0 {
		if err := validator.ValidarValor(r.Valor); err != nil {
			return err
		}
	}

	return nil
}

// ViagemResponse representa a resposta de viagem
type ViagemResponse struct {
	ID          string              `json:"id"`
	VeiculoID   string              `json:"veiculo_id"`
	MotoristaID string              `json:"motorista_id"`
	ClienteID   string              `json:"cliente_id"`
	Origem      string              `json:"origem"`
	Destino     string              `json:"destino"`
	DataInicio  time.Time           `json:"data_inicio"`
	DataFim     time.Time           `json:"data_fim"`
	Valor       float64             `json:"valor"`
	Status      domain.StatusViagem `json:"status"`
	Observacoes string              `json:"observacoes"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

// NewViagemResponse cria uma nova resposta de viagem
func NewViagemResponse(v *domain.Viagem) *ViagemResponse {
	return &ViagemResponse{
		ID:          v.ID.String(),
		VeiculoID:   v.VeiculoID.String(),
		MotoristaID: v.MotoristaID.String(),
		ClienteID:   v.ClienteID.String(),
		Origem:      v.Origem,
		Destino:     v.Destino,
		DataInicio:  v.DataInicio,
		DataFim:     v.DataFim,
		Valor:       v.Valor,
		Status:      v.Status,
		Observacoes: v.Observacoes,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}
}

// ListViagensResponse representa a resposta de listagem de viagens
type ListViagensResponse struct {
	Viagens []*ViagemResponse `json:"viagens"`
	Total   int64             `json:"total"`
}

// NewListViagensResponse cria uma nova resposta de listagem de viagens
func NewListViagensResponse(viagens []*domain.Viagem, total int64) *ListViagensResponse {
	response := &ListViagensResponse{
		Viagens: make([]*ViagemResponse, len(viagens)),
		Total:   total,
	}

	for i, v := range viagens {
		response.Viagens[i] = NewViagemResponse(v)
	}

	return response
}

// ViagemQueryParams representa os parâmetros de query para listagem de viagens
type ViagemQueryParams struct {
	Offset      int       `form:"offset" binding:"min=0"`
	Limit       int       `form:"limit" binding:"min=1,max=100"`
	Status      string    `form:"status"`
	DataInicio  time.Time `form:"data_inicio"`
	DataFim     time.Time `form:"data_fim"`
	VeiculoID   string    `form:"veiculo_id"`
	MotoristaID string    `form:"motorista_id"`
	ClienteID   string    `form:"cliente_id"`
}

// Validate implementa a interface Validator
func (p *ViagemQueryParams) Validate() error {
	if p.Status != "" {
		if err := validator.ValidarStatusViagem(domain.StatusViagem(p.Status)); err != nil {
			return err
		}
	}

	if !p.DataInicio.IsZero() && !p.DataFim.IsZero() {
		if err := validator.ValidarPeriodo(p.DataInicio, p.DataFim); err != nil {
			return err
		}
	}

	if p.VeiculoID != "" {
		if _, err := uuid.Parse(p.VeiculoID); err != nil {
			return validator.ErrIDInvalido
		}
	}

	if p.MotoristaID != "" {
		if _, err := uuid.Parse(p.MotoristaID); err != nil {
			return validator.ErrIDInvalido
		}
	}

	if p.ClienteID != "" {
		if _, err := uuid.Parse(p.ClienteID); err != nil {
			return validator.ErrIDInvalido
		}
	}

	return nil
}

// DisponibilidadeViagemQueryParams representa os parâmetros de query para verificação de disponibilidade
type DisponibilidadeViagemQueryParams struct {
	DataInicio  time.Time `form:"data_inicio" binding:"required"`
	DataFim     time.Time `form:"data_fim" binding:"required"`
	VeiculoID   string    `form:"veiculo_id"`
	MotoristaID string    `form:"motorista_id"`
}

// Validate implementa a interface Validator
func (p *DisponibilidadeViagemQueryParams) Validate() error {
	if err := validator.ValidarPeriodo(p.DataInicio, p.DataFim); err != nil {
		return err
	}

	if p.VeiculoID != "" {
		if _, err := uuid.Parse(p.VeiculoID); err != nil {
			return validator.ErrIDInvalido
		}
	}

	if p.MotoristaID != "" {
		if _, err := uuid.Parse(p.MotoristaID); err != nil {
			return validator.ErrIDInvalido
		}
	}

	return nil
}
