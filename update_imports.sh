#!/bin/bash

# Atualiza todas as importações nos arquivos .go
find . -type f -name "*.go" -exec sed -i 's|github.com/paulopaiva/agencia-viagens|agencia-viagens|g' {} +
find . -type f -name "*.go" -exec sed -i 's|github.com/Paulobpaiva/agencia-viagens|agencia-viagens|g' {} + 