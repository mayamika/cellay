.SUFFIXES:

SHELL	?= /bin/bash
CURL	?= curl --silent
GO		?= go

# Directories

BIN_OUTPUT_DIR 		?= bin
DOWNLOAD_CACHE_DIR	?= build/download
TOOLS_BIN_DIR		:= tools/bin
PROTO_DIR			:= proto

DEP_DIRS = $(BIN_OUTPUT_DIR) $(DOWNLOAD_CACHE_DIR) $(TOOLS_BIN_DIR)
$(DEP_DIRS):
	-mkdir -p $(DEP_DIRS)

# Tools

export TOOLS_PATH = $(CURDIR)/$(TOOLS_BIN_DIR)

TOOLS_MODFILE := tools/go.mod
define install-go-tool =
	$(GO) build \
		-o $(TOOLS_BIN_DIR) \
		-ldflags "-s -w" \
		-modfile $(TOOLS_MODFILE)
endef

clean-tools:
	-rm -rf $(TOOLS_BIN_DIR)

PROTOC_VERSION := 3.14.0
PROTOC_DIST_LINK := https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-linux-x86_64.zip

PROTOC_DIST_ARCHIVE = $(DOWNLOAD_CACHE_DIR)/protoc-$(PROTOC_VERSION).zip
PROTOC_DIST = $(DOWNLOAD_CACHE_DIR)/protoc-$(PROTOC_VERSION)

$(PROTOC_DIST_ARCHIVE): | $(DOWNLOAD_CACHE_DIR)
	$(CURL) --output "$@" -L "$(PROTOC_DIST_LINK)"

$(PROTOC_DIST): | $(PROTOC_DIST_ARCHIVE)
	unzip -d "$@" -o $|

## Codegen tools

CODEGEN_TOOLS := $(TOOLS_BIN_DIR)/enumer
$(TOOLS_BIN_DIR)/enumer: | $(TOOLS_BIN_DIR)
	$(install-go-tool) github.com/alvaroloes/enumer

## CI tools

GO_LINT_TOOL = $(TOOLS_BIN_DIR)/golangci-lint
$(GO_LINT_TOOL): | $(TOOLS_BIN_DIR)
	$(install-go-tool) github.com/golangci/golangci-lint/cmd/golangci-lint

## Proto tools

PROTOC_TOOL = $(TOOLS_BIN_DIR)/protoc-$(PROTOC_VERSION)
PROTO_TOOLS = $(PROTOC_TOOL)
$(PROTOC_TOOL): | $(PROTOC_DIST)
	cp $(PROTOC_DIST)/bin/protoc $@

PROTO_TOOLS += $(TOOLS_BIN_DIR)/protoc-gen-go
$(TOOLS_BIN_DIR)/protoc-gen-go: | $(TOOLS_BIN_DIR)
	$(install-go-tool) google.golang.org/protobuf/cmd/protoc-gen-go

PROTO_TOOLS += $(TOOLS_BIN_DIR)/protoc-gen-go-grpc
$(TOOLS_BIN_DIR)/protoc-gen-go-grpc: | $(TOOLS_BIN_DIR)
	$(install-go-tool) google.golang.org/grpc/cmd/protoc-gen-go-grpc

PROTO_TOOLS += $(TOOLS_BIN_DIR)/protoc-gen-grpc-gateway
$(TOOLS_BIN_DIR)/protoc-gen-grpc-gateway: | $(TOOLS_BIN_DIR)
	$(install-go-tool) github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

PROTO_TOOLS += $(TOOLS_BIN_DIR)/protoc-gen-openapiv2
$(TOOLS_BIN_DIR)/protoc-gen-openapiv2: | $(TOOLS_BIN_DIR)
	$(install-go-tool) github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

PROTO_TOOLS += $(TOOLS_BIN_DIR)/protoc-gen-buf-lint
$(TOOLS_BIN_DIR)/protoc-gen-buf-lint: | $(TOOLS_BIN_DIR)
	$(install-go-tool) github.com/bufbuild/buf/cmd/protoc-gen-buf-lint

# Proto

PROTO_PROTOBUF_FLAGS = \
		-I$(PROTO_DIR) \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(PROTO_DIR) \
		--go_opt=paths=source_relative \
		--buf-lint_out=$(PROTO_DIR) \
		--plugin=protoc-gen-go=$(TOOLS_BIN_DIR)/protoc-gen-go \
		--plugin=protoc-gen-buf-lint=$(TOOLS_BIN_DIR)/protoc-gen-buf-lint

