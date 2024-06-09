PROJECT_NAME := zerome
PKG := "github.com/clwluvw/$(PROJECT_NAME)"
LINTER = golangci-lint
LINTER_VERSION = 1.59.0
CURRENT_LINTER_VERSION := $(shell golangci-lint version 2>/dev/null | awk '{ print $$4 }')

BUILD_TIME := $(shell LANG=en_US date +"%F_%T_%z")
COMMIT := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
VERSION ?= $(shell git describe --tags ${COMMIT} 2>/dev/null | cut -c2-)
VERSION := $(or $(VERSION),$(COMMIT))
LD_FLAGS ?=
LD_FLAGS += -X main.Version=$(VERSION)
LD_FLAGS += -X main.Revision=$(COMMIT)
LD_FLAGS += -X main.Branch=$(BRANCH)
LD_FLAGS += -X main.BuildDate=$(BUILD_TIME)

.PHONY: all dep lint build unused

all: dep build

dep: ## Get the dependencies
	@go mod tidy

lintdeps: ## golangci-lint dependencies
ifneq ($(CURRENT_LINTER_VERSION), $(LINTER_VERSION))
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v$(LINTER_VERSION)
endif

lint: lintdeps ## to lint the files
	$(LINTER) run --config=.golangci.yml ./...

build: ## Build the binary file
	@go build -v -ldflags="$(LD_FLAGS)" $(PKG)/cmd/$(PROJECT_NAME)

test:
	@go test -v -cover -race ./...

unused: dep
	@echo ">> running check for unused/missing packages in go.mod"
	@git diff --exit-code -- go.sum go.mod

