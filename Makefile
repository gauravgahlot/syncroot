BINARY=syncroot
PACKAGE=$(shell go list -m)
GOOS=$(shell go env GOOS)
GIT_VERSION=$(shell git describe --tags --always --dirty 2> /dev/null || echo 0.0.0)
LDFLAGS=-ldflags "-w -X $(PACKAGE)/internal/info.Version=$(GIT_VERSION)"
export DOCKER_BUILDKIT=1

##@ Development
.PHONY: up
up: ## Starts the app using docker compose.
	PACKAGE=$(PACKAGE) GIT_VERSION=$(GIT_VERSION) docker compose up --build --abort-on-container-exit

.PHONY: down
down: ## Stops the app using docker compose.
	docker compose down

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
	go build $(LDFLAGS) -o bin/$(BINARY) cmd/$(BINARY)/main.go

# .PHONY: docker-build
# docker-build: test ## Build docker image with the manager.
# 	docker build \
# 		--platform linux/amd64 \
# 		--target builder \
# 		--cache-from "$(shell img1=$(IMG); echo $${img1%:*}):builder" \
# 		--build-arg BUILDKIT_INLINE_CACHE=1 \
# 		--ssh default=$(SSH_KEY_PATH) \
# 		--build-arg=GIT_VERSION=$(GIT_REF) \
# 		--tag "$(shell img1=$(IMG); echo $${img1%:*}):builder" \
# 		--file Dockerfile \
# 		.

# 	docker build \
# 		--platform linux/amd64 \
# 		--cache-from "$(shell img1=$(IMG); echo $${img1%:*}):builder" \
# 		--cache-from "${IMG}" \
# 		--build-arg BUILDKIT_INLINE_CACHE=1 \
# 		--ssh default=$(SSH_KEY_PATH) \
# 		--build-arg=GIT_VERSION=$(GIT_REF) \
# 		--tag "${IMG}" \
# 		--file Dockerfile \
# 		.

# .PHONY: docker-push
# docker-push: ## Push docker image with the manager.
# 	kind load docker-image ${IMG} --name task-system

##@ Support

.DEFAULT_GOAL := help
.PHONY: help
help: ## Show this help screen.
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
