API_FILE := api/rest/gateway.api
PROTO_FILE := api/rpc/worker.proto

GATEWAY_DIR := gateway
WORKER_DIR := worker
RPC_DIR := api/rpc

.PHONY: gen gen-api gen-rpc clean clean-api clean-rpc

gen: gen-api gen-rpc

gen-api:
	goctl api go --api $(API_FILE) --dir $(GATEWAY_DIR)

gen-rpc:
	goctl rpc protoc $(PROTO_FILE) \
		--go_out=./$(RPC_DIR) \
		--go-grpc_out=./$(RPC_DIR) \
		--zrpc_out=./$(WORKER_DIR) \

clean: clean-api clean-rpc


clean-api:
	rm -rf $(GATEWAY_DIR)

clean-rpc:
	rm -rf $(RPC_DIR)/client
	rm -rf $(RPC_DIR)/worker
	rm -rf $(RPC_DIR)/funnel
	rm -rf .gen
	rm -rf $(WORKER_DIR)/client
	rm -rf $(WORKER_DIR)/etc
	rm -rf $(WORKER_DIR)/internal/logic
	rm -rf $(WORKER_DIR)/internal/server
	rm -f $(WORKER_DIR)/worker.go
