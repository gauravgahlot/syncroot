BINARY=syncroot
PACKAGE=$(shell go list -m)
GOOS=$(shell go env GOOS)
GIT_VERSION=$(shell git describe --tags --always --dirty 2> /dev/null || echo 0.0.0)
export DOCKER_BUILDKIT=1

##@ Development
.PHONY: lint
lint: ## Format files and run lint in the repo.
	go mod tidy
	golangci-lint run ./... --verbose

.PHONY: lint-fix
lint-fix: ## Fix lint issues in the repo.
	golangci-lint run --fix ./... --verbose

.PHONY: test
test: ## Run tests.
	go test $(PACKAGE)/... -covermode=count -coverprofile=coverage.out -v

##@ Build
.PHONY: build
build: ## Build binary.
	go build -o bin/$(BINARY) cmd/main.go


.PHONY: docker-build
docker-build: build ## Build docker image.
	docker build -t $(IMG) .

##@ Support

.DEFAULT_GOAL := help
.PHONY: help
help: ## Show this help screen.
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
