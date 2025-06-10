package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/paulopaiva/agencia-viagens/internal/domain"
	"gorm.io/gorm"
)

type veiculoRepository struct {
	db *gorm.DB
}

// NewVeiculoRepository cria uma nova instância do repositório de veículos
func NewVeiculoRepository(db *gorm.DB) domain.VeiculoRepository {
	return &veiculoRepository{db: db}
}

func (r *veiculoRepository) Create(ctx context.Context, veiculo *domain.Veiculo) error {
	return r.db.WithContext(ctx).Create(veiculo).Error
}

func (r *veiculoRepository) Update(ctx context.Context, veiculo *domain.Veiculo) error {
	return r.db.WithContext(ctx).Save(veiculo).Error
}

func (r *veiculoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Veiculo{}, "id = ?", id).Error
}

func (r *veiculoRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Veiculo, error) {
	var veiculo domain.Veiculo
	err := r.db.WithContext(ctx).First(&veiculo, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &veiculo, nil
}

func (r *veiculoRepository) List(ctx context.Context, offset, limit int) ([]*domain.Veiculo, error) {
	var veiculos []*domain.Veiculo
	err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Order("placa ASC").
		Find(&veiculos).Error
	if err != nil {
		return nil, err
	}
	return veiculos, nil
}

func (r *veiculoRepository) GetByPlaca(ctx context.Context, placa string) (*domain.Veiculo, error) {
	var veiculo domain.Veiculo
	err := r.db.WithContext(ctx).First(&veiculo, "placa = ?", placa).Error
	if err != nil {
		return nil, err
	}
	return &veiculo, nil
}

func (r *veiculoRepository) GetByStatus(ctx context.Context, status domain.StatusVeiculo) ([]*domain.Veiculo, error) {
	var veiculos []*domain.Veiculo
	err := r.db.WithContext(ctx).
		Where("status = ?", status).
		Order("placa ASC").
		Find(&veiculos).Error
	if err != nil {
		return nil, err
	}
	return veiculos, nil
}

func (r *veiculoRepository) GetByTipo(ctx context.Context, tipo domain.TipoVeiculo) ([]*domain.Veiculo, error) {
	var veiculos []*domain.Veiculo
	err := r.db.WithContext(ctx).
		Where("tipo = ?", tipo).
		Order("placa ASC").
		Find(&veiculos).Error
	if err != nil {
		return nil, err
	}
	return veiculos, nil
}

func (r *veiculoRepository) GetDisponiveis(ctx context.Context, dataInicio, dataFim time.Time) ([]*domain.Veiculo, error) {
	var veiculos []*domain.Veiculo
	
	// Subquery para encontrar veículos ocupados no período
	subQuery := r.db.Model(&domain.Viagem{}).
		Select("veiculo_id").
		Where("status != ? AND ((data_inicio BETWEEN ? AND ?) OR (data_fim BETWEEN ? AND ?))",
			domain.StatusCancelada, dataInicio, dataFim, dataInicio, dataFim)
	
	// Query principal para encontrar veículos disponíveis
	err := r.db.WithContext(ctx).
		Where("status = ? AND id NOT IN (?)", domain.StatusDisponivel, subQuery).
		Order("placa ASC").
		Find(&veiculos).Error
	if err != nil {
		return nil, err
	}
	return veiculos, nil
}

// Métodos auxiliares para manutenção

// GetVeiculosProximaManutencao retorna veículos que precisam de manutenção
func (r *veiculoRepository) GetVeiculosProximaManutencao(ctx context.Context) ([]*domain.Veiculo, error) {
	var veiculos []*domain.Veiculo
	err := r.db.WithContext(ctx).
		Where("proxima_manutencao <= ?", time.Now().AddDate(0, 1, 0)). // Próximo mês
		Order("proxima_manutencao ASC").
		Find(&veiculos).Error
	if err != nil {
		return nil, err
	}
	return veiculos, nil
}

// GetVeiculosDocumentacaoVencida retorna veículos com documentação vencida
func (r *veiculoRepository) GetVeiculosDocumentacaoVencida(ctx context.Context) ([]*domain.Veiculo, error) {
	var veiculos []*domain.Veiculo
	err := r.db.WithContext(ctx).
		Where("documentacao_valida = ? AND vencimento_documentacao <= ?", true, time.Now()).
		Order("vencimento_documentacao ASC").
		Find(&veiculos).Error
	if err != nil {
		return nil, err
	}
	return veiculos, nil
} 