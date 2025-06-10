package postgres

import (
	"context"
	"time"

	"agencia-viagens/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type motoristaRepository struct {
	db *gorm.DB
}

// NewMotoristaRepository cria uma nova instância do repositório de motoristas
func NewMotoristaRepository(db *gorm.DB) domain.MotoristaRepository {
	return &motoristaRepository{db: db}
}

func (r *motoristaRepository) Create(ctx context.Context, motorista *domain.Motorista) error {
	return r.db.WithContext(ctx).Create(motorista).Error
}

func (r *motoristaRepository) Update(ctx context.Context, motorista *domain.Motorista) error {
	return r.db.WithContext(ctx).Save(motorista).Error
}

func (r *motoristaRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Motorista{}, "id = ?", id).Error
}

func (r *motoristaRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Motorista, error) {
	var motorista domain.Motorista
	err := r.db.WithContext(ctx).First(&motorista, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &motorista, nil
}

func (r *motoristaRepository) List(ctx context.Context, offset, limit int) ([]*domain.Motorista, error) {
	var motoristas []*domain.Motorista
	err := r.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Order("nome ASC").
		Find(&motoristas).Error
	if err != nil {
		return nil, err
	}
	return motoristas, nil
}

func (r *motoristaRepository) GetByCPF(ctx context.Context, cpf string) (*domain.Motorista, error) {
	var motorista domain.Motorista
	err := r.db.WithContext(ctx).First(&motorista, "cpf = ?", cpf).Error
	if err != nil {
		return nil, err
	}
	return &motorista, nil
}

func (r *motoristaRepository) GetByCNH(ctx context.Context, cnh string) (*domain.Motorista, error) {
	var motorista domain.Motorista
	err := r.db.WithContext(ctx).First(&motorista, "cnh = ?", cnh).Error
	if err != nil {
		return nil, err
	}
	return &motorista, nil
}

func (r *motoristaRepository) GetByStatus(ctx context.Context, status domain.StatusMotorista) ([]*domain.Motorista, error) {
	var motoristas []*domain.Motorista
	err := r.db.WithContext(ctx).
		Where("status = ?", status).
		Order("nome ASC").
		Find(&motoristas).Error
	if err != nil {
		return nil, err
	}
	return motoristas, nil
}

func (r *motoristaRepository) GetDisponiveis(ctx context.Context, dataInicio, dataFim time.Time) ([]*domain.Motorista, error) {
	var motoristas []*domain.Motorista

	// Subquery para encontrar motoristas ocupados no período
	subQuery := r.db.Model(&domain.Viagem{}).
		Select("motorista_id").
		Where("status != ? AND ((data_inicio BETWEEN ? AND ?) OR (data_fim BETWEEN ? AND ?))",
			domain.StatusCancelada, dataInicio, dataFim, dataInicio, dataFim)

	// Query principal para encontrar motoristas disponíveis
	err := r.db.WithContext(ctx).
		Where("status = ? AND disponivel = ? AND id NOT IN (?)",
			domain.StatusDisponivel, true, subQuery).
		Order("nome ASC").
		Find(&motoristas).Error
	if err != nil {
		return nil, err
	}
	return motoristas, nil
}

// Métodos auxiliares para gestão de motoristas

// GetMotoristasCNHVencida retorna motoristas com CNH vencida
func (r *motoristaRepository) GetMotoristasCNHVencida(ctx context.Context) ([]*domain.Motorista, error) {
	var motoristas []*domain.Motorista
	err := r.db.WithContext(ctx).
		Where("validade_cnh <= ?", time.Now()).
		Order("validade_cnh ASC").
		Find(&motoristas).Error
	if err != nil {
		return nil, err
	}
	return motoristas, nil
}

// GetMotoristasProximosVencimentoCNH retorna motoristas com CNH próxima do vencimento
func (r *motoristaRepository) GetMotoristasProximosVencimentoCNH(ctx context.Context) ([]*domain.Motorista, error) {
	var motoristas []*domain.Motorista
	err := r.db.WithContext(ctx).
		Where("validade_cnh BETWEEN ? AND ?",
			time.Now(), time.Now().AddDate(0, 3, 0)). // Próximos 3 meses
		Order("validade_cnh ASC").
		Find(&motoristas).Error
	if err != nil {
		return nil, err
	}
	return motoristas, nil
}

// GetMotoristasBancoHorasExcedido retorna motoristas com banco de horas excedido
func (r *motoristaRepository) GetMotoristasBancoHorasExcedido(ctx context.Context, limiteHoras int) ([]*domain.Motorista, error) {
	var motoristas []*domain.Motorista
	err := r.db.WithContext(ctx).
		Where("banco_horas > ?", limiteHoras*60). // Converte horas para minutos
		Order("banco_horas DESC").
		Find(&motoristas).Error
	if err != nil {
		return nil, err
	}
	return motoristas, nil
}
