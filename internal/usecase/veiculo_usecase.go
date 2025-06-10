package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/paulopaiva/agencia-viagens/internal/domain"
	"github.com/paulopaiva/agencia-viagens/internal/repository"
)

var (
	ErrVeiculoNaoEncontrado = errors.New("veículo não encontrado")
	ErrPlacaInvalida        = errors.New("placa inválida")
	ErrCapacidadeInvalida   = errors.New("capacidade inválida")
)

type VeiculoUseCase struct {
	veiculoRepo repository.VeiculoRepository
}

func NewVeiculoUseCase(veiculoRepo repository.VeiculoRepository) *VeiculoUseCase {
	return &VeiculoUseCase{
		veiculoRepo: veiculoRepo,
	}
}

func (uc *VeiculoUseCase) Criar(ctx context.Context, veiculo *domain.Veiculo) error {
	// Validações básicas
	if err := uc.validarVeiculo(veiculo); err != nil {
		return err
	}

	return uc.veiculoRepo.Create(ctx, veiculo)
}

func (uc *VeiculoUseCase) Listar(ctx context.Context) ([]domain.Veiculo, error) {
	veiculos, err := uc.veiculoRepo.List(ctx, 0, 100) // TODO: Implementar paginação
	if err != nil {
		return nil, err
	}

	result := make([]domain.Veiculo, len(veiculos))
	for i, v := range veiculos {
		result[i] = *v
	}
	return result, nil
}

func (uc *VeiculoUseCase) BuscarPorID(ctx context.Context, id uuid.UUID) (*domain.Veiculo, error) {
	veiculo, err := uc.veiculoRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrVeiculoNaoEncontrado
	}
	return veiculo, nil
}

func (uc *VeiculoUseCase) Atualizar(ctx context.Context, veiculo *domain.Veiculo) error {
	// Verifica se o veículo existe
	if _, err := uc.veiculoRepo.GetByID(ctx, veiculo.ID); err != nil {
		return ErrVeiculoNaoEncontrado
	}

	// Validações básicas
	if err := uc.validarVeiculo(veiculo); err != nil {
		return err
	}

	return uc.veiculoRepo.Update(ctx, veiculo)
}

func (uc *VeiculoUseCase) Remover(ctx context.Context, id uuid.UUID) error {
	// Verifica se o veículo existe
	if _, err := uc.veiculoRepo.GetByID(ctx, id); err != nil {
		return ErrVeiculoNaoEncontrado
	}

	return uc.veiculoRepo.Delete(ctx, id)
}

func (uc *VeiculoUseCase) validarVeiculo(veiculo *domain.Veiculo) error {
	// Validação da placa (formato brasileiro)
	if len(veiculo.Placa) != 7 {
		return ErrPlacaInvalida
	}

	// Validação da capacidade
	if veiculo.Capacidade <= 0 {
		return ErrCapacidadeInvalida
	}

	return nil
}