PROTO_GRPC_FLAGS = \
		--go-grpc_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_DIR) \
		--plugin=protoc-gen-go-grpc=$(TOOLS_BIN_DIR)/protoc-gen-go-grpc

PROTO_GRPC_GATEWAY_FLAGS = \
		--grpc-gateway_out=allow_patch_feature=true,logtostderr=true,request_context=true:$(PROTO_DIR) \
		--grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=allow_merge=true:$(PROTO_OPENAPI_DIR) \
		--plugin=protoc-gen-openapiv2=$(TOOLS_BIN_DIR)/protoc-gen-openapiv2 \
		--plugin=protoc-gen-grpc-gateway=$(TOOLS_BIN_DIR)/protoc-gen-grpc-gateway

define generate-proto-base =
	$(PROTOC_TOOL) $(PROTO_PROTOBUF_FLAGS)
endef

define generate-proto-gw =
	$(generate-proto-base) $(PROTO_GRPC_FLAGS) $(PROTO_GRPC_GATEWAY_FLAGS) $(PROTO_FILES)
endef

define generate-proto-grpc =
	$(generate-proto-base) $(PROTO_GRPC_FLAGS) $(PROTO_FILES)
endef

define generate-proto-protobuf =
	$(generate-proto-base) $(PROTO_FILES)
endef

PROTO_CELLAY_SERVER_FILES = $(addprefix $(PROTO_DIR)/cellay/v1/, \
	cellay.proto \
)
PROTO_CELLAY_SERVER_OPENAPI_DIR := cellay-server/internal/httpserver
PROTO_CELLAY_SERVER_GENERATED = $(PROTO_CELLAY_SERVER_OPENAPI_DIR)/apidocs.swagger.json
PROTO_CELLAY_SERVER_GENERATED += $(addprefix $(PROTO_DIR)/cellay/v1/cellay,\
				_grpc.pb.go \
				.pb.go \
				.pb.gw.go)

.INTERMEDIATE: .proto_cellay_server_generated
.proto_cellay_server_generated: $(PROTO_CELLAY_SERVER_FILES) $(PROTO_TOOLS)
	$(generate-proto-gw)

$(PROTO_CELLAY_SERVER_GENERATED): PROTO_OPENAPI_DIR = $(PROTO_CELLAY_SERVER_OPENAPI_DIR)
$(PROTO_CELLAY_SERVER_GENERATED): PROTO_FILES = $(PROTO_CELLAY_SERVER_FILES)
$(PROTO_CELLAY_SERVER_GENERATED): .proto_cellay_server_generated ;

## Generate

.PHONY: generate-proto-cellay-server
generate-proto-cellay-server: $(PROTO_CELLAY_SERVER_GENERATED)

.PHONY: generate-proto
generate-proto: generate-proto-cellay-server

## Cleanup

.PHONY: clean-proto-cellay-server
clean-proto-cellay-server:
	-rm $(PROTO_CELLAY_SERVER_GENERATED)

.PHONY: clean-proto
clean-proto: clean-proto-cellay-server

# Lint

.PHONY: lint
lint: $(GO_LINT_TOOL)
	$(GO_LINT_TOOL) run --sort-results

# Generate

.PHONY: generate-go
generate-go: $(CODEGEN_TOOLS)
	$(GO) generate ./...

# Build

GO_BUILD_TARGET_CELLAY_SERVER := $(BIN_OUTPUT_DIR)/cellay-server
GO_BUILD_TARGETS = $(GO_BUILD_TARGET_CELLAY_SERVER)

.PHONY: $(GO_BUILD_TARGETS)
$(GO_BUILD_TARGETS): | $(BIN_OUTPUT_DIR)
	$(GO) build -o $(BIN_OUTPUT_DIR) ./cmd/...

# Test

.PHONY: test
test:
	$(GO) test -v -race ./...

.PHONY: test-norace
test-norace:
	$(GO) test -v ./...

# Run

.PHONY: run-server
run-server: $(GO_BUILD_TARGET_CELLAY_SERVER)
	$(BIN_OUTPUT_DIR)/cellay-server

# Cleanup

.PHONY: clean
clean: clean-proto clean-tools
