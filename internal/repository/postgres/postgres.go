package postgres

import (
	"context"
	"fmt"
	"time"

	"agencia-viagens/internal/config"
	"agencia-viagens/internal/domain"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgresDB cria uma nova conexão com o banco de dados PostgreSQL
func NewPostgresDB(cfg config.DatabaseConfig) (*gorm.DB, error) {
	// Configuração do GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	// Conecta ao banco de dados
	db, err := gorm.Open(postgres.Open(cfg.GetDSN()), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}

	// Configura o pool de conexões
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter instância do banco de dados: %v", err)
	}

	// Configurações do pool de conexões
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Executa migrações automáticas
	if err := autoMigrate(db); err != nil {
		return nil, fmt.Errorf("erro ao executar migrações: %v", err)
	}

	return db, nil
}

// autoMigrate executa as migrações automáticas do GORM
func autoMigrate(db *gorm.DB) error {
	// Lista de modelos para migração
	models := []interface{}{
		&domain.Viagem{},
		&domain.Veiculo{},
		&domain.Motorista{},
		&domain.Cliente{},
	}

	// Executa as migrações
	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("erro ao executar migrações automáticas: %v", err)
	}

	return nil
}

// TransactionManager implementa o gerenciador de transações
type TransactionManager struct {
	db *gorm.DB
}

// NewTransactionManager cria uma nova instância do gerenciador de transações
func NewTransactionManager(db *gorm.DB) *TransactionManager {
	return &TransactionManager{db: db}
}

// WithTransaction executa uma função dentro de uma transação
func (tm *TransactionManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := tm.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Cria um novo contexto com a transação
	txCtx := context.WithValue(ctx, "tx", tx)

	// Executa a função dentro da transação
	if err := fn(txCtx); err != nil {
		tx.Rollback()
		return err
	}

	// Commit da transação
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// GetTxFromContext retorna a transação do contexto
func GetTxFromContext(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
		return tx
	}
	return nil
}
