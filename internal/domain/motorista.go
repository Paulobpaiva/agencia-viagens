package domain

import (
	"time"

	"github.com/google/uuid"
)

// StatusMotorista representa os possíveis status de um motorista
type StatusMotorista string

const (
	StatusDisponivel StatusMotorista = "DISPONIVEL"
	StatusEmViagem   StatusMotorista = "EM_VIAGEM"
	StatusFolga      StatusMotorista = "FOLGA"
	StatusInativo    StatusMotorista = "INATIVO"
)

// TipoCNH representa os tipos de CNH
type TipoCNH string

const (
	CNHA TipoCNH = "A"
	CNHB TipoCNH = "B"
	CNHC TipoCNH = "C"
	CNHD TipoCNH = "D"
	CNHE TipoCNH = "E"
)

// Motorista representa um motorista no sistema
type Motorista struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Nome           string    `json:"nome" gorm:"type:varchar(100);not null"`
	CPF            string    `json:"cpf" gorm:"type:varchar(14);unique;not null"`
	RG             string    `json:"rg" gorm:"type:varchar(20);unique;not null"`
	DataNascimento time.Time `json:"data_nascimento" gorm:"not null"`

	// Contato
	Email    string `json:"email" gorm:"type:varchar(100);unique"`
	Telefone string `json:"telefone" gorm:"type:varchar(20);not null"`
	Endereco string `json:"endereco" gorm:"type:varchar(200)"`

	// Documentação
	CNH         string    `json:"cnh" gorm:"type:varchar(20);unique;not null"`
	TipoCNH     TipoCNH   `json:"tipo_cnh" gorm:"type:varchar(2);not null"`
	ValidadeCNH time.Time `json:"validade_cnh" gorm:"not null"`

	// Status e disponibilidade
	Status     StatusMotorista `json:"status" gorm:"type:varchar(20);not null;default:'DISPONIVEL'"`
	Disponivel bool            `json:"disponivel" gorm:"not null;default:true"`

	// Informações adicionais
	Observacoes string `json:"observacoes" gorm:"type:text"`
	BancoHoras  int    `json:"banco_horas" gorm:"type:int;default:0"` // em minutos

	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

// NewMotorista cria uma nova instância de Motorista
func NewMotorista(nome, cpf, rg string, dataNascimento time.Time,
	telefone, cnh string, tipoCNH TipoCNH, validadeCNH time.Time) *Motorista {
	return &Motorista{
		ID:             uuid.New(),
		Nome:           nome,
		CPF:            cpf,
		RG:             rg,
		DataNascimento: dataNascimento,
		Telefone:       telefone,
		CNH:            cnh,
		TipoCNH:        tipoCNH,
		ValidadeCNH:    validadeCNH,
		Status:         StatusDisponivel,
		Disponivel:     true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

// Validar verifica se o motorista é válido
func (m *Motorista) Validar() error {
	if m.Nome == "" {
		return ErrNomeObrigatorio
	}

	if m.CPF == "" {
		return ErrCPFObrigatorio
	}

	if m.RG == "" {
		return ErrRGObrigatorio
	}

	if m.Telefone == "" {
		return ErrTelefoneObrigatorio
	}

	if m.CNH == "" {
		return ErrCNHObrigatoria
	}

	if m.ValidadeCNH.Before(time.Now()) {
		return ErrCNHVencida
	}

	// Verifica idade mínima (18 anos)
	idadeMinima := time.Now().AddDate(-18, 0, 0)
	if m.DataNascimento.After(idadeMinima) {
		return ErrIdadeMinima
	}

	return nil
}

// AtualizarStatus atualiza o status do motorista
func (m *Motorista) AtualizarStatus(status StatusMotorista) {
	m.Status = status
	m.Disponivel = status == StatusDisponivel
	m.UpdatedAt = time.Now()
}

// AtualizarDisponibilidade atualiza a disponibilidade do motorista
func (m *Motorista) AtualizarDisponibilidade(disponivel bool) {
	m.Disponivel = disponivel
	if !disponivel {
		m.Status = StatusEmViagem
	} else {
		m.Status = StatusDisponivel
	}
	m.UpdatedAt = time.Now()
}

// AdicionarBancoHoras adiciona horas ao banco de horas do motorista
func (m *Motorista) AdicionarBancoHoras(minutos int) {
	m.BancoHoras += minutos
	m.UpdatedAt = time.Now()
}

// Erros de domínio
var (
	ErrNomeObrigatorio     = NewDomainError("nome é obrigatório")
	ErrCPFObrigatorio      = NewDomainError("CPF é obrigatório")
	ErrRGObrigatorio       = NewDomainError("RG é obrigatório")
	ErrTelefoneObrigatorio = NewDomainError("telefone é obrigatório")
	ErrCNHObrigatoria      = NewDomainError("CNH é obrigatória")
	ErrCNHVencida          = NewDomainError("CNH está vencida")
	ErrIdadeMinima         = NewDomainError("motorista deve ter no mínimo 18 anos")
)
