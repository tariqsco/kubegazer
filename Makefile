# Kubegazer Makefile
# Author: Tariq Scott
# Date: 2026-07-27
# Description: Convenience targets for local development, testing and building. 

.PHONY: run build test vet

# Default target executed when running 'make'
.DEFAULT_GOAL := help

help: ## Display available commands
	@echo "Kubegazer Developer Commands:"
	@grep -E '[a-zA-Z_-]+:.*?'

run: ## Run the backend application locally 
	go run ./cmd/kubegazer/main.go

build: ## Compile the application binary into ./bin/kubegazer
	go build -o bin/kubegazer ./cmd/main.go

test: ## Run unit test across all packages
	go test ./...

vet: ## Analyze code for potential corrrectness and performance bugs 
	go vet ./...
