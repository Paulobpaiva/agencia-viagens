# Sistema de Gerenciamento de Viagens

Sistema completo para gerenciamento de viagens de veículos, desenvolvido em Go seguindo os princípios da Clean Architecture.

## 🚀 Funcionalidades

- Cadastro completo de viagens (CRUD)
- Calendário de disponibilidade
- Mapa de rotas com integração de mapas
- Sistema de comunicação por e-mail
- Geração de contratos em PDF
- API RESTful para integração

## 🛠️ Tecnologias

- Go 1.21+
- PostgreSQL
- Docker
- Gin Framework
- GORM
- Railway.app (deploy)

## 📋 Pré-requisitos

- Go 1.21 ou superior
- Docker e Docker Compose
- PostgreSQL (opcional para desenvolvimento local)
- Make (opcional, para usar os comandos make)

## 🔧 Instalação

1. Clone o repositório:
```bash
git clone https://github.com/Paulobpaiva/agencia-viagens.git
cd agencia-viagens
```

2. Configure as variáveis de ambiente:
```bash
cp .env.example .env
# Edite o arquivo .env com suas configurações
```

3. Execute com Docker Compose:
```bash
docker-compose up -d
```

4. Execute as migrações:
```bash
make migrate-up
```

5. Inicie o servidor:
```bash
make run
```

## 🏗️ Estrutura do Projeto

```
.
├── cmd/                    # Ponto de entrada da aplicação
│   └── api/               # Servidor API
├── internal/              # Código interno da aplicação
│   ├── domain/           # Entidades e regras de negócio
│   ├── repository/       # Implementações dos repositórios
│   ├── usecase/         # Casos de uso da aplicação
│   └── delivery/        # Handlers HTTP
├── pkg/                  # Pacotes públicos
├── migrations/           # Migrações do banco de dados
├── docs/                # Documentação
└── scripts/             # Scripts utilitários
```

## 🧪 Testes

Execute os testes com:
```bash
make test
```

## 📦 Deploy

O projeto está configurado para deploy no Railway.app. Para fazer o deploy:

1. Crie uma conta no Railway.app
2. Conecte seu repositório
3. Configure as variáveis de ambiente necessárias
4. O Railway detectará automaticamente o Dockerfile e fará o deploy

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
