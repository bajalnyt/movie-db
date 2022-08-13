run:
	go run ./cmd/api


lint: ## checking cleanup
	golangci-lint run
