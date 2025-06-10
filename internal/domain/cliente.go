package domain

import (
	"time"

	"github.com/google/uuid"
)

// TipoCliente representa os tipos de clientes
type TipoCliente string

const (
	TipoPessoaFisica  TipoCliente = "PF"
	TipoPessoaJuridica TipoCliente = "PJ"
)

// Cliente representa um cliente no sistema
type Cliente struct {
	ID          uuid.UUID   `json:"id" gorm:"type:uuid;primary_key"`
	Tipo        TipoCliente `json:"tipo" gorm:"type:varchar(2);not null"`
	
	// Dados básicos
	Nome        string      `json:"nome" gorm:"type:varchar(100);not null"`
	CPFCNPJ     string      `json:"cpf_cnpj" gorm:"type:varchar(14);unique;not null"`
	RG          string      `json:"rg" gorm:"type:varchar(20);unique"` // Apenas para PF
	DataNascimento time.Time `json:"data_nascimento" gorm:"type:date"` // Apenas para PF
	
	// Contato
	Email       string      `json:"email" gorm:"type:varchar(100);unique"`
	Telefone    string      `json:"telefone" gorm:"type:varchar(20);not null"`
	Celular     string      `json:"celular" gorm:"type:varchar(20)"`
	
	// Endereço
	Endereco    string      `json:"endereco" gorm:"type:varchar(200)"`
	Numero      string      `json:"numero" gorm:"type:varchar(10)"`
	Complemento string      `json:"complemento" gorm:"type:varchar(100)"`
	Bairro      string      `json:"bairro" gorm:"type:varchar(100)"`
	Cidade      string      `json:"cidade" gorm:"type:varchar(100)"`
	Estado      string      `json:"estado" gorm:"type:varchar(2)"`
	CEP         string      `json:"cep" gorm:"type:varchar(9)"`
	
	// Dados adicionais
	Observacoes string      `json:"observacoes" gorm:"type:text"`
	LimiteCredito float64   `json:"limite_credito" gorm:"type:decimal(10,2);default:0"`
	
	// Status
	Ativo       bool        `json:"ativo" gorm:"not null;default:true"`
	
	CreatedAt   time.Time   `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"not null"`
}

// NewClientePF cria uma nova instância de Cliente (Pessoa Física)
func NewClientePF(nome, cpf, rg string, dataNascimento time.Time, 
	telefone, email string) *Cliente {
	return &Cliente{
		ID:          uuid.New(),
		Tipo:        TipoPessoaFisica,
		Nome:        nome,
		CPFCNPJ:     cpf,
		RG:          rg,
		DataNascimento: dataNascimento,
		Telefone:    telefone,
		Email:       email,
		Ativo:       true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// NewClientePJ cria uma nova instância de Cliente (Pessoa Jurídica)
func NewClientePJ(nome, cnpj, telefone, email string) *Cliente {
	return &Cliente{
		ID:          uuid.New(),
		Tipo:        TipoPessoaJuridica,
		Nome:        nome,
		CPFCNPJ:     cnpj,
		Telefone:    telefone,
		Email:       email,
		Ativo:       true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// Validar verifica se o cliente é válido
func (c *Cliente) Validar() error {
	if c.Nome == "" {
		return ErrNomeObrigatorio
	}
	
	if c.CPFCNPJ == "" {
		return ErrCPFCNPJObrigatorio
	}
	
	if c.Telefone == "" {
		return ErrTelefoneObrigatorio
	}
	
	if c.Tipo == TipoPessoaFisica {
		if c.RG == "" {
			return ErrRGObrigatorio
		}
		if c.DataNascimento.IsZero() {
			return ErrDataNascimentoObrigatoria
		}
	}
	
	return nil
}

// AtualizarStatus atualiza o status do cliente
func (c *Cliente) AtualizarStatus(ativo bool) {
	c.Ativo = ativo
	c.UpdatedAt = time.Now()
}

// AtualizarLimiteCredito atualiza o limite de crédito do cliente
func (c *Cliente) AtualizarLimiteCredito(limite float64) {
	c.LimiteCredito = limite
	c.UpdatedAt = time.Now()
}

// AtualizarEndereco atualiza os dados de endereço do cliente
func (c *Cliente) AtualizarEndereco(endereco, numero, complemento, bairro, 
	cidade, estado, cep string) {
	c.Endereco = endereco
	c.Numero = numero
	c.Complemento = complemento
	c.Bairro = bairro
	c.Cidade = cidade
	c.Estado = estado
	c.CEP = cep
	c.UpdatedAt = time.Now()
}

// Erros de domínio
var (
	ErrCPFCNPJObrigatorio = NewDomainError("CPF/CNPJ é obrigatório")
	ErrDataNascimentoObrigatoria = NewDomainError("data de nascimento é obrigatória para pessoa física")
) 