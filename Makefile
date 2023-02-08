#!/usr/bin/env make

.DEFAULT_GOAL  := help
.DEFAULT_SHELL := /bin/bash

GOARCH                    := $(shell go env GOARCH)
GOOS                      := $(shell go env GOOS)
GO                        := $(GOBIN)/go
GOTEST                    ?= $(shell command -v gotest 2>/dev/null)
GOLINT                    ?= $(shell command -v golangci-lint 2>/dev/null)
EXECUTABLE                ?= dist/simple-prober
#EXECUTABLE               ?= dist/simple-prober-$(GOOS)_$(GOARCH)

DOCKER_REGISTRY           ?= docker.io
DOCKER_REGISTRY_NAMESPACE ?= jjuarez
DOCKER_SERVICE_NAME       ?= simple-prober
DOCKER_IMAGE              := $(DOCKER_REGISTRY)/$(DOCKER_REGISTRY_NAMESPACE)/$(DOCKER_SERVICE_NAME)

PROJECT_CHANGESET := $(shell git rev-parse --verify HEAD 2>/dev/null)


define assert-set
	@$(if $($1),,$(error $(1) environment variable is not defined))
endef

define assert-command
	@$(if $(shell command -v $1 2>/dev/null),,$(error $(1) command not found))
endef

define assert-file
	@$(if $(wildcard $($1) 2>/dev/null),,$(error $($1) does not exist))
endef


.PHONY: help
help: ## Shows this pretty help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target>\n\nTargets:\n"} /^[a-zA-Z\/_-]+:.*?##/ { printf " %-20s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: lint
lint: ## Lint the source code
	$(call assert-command,$(GOLINT))
	@$(GOLINT) run ./...

.PHONY: clean
clean: ## Clean the project executable
	@$(GO) clean -v ./...
	@rm -f $(EXECUTABLE)

.PHONY: build
build: ## Build the project
ifdef PROJECT_VERSION
	@$(GO) build -v -ldflags="-X github.com/jjuarez/simple-prober/cmd.Version='v$(PROJECT_VERSION)+$(PROJECT_CHANGESET)'" -o $(EXECUTABLE) main.go
else
	@$(GO) build -v -ldflags="-X github.com/jjuarez/simple-prober/cmd.Version='$(PROJECT_CHANGESET)'" -o $(EXECUTABLE) main.go
endif

.PHONY: test
test: ## Unit tests
ifdef GOTEST
	@$(GOTEST) -v ./...
else
	@$(GO) test -v ./...
endif

.PHONY: docker/login
docker/login:
	$(call assert-set,DOCKER_USERNAME)
	$(call assert-set,DOCKER_TOKEN)
	@echo $(DOCKER_TOKEN)|docker login --username $(DOCKER_USERNAME) --password-stdin $(DOCKER_REGISTRY)

.PHONY: docker/build
docker/build: docker/login ## Makes the Docker build and takes care of the remote cache by target
ifdef PROJECT_VERSION
	@docker image build \
    --build-arg BUILDKIT_INLINE_CACHE=1 \
    --build-arg VERSION=v$(PROJECT_VERSION)+$(PROJECT_CHANGESET) \
    --cache-from $(DOCKER_IMAGE):latest \
    --tag $(DOCKER_IMAGE):$(PROJECT_CHANGESET) \
    --tag $(DOCKER_IMAGE):latest \
    --file Dockerfile \
    .
else
	@docker image build \
    --build-arg BUILDKIT_INLINE_CACHE=1 \
    --build-arg VERSION=$(PROJECT_CHANGESET) \
    --cache-from $(DOCKER_IMAGE):latest \
    --tag $(DOCKER_IMAGE):$(PROJECT_CHANGESET) \
    --tag $(DOCKER_IMAGE):latest \
    --file Dockerfile \
    .
endif
	@docker image push $(DOCKER_IMAGE):latest

.PHONY: docker/release
docker/release: docker/build ## Builds and release over the Docker registry the image
	@docker image push $(DOCKER_IMAGE):$(PROJECT_CHANGESET)
ifdef PROJECT_VERSION
	@docker image tag  $(DOCKER_IMAGE):$(PROJECT_CHANGESET) $(DOCKER_IMAGE):$(PROJECT_VERSION)
	@docker image push $(DOCKER_IMAGE):$(PROJECT_VERSION)
else
	$(warning The release rule should have a PROJECT_VERSION defined)
endif
