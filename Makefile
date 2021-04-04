.PHONY: all build lint test vet
CHECK_FILES?=$$(go list ./... | grep -v /vendor/)
APP_NAME=git-fork

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

all: quality build ## Run the tests and build the binary.

quality: format lint vet test ## Run quality operations

mocks: ## Generate mock for the application
	docker run --user=$$(id -u):$$(id -g) -v $$PWD:/app -w /app vektra/mockery --all --recursive --keeptree --inpackage

build: ## Build the binary.
	go build -o bin/$(APP_NAME) *.go

lint: ## Lint the code.
	golint -set_exit_status $(CHECK_FILES)

test: test-unit test-integration

test-unit: ## Run unit tests.
	go test -race -cover -p 1 -coverprofile=c.out

test-integration: ## Run unit and integration tests
	go test ./... -race -tags integration -cover -p 1 -coverprofile=c.out

vet: ## Vet the code
	go vet $(CHECK_FILES)

format: ## Format the code
	go fmt ./...

run: ## Run application
	go run .

install: ## Install application on local machine or container
	go install *.go