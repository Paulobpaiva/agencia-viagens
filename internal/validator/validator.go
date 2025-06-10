package validator

import (
	"errors"
	"regexp"
	"time"

	"agencia-viagens/internal/domain"
)

var (
	ErrPlacaInvalida           = errors.New("placa inválida")
	ErrAnoInvalido             = errors.New("ano inválido")
	ErrCapacidadeInvalida      = errors.New("capacidade inválida")
	ErrTipoVeiculoInvalido     = errors.New("tipo de veículo inválido")
	ErrStatusVeiculoInvalido   = errors.New("status de veículo inválido")
	ErrCPFInvalido             = errors.New("CPF inválido")
	ErrCNHInvalida             = errors.New("CNH inválida")
	ErrCategoriaCNHInvalida    = errors.New("categoria de CNH inválida")
	ErrStatusMotoristaInvalido = errors.New("status de motorista inválido")
	ErrEmailInvalido           = errors.New("email inválido")
	ErrTelefoneInvalido        = errors.New("telefone inválido")
	ErrDataNascimentoInvalida  = errors.New("data de nascimento inválida")
	ErrPeriodoInvalido         = errors.New("período inválido")
	ErrValorInvalido           = errors.New("valor inválido")
	ErrStatusViagemInvalido    = errors.New("status de viagem inválido")
	ErrIDInvalido              = errors.New("ID inválido")
)

// ValidarPlaca valida se a placa do veículo está no formato correto
func ValidarPlaca(placa string) error {
	// Formato Mercosul: ABC1D23 ou ABC1234
	regex := regexp.MustCompile(`^[A-Z]{3}[0-9][A-Z0-9][0-9]{2}$`)
	if !regex.MatchString(placa) {
		return ErrPlacaInvalida
	}
	return nil
}

// ValidarAno valida se o ano do veículo é válido
func ValidarAno(ano int) error {
	anoAtual := time.Now().Year()
	if ano < 1900 || ano > anoAtual+1 {
		return ErrAnoInvalido
	}
	return nil
}

// ValidarCapacidade valida se a capacidade do veículo é válida
func ValidarCapacidade(capacidade int) error {
	if capacidade <= 0 || capacidade > 100 {
		return ErrCapacidadeInvalida
	}
	return nil
}

// ValidarTipoVeiculo valida se o tipo de veículo é válido
func ValidarTipoVeiculo(tipo domain.TipoVeiculo) error {
	switch tipo {
	case domain.TipoVan, domain.TipoOnibus, domain.TipoMicroOnibus:
		return nil
	default:
		return ErrTipoVeiculoInvalido
	}
}

// ValidarStatusVeiculo valida se o status do veículo é válido
func ValidarStatusVeiculo(status domain.StatusVeiculo) error {
	switch string(status) {
	case string(domain.StatusDisponivel), string(domain.StatusEmUso), string(domain.StatusManutencao), string(domain.StatusInativo):
		return nil
	default:
		return ErrStatusVeiculoInvalido
	}
}

// ValidarCPF valida se o CPF está no formato correto
func ValidarCPF(cpf string) error {
	// Remove caracteres não numéricos
	regex := regexp.MustCompile(`[^0-9]`)
	cpf = regex.ReplaceAllString(cpf, "")

	// Verifica se tem 11 dígitos
	if len(cpf) != 11 {
		return ErrCPFInvalido
	}

	// Verifica se todos os dígitos são iguais
	if cpf[0] == cpf[1] && cpf[1] == cpf[2] && cpf[2] == cpf[3] && cpf[3] == cpf[4] && cpf[4] == cpf[5] && cpf[5] == cpf[6] && cpf[6] == cpf[7] && cpf[7] == cpf[8] && cpf[8] == cpf[9] && cpf[9] == cpf[10] {
		return ErrCPFInvalido
	}

	// Validação do primeiro dígito verificador
	soma := 0
	for i := 0; i < 9; i++ {
		soma += int(cpf[i]-'0') * (10 - i)
	}
	resto := soma % 11
	if resto < 2 {
		if int(cpf[9]-'0') != 0 {
			return ErrCPFInvalido
		}
	} else {
		if int(cpf[9]-'0') != 11-resto {
			return ErrCPFInvalido
		}
	}

	// Validação do segundo dígito verificador
	soma = 0
	for i := 0; i < 10; i++ {
		soma += int(cpf[i]-'0') * (11 - i)
	}
	resto = soma % 11
	if resto < 2 {
		if int(cpf[10]-'0') != 0 {
			return ErrCPFInvalido
		}
	} else {
		if int(cpf[10]-'0') != 11-resto {
			return ErrCPFInvalido
		}
	}

	return nil
}

