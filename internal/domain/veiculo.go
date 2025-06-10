package domain

import (
	"time"

	"github.com/google/uuid"
)

// StatusVeiculo representa os possíveis status de um veículo
type StatusVeiculo string

const (
	StatusDisponivel StatusVeiculo = "DISPONIVEL"
	StatusEmUso      StatusVeiculo = "EM_USO"
	StatusManutencao StatusVeiculo = "MANUTENCAO"
	StatusInativo    StatusVeiculo = "INATIVO"
)

// TipoVeiculo representa os tipos de veículos disponíveis
type TipoVeiculo string

const (
	TipoVan         TipoVeiculo = "VAN"
	TipoOnibus      TipoVeiculo = "ONIBUS"
	TipoMicroOnibus TipoVeiculo = "MICRO_ONIBUS"
)

// Veiculo representa um veículo no sistema
type Veiculo struct {
	ID         uuid.UUID     `json:"id" gorm:"type:uuid;primary_key"`
	Placa      string        `json:"placa" gorm:"type:varchar(8);unique;not null"`
	Modelo     string        `json:"modelo" gorm:"type:varchar(100);not null"`
	Marca      string        `json:"marca" gorm:"type:varchar(100);not null"`
	Ano        int           `json:"ano" gorm:"not null"`
	Tipo       TipoVeiculo   `json:"tipo" gorm:"type:varchar(20);not null"`
	Capacidade int           `json:"capacidade" gorm:"not null"`
	Status     StatusVeiculo `json:"status" gorm:"type:varchar(20);not null;default:'DISPONIVEL'"`

	// Informações adicionais
	Chassi      string `json:"chassi" gorm:"type:varchar(50);unique"`
	Renavam     string `json:"renavam" gorm:"type:varchar(50);unique"`
	Cor         string `json:"cor" gorm:"type:varchar(50)"`
	Observacoes string `json:"observacoes" gorm:"type:text"`

	// Documentação
	DocumentacaoValida     bool      `json:"documentacao_valida" gorm:"not null;default:true"`
	VencimentoDocumentacao time.Time `json:"vencimento_documentacao" gorm:"not null"`

	// Manutenção
	UltimaManutencao  time.Time `json:"ultima_manutencao"`
	ProximaManutencao time.Time `json:"proxima_manutencao"`

	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

// NewVeiculo cria uma nova instância de Veículo
func NewVeiculo(placa, modelo, marca string, ano int, tipo TipoVeiculo,
	capacidade int, chassi, renavam string) *Veiculo {
	return &Veiculo{
		ID:                     uuid.New(),
		Placa:                  placa,
		Modelo:                 modelo,
		Marca:                  marca,
		Ano:                    ano,
		Tipo:                   tipo,
		Capacidade:             capacidade,
		Status:                 StatusDisponivel,
		Chassi:                 chassi,
		Renavam:                renavam,
		DocumentacaoValida:     true,
		VencimentoDocumentacao: time.Now().AddDate(1, 0, 0), // 1 ano
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}
}

// Validar verifica se o veículo é válido
func (v *Veiculo) Validar() error {
	if v.Placa == "" {
		return ErrPlacaObrigatoria
	}

	if v.Modelo == "" || v.Marca == "" {
		return ErrModeloMarcaObrigatorios
	}

	if v.Ano < 1900 || v.Ano > time.Now().Year() {
		return ErrAnoInvalido
	}

	if v.Capacidade <= 0 {
		return ErrCapacidadeInvalida
	}

	if v.Chassi == "" || v.Renavam == "" {
		return ErrChassiRenavamObrigatorios
	}

	return nil
}

// AtualizarStatus atualiza o status do veículo
func (v *Veiculo) AtualizarStatus(status StatusVeiculo) {
	v.Status = status
	v.UpdatedAt = time.Now()
}

// AtualizarDocumentacao atualiza as informações de documentação
func (v *Veiculo) AtualizarDocumentacao(valida bool, vencimento time.Time) {
	v.DocumentacaoValida = valida
	v.VencimentoDocumentacao = vencimento
	v.UpdatedAt = time.Now()
}

// RegistrarManutencao registra uma manutenção realizada
func (v *Veiculo) RegistrarManutencao() {
	v.UltimaManutencao = time.Now()
	v.ProximaManutencao = time.Now().AddDate(0, 6, 0) // 6 meses
	v.Status = StatusDisponivel
	v.UpdatedAt = time.Now()
}

// Erros de domínio
var (
	ErrPlacaObrigatoria          = NewDomainError("placa é obrigatória")
	ErrModeloMarcaObrigatorios   = NewDomainError("modelo e marca são obrigatórios")
	ErrAnoInvalido               = NewDomainError("ano inválido")
	ErrCapacidadeInvalida        = NewDomainError("capacidade deve ser maior que zero")
	ErrChassiRenavamObrigatorios = NewDomainError("chassi e renavam são obrigatórios")
)
