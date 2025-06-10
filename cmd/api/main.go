package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/Paulobpaiva/agencia-viagens/docs" // Importa a documentação gerada
	"github.com/Paulobpaiva/agencia-viagens/internal/config"
	"github.com/Paulobpaiva/agencia-viagens/internal/delivery/http"
	"github.com/Paulobpaiva/agencia-viagens/internal/repository"
	"github.com/Paulobpaiva/agencia-viagens/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API de Agência de Viagens
// @version         1.0
// @description     API para gerenciamento de viagens, veículos e motoristas.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Suporte
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func init() {
	// Carrega variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}
}

func main() {
	// Configuração do ambiente
	env := os.Getenv("APP_ENV")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Inicializa configurações
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Inicializa conexão com o banco de dados
	db, err := repository.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Inicializa repositórios
	viagemRepo := repository.NewViagemRepository(db)
	veiculoRepo := repository.NewVeiculoRepository(db)
	motoristaRepo := repository.NewMotoristaRepository(db)

	// Inicializa casos de uso
	viagemUseCase := usecase.NewViagemUseCase(viagemRepo, veiculoRepo, motoristaRepo)
	veiculoUseCase := usecase.NewVeiculoUseCase(veiculoRepo)
	motoristaUseCase := usecase.NewMotoristaUseCase(motoristaRepo)

	// Inicializa handlers HTTP
	handler := http.NewHandler(viagemUseCase, veiculoUseCase, motoristaUseCase)

	// Configura o router
	router := gin.Default()

	// Adiciona documentação Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Inicializa rotas da API
	handler.InitRoutes(router)

	// Inicia o servidor
	addr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	log.Printf("Servidor iniciado em %s", addr)
	log.Printf("Documentação Swagger disponível em http://%s/swagger/index.html", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
