package usecase

import (
	"context"
	"errors"
	"regexp"

	"agencia-viagens/internal/domain"
	"agencia-viagens/internal/repository"

	"github.com/google/uuid"
)

var (
	ErrMotoristaNaoEncontrado = errors.New("motorista não encontrado")
	ErrCNHInvalida            = errors.New("CNH inválida")
	ErrEmailInvalido          = errors.New("email inválido")
	ErrTelefoneInvalido       = errors.New("telefone inválido")
)

type MotoristaUseCase struct {
	motoristaRepo repository.MotoristaRepository
}

func NewMotoristaUseCase(motoristaRepo repository.MotoristaRepository) *MotoristaUseCase {
	return &MotoristaUseCase{
		motoristaRepo: motoristaRepo,
	}
}

func (uc *MotoristaUseCase) Criar(ctx context.Context, motorista *domain.Motorista) error {
	// Validações básicas
	if err := uc.validarMotorista(motorista); err != nil {
		return err
	}

	return uc.motoristaRepo.Create(ctx, motorista)
}

func (uc *MotoristaUseCase) Listar(ctx context.Context) ([]domain.Motorista, error) {
	motoristas, err := uc.motoristaRepo.List(ctx, 0, 100) // TODO: Implementar paginação
	if err != nil {
		return nil, err
	}

	result := make([]domain.Motorista, len(motoristas))
	for i, m := range motoristas {
		result[i] = *m
	}
	return result, nil
}

func (uc *MotoristaUseCase) BuscarPorID(ctx context.Context, id uuid.UUID) (*domain.Motorista, error) {
	motorista, err := uc.motoristaRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrMotoristaNaoEncontrado
	}
	return motorista, nil
}

func (uc *MotoristaUseCase) Atualizar(ctx context.Context, motorista *domain.Motorista) error {
	// Verifica se o motorista existe
	if _, err := uc.motoristaRepo.GetByID(ctx, motorista.ID); err != nil {
		return ErrMotoristaNaoEncontrado
	}

	// Validações básicas
	if err := uc.validarMotorista(motorista); err != nil {
		return err
	}

	return uc.motoristaRepo.Update(ctx, motorista)
}

func (uc *MotoristaUseCase) Remover(ctx context.Context, id uuid.UUID) error {
	// Verifica se o motorista existe
	if _, err := uc.motoristaRepo.GetByID(ctx, id); err != nil {
		return ErrMotoristaNaoEncontrado
	}

	return uc.motoristaRepo.Delete(ctx, id)
}

func (uc *MotoristaUseCase) validarMotorista(motorista *domain.Motorista) error {
	// Validação da CNH (formato brasileiro)
	cnhRegex := regexp.MustCompile(`^[0-9]{11}$`)
	if !cnhRegex.MatchString(motorista.CNH) {
		return ErrCNHInvalida
	}

	// Validação do email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(motorista.Email) {
		return ErrEmailInvalido
	}

	// Validação do telefone (formato brasileiro)
	telefoneRegex := regexp.MustCompile(`^\([0-9]{2}\) [0-9]{5}-[0-9]{4}$`)
	if !telefoneRegex.MatchString(motorista.Telefone) {
		return ErrTelefoneInvalido
	}

	return nil
}
