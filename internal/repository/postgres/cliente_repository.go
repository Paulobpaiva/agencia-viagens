package postgres

import (
	"context"

	"agencia-viagens/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type clienteRepository struct {
	db *gorm.DB
}

// NewClienteRepository cria uma nova instância do repositório de clientes
func NewClienteRepository(db *gorm.DB) domain.ClienteRepository {
	return &clienteRepository{db: db}
}

func (r *clienteRepository) Create(ctx context.Context, cliente *domain.Cliente) error {
	return r.db.WithContext(ctx).Create(cliente).Error
}

func (r *clienteRepository) Update(ctx context.Context, cliente *domain.Cliente) error {
	return r.db.WithContext(ctx).Save(cliente).Error
}

func (r *clienteRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Cliente{}, "id = ?", id).Error
}

func (r *clienteRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Cliente, error) {
	var cliente domain.Cliente
	err := r.db.WithContext(ctx).First(&cliente, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &cliente, nil
}

func (r *clienteRepository) List(ctx context.Context, offset, limit int) ([]*domain.Cliente, error) {
	var clientes []*domain.Cliente
	err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Order("nome ASC").
		Find(&clientes).Error
	if err != nil {
		return nil, err
	}
	return clientes, nil
}

func (r *clienteRepository) GetByCPFCNPJ(ctx context.Context, cpfCnpj string) (*domain.Cliente, error) {
	var cliente domain.Cliente
	err := r.db.WithContext(ctx).First(&cliente, "cpf_cnpj = ?", cpfCnpj).Error
	if err != nil {
		return nil, err
	}
	return &cliente, nil
}

func (r *clienteRepository) GetByTipo(ctx context.Context, tipo domain.TipoCliente) ([]*domain.Cliente, error) {
	var clientes []*domain.Cliente
	err := r.db.WithContext(ctx).
		Where("tipo = ?", tipo).
		Order("nome ASC").
		Find(&clientes).Error
	if err != nil {
		return nil, err
	}
	return clientes, nil
}

func (r *clienteRepository) GetAtivos(ctx context.Context) ([]*domain.Cliente, error) {
	var clientes []*domain.Cliente
	err := r.db.WithContext(ctx).
		Where("ativo = ?", true).
		Order("nome ASC").
		Find(&clientes).Error
	if err != nil {
		return nil, err
	}
	return clientes, nil
}

// Métodos auxiliares para gestão de clientes

// GetClientesPorCidade retorna clientes agrupados por cidade
func (r *clienteRepository) GetClientesPorCidade(ctx context.Context) (map[string]int, error) {
	type Result struct {
		Cidade string
		Total  int
	}

	var results []Result
	err := r.db.WithContext(ctx).
		Model(&domain.Cliente{}).
		Select("cidade, count(*) as total").
		Where("cidade IS NOT NULL").
		Group("cidade").
		Order("total DESC").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	clientesPorCidade := make(map[string]int)
	for _, r := range results {
		clientesPorCidade[r.Cidade] = r.Total
	}

	return clientesPorCidade, nil
}

// GetClientesPorEstado retorna clientes agrupados por estado
func (r *clienteRepository) GetClientesPorEstado(ctx context.Context) (map[string]int, error) {
	type Result struct {
		Estado string
		Total  int
	}

	var results []Result
	err := r.db.WithContext(ctx).
		Model(&domain.Cliente{}).
		Select("estado, count(*) as total").
		Where("estado IS NOT NULL").
		Group("estado").
		Order("total DESC").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	clientesPorEstado := make(map[string]int)
	for _, r := range results {
		clientesPorEstado[r.Estado] = r.Total
	}

	return clientesPorEstado, nil
}

// GetClientesPorTipo retorna a quantidade de clientes por tipo (PF/PJ)
func (r *clienteRepository) GetClientesPorTipo(ctx context.Context) (map[domain.TipoCliente]int, error) {
	type Result struct {
		Tipo  domain.TipoCliente
		Total int
	}

	var results []Result
	err := r.db.WithContext(ctx).
		Model(&domain.Cliente{}).
		Select("tipo, count(*) as total").
		Group("tipo").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	clientesPorTipo := make(map[domain.TipoCliente]int)
	for _, r := range results {
		clientesPorTipo[r.Tipo] = r.Total
	}

	return clientesPorTipo, nil
}

// GetClientesLimiteCreditoExcedido retorna clientes com limite de crédito excedido
func (r *clienteRepository) GetClientesLimiteCreditoExcedido(ctx context.Context, limite float64) ([]*domain.Cliente, error) {
	var clientes []*domain.Cliente
	err := r.db.WithContext(ctx).
		Where("limite_credito > ?", limite).
		Order("limite_credito DESC").
		Find(&clientes).Error
	if err != nil {
		return nil, err
	}
	return clientes, nil
}
