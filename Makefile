#!/usr/bin/env make

.DEFAULT_GOAL  := help
.DEFAULT_SHELL := /bin/bash

GOARCH        := $(shell go env GOARCH)
GOOS          := $(shell go env GOOS)
GO            := $(GOBIN)/go
GOTEST        ?= $(shell command -v gotest 2>/dev/null)
GOLINT        ?= $(shell command -v golangci-lint 2>/dev/null)
PROJECT_MAIN  := $(shell find . -type f -name main.go 2>/dev/null)
#EXECUTABLE    ?= dist/simple-prober-$(GOOS)_$(GOARCH)
EXECUTABLE    ?= dist/simple-prober

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

$(EXECUTABLE): $(PROJECT_MAIN)
ifdef PROJECT_VERSION
	@$(GO) build -v -ldflags "-X 'main.Version=$(PROJECT_VERSION)'" -o $(EXECUTABLE) $<
else
	@$(GO) build -v -ldflags "-X 'main.Version=$(PROJECT_CHANGESET)'" -o $(EXECUTABLE) $<
endif

.PHONY: build
build: $(EXECUTABLE) ## Build the project

.PHONY: test
test: ## Unit tests
ifdef GOTEST
	@$(GOTEST) -v ./...
else
	@$(GO) test -v ./...
endif
