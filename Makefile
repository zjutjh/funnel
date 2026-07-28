API_DIR := apps/funnel-api
API_FILE := $(API_DIR)/funnel.api
RPC_DIR := apps/funnel-rpc
RPC_PROTO := $(RPC_DIR)/funnel.proto
MODULE := funnel
STYLE ?= go_zero

.PHONY: build
build:
	go build -o bin/funnel .

.PHONY: build-api
build-api:
	go build -o bin/funnel-api ./$(API_DIR)

.PHONY: build-rpc
build-rpc:
	go build -o bin/funnel-rpc ./$(RPC_DIR)

.PHONY: generate
generate: generate-api generate-rpc

.PHONY: generate-api
generate-api:
	goctl api go --api $(API_FILE) --dir $(API_DIR) --style $(STYLE)

.PHONY: generate-rpc
generate-rpc:
	goctl rpc protoc $(RPC_PROTO) --go_out=$(RPC_DIR) --go-grpc_out=$(RPC_DIR) --zrpc_out=$(RPC_DIR) --style $(STYLE) --module $(MODULE)

.PHONY: configure
configure: tools

.PHONY: tools
tools:
	go install github.com/zeromicro/go-zero/tools/goctl@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: fmt
fmt:
	goctl api format --dir $(API_DIR)
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	go run .

.PHONY: run-api
run-api:
	go run ./$(API_DIR)

.PHONY: run-rpc
run-rpc:
	go run ./$(RPC_DIR)
