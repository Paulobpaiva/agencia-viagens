package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// ViagemRepository define as operações do repositório de viagens
type ViagemRepository interface {
	Create(ctx context.Context, viagem *Viagem) error
	Update(ctx context.Context, viagem *Viagem) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*Viagem, error)
	List(ctx context.Context, offset, limit int) ([]*Viagem, error)
	GetByVeiculo(ctx context.Context, veiculoID uuid.UUID, dataInicio, dataFim time.Time) ([]*Viagem, error)
	GetByMotorista(ctx context.Context, motoristaID uuid.UUID, dataInicio, dataFim time.Time) ([]*Viagem, error)
	GetByCliente(ctx context.Context, clienteID uuid.UUID) ([]*Viagem, error)
	CheckDisponibilidade(ctx context.Context, veiculoID uuid.UUID, dataInicio, dataFim time.Time) (bool, error)
}

// VeiculoRepository define as operações do repositório de veículos
type VeiculoRepository interface {
	Create(ctx context.Context, veiculo *Veiculo) error
	Update(ctx context.Context, veiculo *Veiculo) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*Veiculo, error)
	List(ctx context.Context, offset, limit int) ([]*Veiculo, error)
	GetByPlaca(ctx context.Context, placa string) (*Veiculo, error)
	GetByStatus(ctx context.Context, status StatusVeiculo) ([]*Veiculo, error)
	GetByTipo(ctx context.Context, tipo TipoVeiculo) ([]*Veiculo, error)
	GetDisponiveis(ctx context.Context, dataInicio, dataFim time.Time) ([]*Veiculo, error)
	GetVeiculosProximaManutencao(ctx context.Context) ([]*Veiculo, error)
	GetVeiculosDocumentacaoVencida(ctx context.Context) ([]*Veiculo, error)
}

// MotoristaRepository define as operações do repositório de motoristas
type MotoristaRepository interface {
	Create(ctx context.Context, motorista *Motorista) error
	Update(ctx context.Context, motorista *Motorista) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*Motorista, error)
	List(ctx context.Context, offset, limit int) ([]*Motorista, error)
	GetByCPF(ctx context.Context, cpf string) (*Motorista, error)
	GetByCNH(ctx context.Context, cnh string) (*Motorista, error)
	GetByStatus(ctx context.Context, status StatusMotorista) ([]*Motorista, error)
	GetDisponiveis(ctx context.Context, dataInicio, dataFim time.Time) ([]*Motorista, error)
	GetMotoristasCNHVencida(ctx context.Context) ([]*Motorista, error)
	GetMotoristasProximosVencimentoCNH(ctx context.Context) ([]*Motorista, error)
	GetMotoristasBancoHorasExcedido(ctx context.Context, limiteHoras int) ([]*Motorista, error)
}

// ClienteRepository define as operações do repositório de clientes
type ClienteRepository interface {
	Create(ctx context.Context, cliente *Cliente) error
	Update(ctx context.Context, cliente *Cliente) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*Cliente, error)
	List(ctx context.Context, offset, limit int) ([]*Cliente, error)
	GetByCPFCNPJ(ctx context.Context, cpfCnpj string) (*Cliente, error)
	GetByTipo(ctx context.Context, tipo TipoCliente) ([]*Cliente, error)
	GetAtivos(ctx context.Context) ([]*Cliente, error)
	GetClientesPorCidade(ctx context.Context) (map[string]int, error)
	GetClientesPorEstado(ctx context.Context) (map[string]int, error)
	GetClientesPorTipo(ctx context.Context) (map[TipoCliente]int, error)
	GetClientesLimiteCreditoExcedido(ctx context.Context, limite float64) ([]*Cliente, error)
}
