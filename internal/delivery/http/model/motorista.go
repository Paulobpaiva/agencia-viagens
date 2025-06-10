package model

import (
	"agencia-viagens/internal/domain"
	"agencia-viagens/internal/validator"
	"time"
)

// CreateMotoristaRequest representa a requisição de criação de motorista
type CreateMotoristaRequest struct {
	Nome           string              `json:"nome" binding:"required"`
	CPF            string              `json:"cpf" binding:"required"`
	CNH            string              `json:"cnh" binding:"required"`
	ValidadeCNH    time.Time           `json:"validade_cnh" binding:"required"`
	CategoriaCNH   domain.CategoriaCNH `json:"categoria_cnh" binding:"required"`
	DataNascimento time.Time           `json:"data_nascimento" binding:"required"`
	Email          string              `json:"email" binding:"required,email"`
	Telefone       string              `json:"telefone" binding:"required"`
	Endereco       string              `json:"endereco" binding:"required"`
	Observacoes    string              `json:"observacoes"`
}

// Validate implementa a interface Validator
func (r *CreateMotoristaRequest) Validate() error {
	if err := validator.ValidarCPF(r.CPF); err != nil {
		return err
	}

	if err := validator.ValidarCNH(r.CNH); err != nil {
		return err
	}

	if err := validator.ValidarCategoriaCNH(r.CategoriaCNH); err != nil {
		return err
	}

	if err := validator.ValidarEmail(r.Email); err != nil {
		return err
	}

	if err := validator.ValidarTelefone(r.Telefone); err != nil {
		return err
	}

	if err := validator.ValidarDataNascimento(r.DataNascimento); err != nil {
		return err
	}

	return nil
}

// UpdateMotoristaRequest representa a requisição de atualização de motorista
type UpdateMotoristaRequest struct {
	Nome         string                 `json:"nome"`
	ValidadeCNH  time.Time              `json:"validade_cnh"`
	CategoriaCNH domain.CategoriaCNH    `json:"categoria_cnh"`
	Email        string                 `json:"email"`
	Telefone     string                 `json:"telefone"`
	Endereco     string                 `json:"endereco"`
	Observacoes  string                 `json:"observacoes"`
	Status       domain.StatusMotorista `json:"status"`
}

// Validate implementa a interface Validator
func (r *UpdateMotoristaRequest) Validate() error {
	if r.Email != "" {
		if err := validator.ValidarEmail(r.Email); err != nil {
			return err
		}
	}

	if r.Telefone != "" {
		if err := validator.ValidarTelefone(r.Telefone); err != nil {
			return err
		}
	}

	if r.CategoriaCNH != "" {
		if err := validator.ValidarCategoriaCNH(r.CategoriaCNH); err != nil {
			return err
		}
	}

	if r.Status != "" {
		if err := validator.ValidarStatusMotorista(r.Status); err != nil {
			return err
		}
	}

	return nil
}

// MotoristaResponse representa a resposta de motorista
type MotoristaResponse struct {
	ID                 string                 `json:"id"`
	Nome               string                 `json:"nome"`
	CPF                string                 `json:"cpf"`
	CNH                string                 `json:"cnh"`
	ValidadeCNH        time.Time              `json:"validade_cnh"`
	CategoriaCNH       domain.CategoriaCNH    `json:"categoria_cnh"`
	DataNascimento     time.Time              `json:"data_nascimento"`
	Email              string                 `json:"email"`
	Telefone           string                 `json:"telefone"`
	Endereco           string                 `json:"endereco"`
	Status             domain.StatusMotorista `json:"status"`
	Observacoes        string                 `json:"observacoes"`
	DocumentacaoValida bool                   `json:"documentacao_valida"`
	UltimaAvaliacao    time.Time              `json:"ultima_avaliacao"`
	ProximaAvaliacao   time.Time              `json:"proxima_avaliacao"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
}

// NewMotoristaResponse cria uma nova resposta de motorista
func NewMotoristaResponse(m *domain.Motorista) *MotoristaResponse {
	return &MotoristaResponse{
		ID:                 m.ID.String(),
		Nome:               m.Nome,
		CPF:                m.CPF,
		CNH:                m.CNH,
		ValidadeCNH:        m.ValidadeCNH,
		CategoriaCNH:       m.CategoriaCNH,
		DataNascimento:     m.DataNascimento,
		Email:              m.Email,
		Telefone:           m.Telefone,
		Endereco:           m.Endereco,
		Status:             m.Status,
		Observacoes:        m.Observacoes,
		DocumentacaoValida: m.DocumentacaoValida,
		UltimaAvaliacao:    m.UltimaAvaliacao,
		ProximaAvaliacao:   m.ProximaAvaliacao,
		CreatedAt:          m.CreatedAt,
		UpdatedAt:          m.UpdatedAt,
	}
}

// ListMotoristasResponse representa a resposta de listagem de motoristas
type ListMotoristasResponse struct {
	Motoristas []*MotoristaResponse `json:"motoristas"`
	Total      int64                `json:"total"`
}

// NewListMotoristasResponse cria uma nova resposta de listagem de motoristas
func NewListMotoristasResponse(motoristas []*domain.Motorista, total int64) *ListMotoristasResponse {
	response := &ListMotoristasResponse{
		Motoristas: make([]*MotoristaResponse, len(motoristas)),
		Total:      total,
	}

	for i, m := range motoristas {
		response.Motoristas[i] = NewMotoristaResponse(m)
	}

	return response
}

// MotoristaQueryParams representa os parâmetros de query para listagem de motoristas
type MotoristaQueryParams struct {
	Offset       int    `form:"offset" binding:"min=0"`
	Limit        int    `form:"limit" binding:"min=1,max=100"`
	Status       string `form:"status"`
	CategoriaCNH string `form:"categoria_cnh"`
}

// Validate implementa a interface Validator
func (p *MotoristaQueryParams) Validate() error {
	if p.Status != "" {
		if err := validator.ValidarStatusMotorista(domain.StatusMotorista(p.Status)); err != nil {
			return err
		}
	}

	if p.CategoriaCNH != "" {
		if err := validator.ValidarCategoriaCNH(domain.CategoriaCNH(p.CategoriaCNH)); err != nil {
			return err
		}
	}

	return nil
}

// DisponibilidadeMotoristaQueryParams representa os parâmetros de query para verificação de disponibilidade
type DisponibilidadeMotoristaQueryParams struct {
	DataInicio time.Time `form:"data_inicio" binding:"required"`
	DataFim    time.Time `form:"data_fim" binding:"required"`
}

// Validate implementa a interface Validator
func (p *DisponibilidadeMotoristaQueryParams) Validate() error {
	return validator.ValidarPeriodo(p.DataInicio, p.DataFim)
}
