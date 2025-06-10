package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/paulopaiva/agencia-viagens/internal/domain"
	"gorm.io/gorm"
)

type viagemRepository struct {
	db *gorm.DB
}

// NewViagemRepository cria uma nova instância do repositório de viagens
func NewViagemRepository(db *gorm.DB) domain.ViagemRepository {
	return &viagemRepository{db: db}
}

func (r *viagemRepository) Create(ctx context.Context, viagem *domain.Viagem) error {
	return r.db.WithContext(ctx).Create(viagem).Error
}

func (r *viagemRepository) Update(ctx context.Context, viagem *domain.Viagem) error {
	return r.db.WithContext(ctx).Save(viagem).Error
}

func (r *viagemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Viagem{}, "id = ?", id).Error
}

func (r *viagemRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Viagem, error) {
	var viagem domain.Viagem
	err := r.db.WithContext(ctx).
		Preload("Veiculo").
		Preload("Motorista").
		Preload("Cliente").
		First(&viagem, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &viagem, nil
}

func (r *viagemRepository) List(ctx context.Context, offset, limit int) ([]*domain.Viagem, error) {
	var viagens []*domain.Viagem
	err := r.db.WithContext(ctx).
		Preload("Veiculo").
		Preload("Motorista").
		Preload("Cliente").
		Offset(offset).
		Limit(limit).
		Order("data_inicio DESC").
		Find(&viagens).Error
	if err != nil {
		return nil, err
	}
	return viagens, nil
}

func (r *viagemRepository) GetByVeiculo(ctx context.Context, veiculoID uuid.UUID, 
	dataInicio, dataFim time.Time) ([]*domain.Viagem, error) {
	var viagens []*domain.Viagem
	err := r.db.WithContext(ctx).
		Where("veiculo_id = ? AND ((data_inicio BETWEEN ? AND ?) OR (data_fim BETWEEN ? AND ?))",
			veiculoID, dataInicio, dataFim, dataInicio, dataFim).
		Preload("Veiculo").
		Preload("Motorista").
		Preload("Cliente").
		Order("data_inicio ASC").
		Find(&viagens).Error
	if err != nil {
		return nil, err
	}
	return viagens, nil
}

func (r *viagemRepository) GetByMotorista(ctx context.Context, motoristaID uuid.UUID,
	dataInicio, dataFim time.Time) ([]*domain.Viagem, error) {
	var viagens []*domain.Viagem
	err := r.db.WithContext(ctx).
		Where("motorista_id = ? AND ((data_inicio BETWEEN ? AND ?) OR (data_fim BETWEEN ? AND ?))",
			motoristaID, dataInicio, dataFim, dataInicio, dataFim).
		Preload("Veiculo").
		Preload("Motorista").
		Preload("Cliente").
		Order("data_inicio ASC").
		Find(&viagens).Error
	if err != nil {
		return nil, err
	}
	return viagens, nil
}

func (r *viagemRepository) GetByCliente(ctx context.Context, clienteID uuid.UUID) ([]*domain.Viagem, error) {
	var viagens []*domain.Viagem
	err := r.db.WithContext(ctx).
		Where("cliente_id = ?", clienteID).
		Preload("Veiculo").
		Preload("Motorista").
		Preload("Cliente").
		Order("data_inicio DESC").
		Find(&viagens).Error
	if err != nil {
		return nil, err
	}
	return viagens, nil
}

func (r *viagemRepository) CheckDisponibilidade(ctx context.Context, veiculoID uuid.UUID,
	dataInicio, dataFim time.Time) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&domain.Viagem{}).
		Where("veiculo_id = ? AND status != ? AND ((data_inicio BETWEEN ? AND ?) OR (data_fim BETWEEN ? AND ?))",
			veiculoID, domain.StatusCancelada, dataInicio, dataFim, dataInicio, dataFim).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 0, nil
} 