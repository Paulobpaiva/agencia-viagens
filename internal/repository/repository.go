package repository

import (
	"context"
	"time"

	"agencia-viagens/internal/config"
	"agencia-viagens/internal/domain"
	"agencia-viagens/internal/repository/postgres"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ViagemRepository define as operações do repositório de viagens
type ViagemRepository interface {
	Create(ctx context.Context, viagem *domain.Viagem) error
	Update(ctx context.Context, viagem *domain.Viagem) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Viagem, error)
	List(ctx context.Context, offset, limit int) ([]*domain.Viagem, error)

	// Métodos específicos
	GetByVeiculo(ctx context.Context, veiculoID uuid.UUID, dataInicio, dataFim time.Time) ([]*domain.Viagem, error)
	GetByMotorista(ctx context.Context, motoristaID uuid.UUID, dataInicio, dataFim time.Time) ([]*domain.Viagem, error)
	GetByCliente(ctx context.Context, clienteID uuid.UUID) ([]*domain.Viagem, error)
	CheckDisponibilidade(ctx context.Context, veiculoID uuid.UUID, dataInicio, dataFim time.Time) (bool, error)
}

// VeiculoRepository define as operações do repositório de veículos
type VeiculoRepository interface {
	Create(ctx context.Context, veiculo *domain.Veiculo) error
	Update(ctx context.Context, veiculo *domain.Veiculo) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Veiculo, error)
	List(ctx context.Context, offset, limit int) ([]*domain.Veiculo, error)

	// Métodos específicos
	GetByPlaca(ctx context.Context, placa string) (*domain.Veiculo, error)
	GetByStatus(ctx context.Context, status domain.StatusVeiculo) ([]*domain.Veiculo, error)
	GetByTipo(ctx context.Context, tipo domain.TipoVeiculo) ([]*domain.Veiculo, error)
	GetDisponiveis(ctx context.Context, dataInicio, dataFim time.Time) ([]*domain.Veiculo, error)
}

// MotoristaRepository define as operações do repositório de motoristas
type MotoristaRepository interface {
	Create(ctx context.Context, motorista *domain.Motorista) error
	Update(ctx context.Context, motorista *domain.Motorista) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Motorista, error)
	List(ctx context.Context, offset, limit int) ([]*domain.Motorista, error)

	// Métodos específicos
	GetByCPF(ctx context.Context, cpf string) (*domain.Motorista, error)
	GetByCNH(ctx context.Context, cnh string) (*domain.Motorista, error)
	GetByStatus(ctx context.Context, status domain.StatusMotorista) ([]*domain.Motorista, error)
	GetDisponiveis(ctx context.Context, dataInicio, dataFim time.Time) ([]*domain.Motorista, error)
}

// ClienteRepository define as operações do repositório de clientes
type ClienteRepository interface {
	Create(ctx context.Context, cliente *domain.Cliente) error
	Update(ctx context.Context, cliente *domain.Cliente) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Cliente, error)
	List(ctx context.Context, offset, limit int) ([]*domain.Cliente, error)

	// Métodos específicos
	GetByCPFCNPJ(ctx context.Context, cpfCnpj string) (*domain.Cliente, error)
	GetByTipo(ctx context.Context, tipo domain.TipoCliente) ([]*domain.Cliente, error)
	GetAtivos(ctx context.Context) ([]*domain.Cliente, error)
}

// TransactionManager define a interface para gerenciamento de transações
type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

// NewPostgresDB cria uma nova conexão com o banco de dados PostgreSQL
func NewPostgresDB(cfg config.DatabaseConfig) (*gorm.DB, error) {
	return postgres.NewPostgresDB(cfg)
}

// NewViagemRepository cria uma nova instância do repositório de viagens
func NewViagemRepository(db *gorm.DB) domain.ViagemRepository {
	return postgres.NewViagemRepository(db)
}

// NewVeiculoRepository cria uma nova instância do repositório de veículos
func NewVeiculoRepository(db *gorm.DB) domain.VeiculoRepository {
	return postgres.NewVeiculoRepository(db)
}

// NewMotoristaRepository cria uma nova instância do repositório de motoristas
func NewMotoristaRepository(db *gorm.DB) domain.MotoristaRepository {
	return postgres.NewMotoristaRepository(db)
}

// NewClienteRepository cria uma nova instância do repositório de clientes
func NewClienteRepository(db *gorm.DB) domain.ClienteRepository {
	return postgres.NewClienteRepository(db)
}

// NewTransactionManager cria uma nova instância do gerenciador de transações
func NewTransactionManager(db *gorm.DB) TransactionManager {
	return postgres.NewTransactionManager(db)
}
