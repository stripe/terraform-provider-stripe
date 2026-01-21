# Terraform Provider for Stripe

HOSTNAME=registry.terraform.io
NAMESPACE=stripe
NAME=stripe
BINARY=terraform-provider-${NAME}
VERSION=0.1.0

# Detect OS and architecture
OS=$(shell go env GOOS)
ARCH=$(shell go env GOARCH)

# Installation path for local development
INSTALL_PATH=~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS}_${ARCH}

default: install

.PHONY: build
build:
	go build -o ${BINARY}

.PHONY: install
install: build
	mkdir -p ${INSTALL_PATH}
	cp ${BINARY} ${INSTALL_PATH}/${BINARY}
	@echo ""
	@echo "✓ Provider installed to ${INSTALL_PATH}"
	@echo ""
	@echo "Add this to your Terraform configuration:"
	@echo ""
	@echo '  terraform {'
	@echo '    required_providers {'
	@echo '      stripe = {'
	@echo '        source  = "${HOSTNAME}/${NAMESPACE}/${NAME}"'
	@echo '        version = "${VERSION}"'
	@echo '      }'
	@echo '    }'
	@echo '  }'
	@echo ""

.PHONY: reinstall
reinstall: clean install

.PHONY: uninstall
uninstall:
	rm -rf ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}

.PHONY: clean
clean:
	rm -f ${BINARY}

.PHONY: test
test:
	go test ./... -v

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: deps
deps:
	go mod download
	go mod tidy

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  install    - Build and install the provider locally (default)"
	@echo "  reinstall  - Clean and reinstall the provider"
	@echo "  build      - Build the provider binary only"
	@echo "  uninstall  - Remove the locally installed provider"
	@echo "  clean      - Remove built binary"
	@echo "  test       - Run tests"
	@echo "  fmt        - Format Go code"
	@echo "  vet        - Run go vet"
	@echo "  deps       - Download and tidy dependencies"
	@echo "  help       - Show this help message"

