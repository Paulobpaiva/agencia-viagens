package domain

import (
	"time"

	"github.com/google/uuid"
)

// StatusViagem representa os possíveis status de uma viagem
type StatusViagem string

const (
	StatusAgendada   StatusViagem = "AGENDADA"
	StatusEmAndamento StatusViagem = "EM_ANDAMENTO"
	StatusConcluida  StatusViagem = "CONCLUIDA"
	StatusCancelada  StatusViagem = "CANCELADA"
)

// Viagem representa uma viagem no sistema
type Viagem struct {
	ID          uuid.UUID   `json:"id" gorm:"type:uuid;primary_key"`
	VeiculoID   uuid.UUID   `json:"veiculo_id" gorm:"type:uuid;not null"`
	MotoristaID uuid.UUID   `json:"motorista_id" gorm:"type:uuid;not null"`
	ClienteID   uuid.UUID   `json:"cliente_id" gorm:"type:uuid;not null"`
	
	Origem      string      `json:"origem" gorm:"not null"`
	Destino     string      `json:"destino" gorm:"not null"`
	DataInicio  time.Time   `json:"data_inicio" gorm:"not null"`
	DataFim     time.Time   `json:"data_fim" gorm:"not null"`
	
	Status      StatusViagem `json:"status" gorm:"type:varchar(20);not null;default:'AGENDADA'"`
	Valor       float64     `json:"valor" gorm:"type:decimal(10,2);not null"`
	Observacoes string      `json:"observacoes" gorm:"type:text"`
	
	// Coordenadas da rota
	CoordenadasOrigem  string `json:"coordenadas_origem" gorm:"type:varchar(100)"`
	CoordenadasDestino string `json:"coordenadas_destino" gorm:"type:varchar(100)"`
	RotaCompleta      string `json:"rota_completa" gorm:"type:text"`
	
	// Relacionamentos
	Veiculo    *Veiculo    `json:"veiculo,omitempty" gorm:"foreignKey:VeiculoID"`
	Motorista  *Motorista  `json:"motorista,omitempty" gorm:"foreignKey:MotoristaID"`
	Cliente    *Cliente    `json:"cliente,omitempty" gorm:"foreignKey:ClienteID"`
	
	CreatedAt  time.Time   `json:"created_at" gorm:"not null"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"not null"`
}

// NewViagem cria uma nova instância de Viagem
func NewViagem(veiculoID, motoristaID, clienteID uuid.UUID, origem, destino string, 
	dataInicio, dataFim time.Time, valor float64) *Viagem {
	return &Viagem{
		ID:          uuid.New(),
		VeiculoID:   veiculoID,
		MotoristaID: motoristaID,
		ClienteID:   clienteID,
		Origem:      origem,
		Destino:     destino,
		DataInicio:  dataInicio,
		DataFim:     dataFim,
		Status:      StatusAgendada,
		Valor:       valor,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// Validar verifica se a viagem é válida
func (v *Viagem) Validar() error {
	if v.DataInicio.After(v.DataFim) {
		return ErrDataInicioMaiorQueFim
	}
	
	if v.DataInicio.Before(time.Now()) {
		return ErrDataInicioPassada
	}
	
	if v.Valor <= 0 {
		return ErrValorInvalido
	}
	
	if v.Origem == "" || v.Destino == "" {
		return ErrOrigemDestinoObrigatorios
	}
	
	return nil
}

// AtualizarStatus atualiza o status da viagem
func (v *Viagem) AtualizarStatus(status StatusViagem) {
	v.Status = status
	v.UpdatedAt = time.Now()
}

// AtualizarRota atualiza as informações da rota
func (v *Viagem) AtualizarRota(coordsOrigem, coordsDestino, rotaCompleta string) {
	v.CoordenadasOrigem = coordsOrigem
	v.CoordenadasDestino = coordsDestino
	v.RotaCompleta = rotaCompleta
	v.UpdatedAt = time.Now()
}

// Erros de domínio
var (
	ErrDataInicioMaiorQueFim    = NewDomainError("data de início não pode ser maior que data de fim")
	ErrDataInicioPassada        = NewDomainError("data de início não pode ser no passado")
	ErrValorInvalido            = NewDomainError("valor deve ser maior que zero")
	ErrOrigemDestinoObrigatorios = NewDomainError("origem e destino são obrigatórios")
)

// DomainError representa um erro de domínio
type DomainError struct {
	message string
}

func NewDomainError(message string) *DomainError {
	return &DomainError{message: message}
}

func (e *DomainError) Error() string {
	return e.message
} 