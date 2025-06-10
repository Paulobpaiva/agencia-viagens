# Stage 1: Builder
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Instala dependências necessárias
RUN apk add --no-cache \
    git \
    make \
    gcc \
    libc-dev

# Copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte
COPY . .

# Compila a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Stage 2: Runner
FROM alpine:3.19

WORKDIR /app

# Instala certificados CA
RUN apk add --no-cache ca-certificates tzdata

# Copia o binário compilado
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations

# Expõe a porta da aplicação
EXPOSE 8080

# Define variáveis de ambiente padrão
ENV APP_ENV=production

# Comando para iniciar a aplicação
CMD ["./main"] 