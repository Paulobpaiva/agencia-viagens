package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/paulopaiva/agencia-viagens/internal/domain"
	"github.com/paulopaiva/agencia-viagens/internal/repository"
)

var (
	ErrViagemNaoEncontrada   = errors.New("viagem não encontrada")
	ErrVeiculoIndisponivel   = errors.New("veículo indisponível para o período")
	ErrMotoristaIndisponivel = errors.New("motorista indisponível para o período")
	ErrDataInvalida          = errors.New("data inválida")
)

type ViagemUseCase struct {
	viagemRepo    repository.ViagemRepository
	veiculoRepo   repository.VeiculoRepository
	motoristaRepo repository.MotoristaRepository
}

func NewViagemUseCase(
	viagemRepo repository.ViagemRepository,
	veiculoRepo repository.VeiculoRepository,
	motoristaRepo repository.MotoristaRepository,
) *ViagemUseCase {
	return &ViagemUseCase{
		viagemRepo:    viagemRepo,
		veiculoRepo:   veiculoRepo,
		motoristaRepo: motoristaRepo,
	}
}

func (uc *ViagemUseCase) Criar(ctx context.Context, viagem *domain.Viagem) error {
	// Validações básicas
	if viagem.DataInicio.After(viagem.DataFim) {
		return ErrDataInvalida
	}

	// Verifica disponibilidade do veículo
	disponivel, err := uc.viagemRepo.CheckDisponibilidade(ctx, viagem.VeiculoID, viagem.DataInicio, viagem.DataFim)
	if err != nil {
		return err
	}
	if !disponivel {
		return ErrVeiculoIndisponivel
	}

	// Verifica disponibilidade do motorista
	viagens, err := uc.viagemRepo.GetByMotorista(ctx, viagem.MotoristaID, viagem.DataInicio, viagem.DataFim)
	if err != nil {
		return err
	}
	if len(viagens) > 0 {
		return ErrMotoristaIndisponivel
	}

	// Define status inicial
	viagem.Status = domain.StatusAgendada

	return uc.viagemRepo.Create(ctx, viagem)
}

func (uc *ViagemUseCase) Listar(ctx context.Context) ([]domain.Viagem, error) {
	viagens, err := uc.viagemRepo.List(ctx, 0, 100) // TODO: Implementar paginação
	if err != nil {
		return nil, err
	}

	result := make([]domain.Viagem, len(viagens))
	for i, v := range viagens {
		result[i] = *v
	}
	return result, nil
}

func (uc *ViagemUseCase) BuscarPorID(ctx context.Context, id uuid.UUID) (*domain.Viagem, error) {
	viagem, err := uc.viagemRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrViagemNaoEncontrada
	}
	return viagem, nil
}

func (uc *ViagemUseCase) Atualizar(ctx context.Context, viagem *domain.Viagem) error {
	// Verifica se a viagem existe
	existente, err := uc.viagemRepo.GetByID(ctx, viagem.ID)
	if err != nil {
		return ErrViagemNaoEncontrada
	}

	// Se a data foi alterada, verifica disponibilidade
	if !existente.DataInicio.Equal(viagem.DataInicio) || !existente.DataFim.Equal(viagem.DataFim) {
		if viagem.DataInicio.After(viagem.DataFim) {
			return ErrDataInvalida
		}

		// Verifica disponibilidade do veículo
		disponivel, err := uc.viagemRepo.CheckDisponibilidade(ctx, viagem.VeiculoID, viagem.DataInicio, viagem.DataFim)
		if err != nil {
			return err
		}
		if !disponivel {
			return ErrVeiculoIndisponivel
		}

		// Verifica disponibilidade do motorista
		viagens, err := uc.viagemRepo.GetByMotorista(ctx, viagem.MotoristaID, viagem.DataInicio, viagem.DataFim)
		if err != nil {
			return err
		}
		if len(viagens) > 0 {
			return ErrMotoristaIndisponivel
		}
	}

	return uc.viagemRepo.Update(ctx, viagem)
}

func (uc *ViagemUseCase) Cancelar(ctx context.Context, id uuid.UUID) error {
	viagem, err := uc.viagemRepo.GetByID(ctx, id)
	if err != nil {
		return ErrViagemNaoEncontrada
	}

	viagem.Status = domain.StatusCancelada
	return uc.viagemRepo.Update(ctx, viagem)
}
