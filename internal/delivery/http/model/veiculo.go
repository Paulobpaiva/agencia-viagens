package model

import (
	"agencia-viagens/internal/domain"
	"agencia-viagens/internal/validator"
	"time"
)

// CreateVeiculoRequest representa a requisição de criação de veículo
type CreateVeiculoRequest struct {
	Placa       string             `json:"placa" binding:"required"`
	Modelo      string             `json:"modelo" binding:"required"`
	Marca       string             `json:"marca" binding:"required"`
	Ano         int                `json:"ano" binding:"required"`
	Tipo        domain.TipoVeiculo `json:"tipo" binding:"required"`
	Capacidade  int                `json:"capacidade" binding:"required"`
	Chassi      string             `json:"chassi" binding:"required"`
	Renavam     string             `json:"renavam" binding:"required"`
	Cor         string             `json:"cor"`
	Observacoes string             `json:"observacoes"`
}

// Validate implementa a interface Validator
func (r *CreateVeiculoRequest) Validate() error {
	if err := validator.ValidarPlaca(r.Placa); err != nil {
		return err
	}

	if err := validator.ValidarAno(r.Ano); err != nil {
		return err
	}

	if err := validator.ValidarTipoVeiculo(r.Tipo); err != nil {
		return err
	}

	if err := validator.ValidarCapacidade(r.Capacidade); err != nil {
		return err
	}

	return nil
}

// UpdateVeiculoRequest representa a requisição de atualização de veículo
type UpdateVeiculoRequest struct {
	Modelo      string               `json:"modelo"`
	Marca       string               `json:"marca"`
	Ano         int                  `json:"ano"`
	Tipo        domain.TipoVeiculo   `json:"tipo"`
	Capacidade  int                  `json:"capacidade"`
	Cor         string               `json:"cor"`
	Observacoes string               `json:"observacoes"`
	Status      domain.StatusVeiculo `json:"status"`
}

// Validate implementa a interface Validator
func (r *UpdateVeiculoRequest) Validate() error {
	if r.Ano != 0 {
		if err := validator.ValidarAno(r.Ano); err != nil {
			return err
		}
	}

	if r.Tipo != "" {
		if err := validator.ValidarTipoVeiculo(r.Tipo); err != nil {
			return err
		}
	}

	if r.Capacidade != 0 {
		if err := validator.ValidarCapacidade(r.Capacidade); err != nil {
			return err
		}
	}

	if r.Status != "" {
		if err := validator.ValidarStatusVeiculo(r.Status); err != nil {
			return err
		}
	}

	return nil
}

// VeiculoResponse representa a resposta de veículo
type VeiculoResponse struct {
	ID                     string               `json:"id"`
	Placa                  string               `json:"placa"`
	Modelo                 string               `json:"modelo"`
	Marca                  string               `json:"marca"`
	Ano                    int                  `json:"ano"`
	Tipo                   domain.TipoVeiculo   `json:"tipo"`
	Capacidade             int                  `json:"capacidade"`
	Status                 domain.StatusVeiculo `json:"status"`
	Chassi                 string               `json:"chassi"`
	Renavam                string               `json:"renavam"`
	Cor                    string               `json:"cor"`
	Observacoes            string               `json:"observacoes"`
	DocumentacaoValida     bool                 `json:"documentacao_valida"`
	VencimentoDocumentacao time.Time            `json:"vencimento_documentacao"`
	UltimaManutencao       time.Time            `json:"ultima_manutencao"`
	ProximaManutencao      time.Time            `json:"proxima_manutencao"`
	CreatedAt              time.Time            `json:"created_at"`
	UpdatedAt              time.Time            `json:"updated_at"`
}

// NewVeiculoResponse cria uma nova resposta de veículo
func NewVeiculoResponse(v *domain.Veiculo) *VeiculoResponse {
	return &VeiculoResponse{
		ID:                     v.ID.String(),
		Placa:                  v.Placa,
		Modelo:                 v.Modelo,
		Marca:                  v.Marca,
		Ano:                    v.Ano,
		Tipo:                   v.Tipo,
		Capacidade:             v.Capacidade,
		Status:                 v.Status,
		Chassi:                 v.Chassi,
		Renavam:                v.Renavam,
		Cor:                    v.Cor,
		Observacoes:            v.Observacoes,
		DocumentacaoValida:     v.DocumentacaoValida,
		VencimentoDocumentacao: v.VencimentoDocumentacao,
		UltimaManutencao:       v.UltimaManutencao,
		ProximaManutencao:      v.ProximaManutencao,
		CreatedAt:              v.CreatedAt,
		UpdatedAt:              v.UpdatedAt,
	}
}

// ListVeiculosResponse representa a resposta de listagem de veículos
type ListVeiculosResponse struct {
	Veiculos []*VeiculoResponse `json:"veiculos"`
	Total    int64              `json:"total"`
}

// NewListVeiculosResponse cria uma nova resposta de listagem de veículos
func NewListVeiculosResponse(veiculos []*domain.Veiculo, total int64) *ListVeiculosResponse {
	response := &ListVeiculosResponse{
		Veiculos: make([]*VeiculoResponse, len(veiculos)),
		Total:    total,
	}

	for i, v := range veiculos {
		response.Veiculos[i] = NewVeiculoResponse(v)
	}

	return response
}

// VeiculoQueryParams representa os parâmetros de query para listagem de veículos
type VeiculoQueryParams struct {
	Offset int    `form:"offset" binding:"min=0"`
	Limit  int    `form:"limit" binding:"min=1,max=100"`
	Status string `form:"status"`
	Tipo   string `form:"tipo"`
}

// Validate implementa a interface Validator
func (p *VeiculoQueryParams) Validate() error {
	if p.Status != "" {
		if err := validator.ValidarStatusVeiculo(domain.StatusVeiculo(p.Status)); err != nil {
			return err
		}
	}

	if p.Tipo != "" {
		if err := validator.ValidarTipoVeiculo(domain.TipoVeiculo(p.Tipo)); err != nil {
			return err
		}
	}

	return nil
}

// DisponibilidadeVeiculoQueryParams representa os parâmetros de query para verificação de disponibilidade
type DisponibilidadeVeiculoQueryParams struct {
	DataInicio time.Time `form:"data_inicio" binding:"required"`
	DataFim    time.Time `form:"data_fim" binding:"required"`
}

// Validate implementa a interface Validator
func (p *DisponibilidadeVeiculoQueryParams) Validate() error {
	return validator.ValidarPeriodo(p.DataInicio, p.DataFim)
}