// ValidarCNH valida se a CNH está no formato correto
func ValidarCNH(cnh string) error {
	// Remove caracteres não numéricos
	regex := regexp.MustCompile(`[^0-9]`)
	cnh = regex.ReplaceAllString(cnh, "")

	// Verifica se tem 11 dígitos
	if len(cnh) != 11 {
		return ErrCNHInvalida
	}

	// Verifica se todos os dígitos são iguais
	if cnh[0] == cnh[1] && cnh[1] == cnh[2] && cnh[2] == cnh[3] && cnh[3] == cnh[4] && cnh[4] == cnh[5] && cnh[5] == cnh[6] && cnh[6] == cnh[7] && cnh[7] == cnh[8] && cnh[8] == cnh[9] && cnh[9] == cnh[10] {
		return ErrCNHInvalida
	}

	// Validação do primeiro dígito verificador
	soma := 0
	for i := 0; i < 9; i++ {
		soma += int(cnh[i]-'0') * (9 - i)
	}
	resto := soma % 11
	if resto < 2 {
		if int(cnh[9]-'0') != 0 {
			return ErrCNHInvalida
		}
	} else {
		if int(cnh[9]-'0') != 11-resto {
			return ErrCNHInvalida
		}
	}

	// Validação do segundo dígito verificador
	soma = 0
	for i := 0; i < 10; i++ {
		soma += int(cnh[i]-'0') * (10 - i)
	}
	resto = soma % 11
	if resto < 2 {
		if int(cnh[10]-'0') != 0 {
			return ErrCNHInvalida
		}
	} else {
		if int(cnh[10]-'0') != 11-resto {
			return ErrCNHInvalida
		}
	}

	return nil
}

// ValidarCategoriaCNH valida se a categoria da CNH é válida
func ValidarCategoriaCNH(categoria domain.TipoCNH) error {
	switch categoria {
	case domain.CNHA, domain.CNHB, domain.CNHC, domain.CNHD, domain.CNHE:
		return nil
	default:
		return ErrCategoriaCNHInvalida
	}
}

// ValidarStatusMotorista valida se o status do motorista é válido
func ValidarStatusMotorista(status domain.StatusMotorista) error {
	switch string(status) {
	case string(domain.StatusDisponivel), string(domain.StatusEmViagem), string(domain.StatusFolga), string(domain.StatusInativo):
		return nil
	default:
		return ErrStatusMotoristaInvalido
	}
}

// ValidarEmail valida se o email está no formato correto
func ValidarEmail(email string) error {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !regex.MatchString(email) {
		return ErrEmailInvalido
	}
	return nil
}

// ValidarTelefone valida se o telefone está no formato correto
func ValidarTelefone(telefone string) error {
	// Remove caracteres não numéricos
	regex := regexp.MustCompile(`[^0-9]`)
	telefone = regex.ReplaceAllString(telefone, "")

	// Verifica se tem entre 10 e 11 dígitos (com DDD)
	if len(telefone) < 10 || len(telefone) > 11 {
		return ErrTelefoneInvalido
	}

	return nil
}

// ValidarDataNascimento valida se a data de nascimento é válida
func ValidarDataNascimento(dataNascimento time.Time) error {
	idadeMinima := 18
	idadeMaxima := 70
	hoje := time.Now()
	idade := hoje.Year() - dataNascimento.Year()

	// Ajusta a idade se o aniversário ainda não ocorreu este ano
	if hoje.Month() < dataNascimento.Month() || (hoje.Month() == dataNascimento.Month() && hoje.Day() < dataNascimento.Day()) {
		idade--
	}

	if idade < idadeMinima || idade > idadeMaxima {
		return ErrDataNascimentoInvalida
	}

	return nil
}

// ValidarPeriodo valida se o período é válido
func ValidarPeriodo(dataInicio, dataFim time.Time) error {
	if dataInicio.IsZero() || dataFim.IsZero() {
		return ErrPeriodoInvalido
	}

	if dataInicio.After(dataFim) {
		return ErrPeriodoInvalido
	}

	// Verifica se a data de início é futura
	if dataInicio.Before(time.Now()) {
		return ErrPeriodoInvalido
	}

	// Verifica se o período não é maior que 30 dias
	if dataFim.Sub(dataInicio) > 30*24*time.Hour {
		return ErrPeriodoInvalido
	}

	return nil
}

// ValidarValor valida se o valor é válido
func ValidarValor(valor float64) error {
	if valor <= 0 {
		return ErrValorInvalido
	}
	return nil
}

// ValidarStatusViagem valida se o status da viagem é válido
func ValidarStatusViagem(status domain.StatusViagem) error {
	switch string(status) {
	case string(domain.StatusAgendada), string(domain.StatusEmAndamento), string(domain.StatusConcluida), string(domain.StatusCancelada):
		return nil
	default:
		return ErrStatusViagemInvalido
	}
}
