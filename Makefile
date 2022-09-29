GEN_DIR='./pb'
OPENV2_DIR='./swagger'

.PHONY: server
server:
	@go run cmd/server/main.go

.PHONY: client
client:
	@go run cmd/client/main.go

.PHONY: gen
gen: genclean
	mkdir -p $(GEN_DIR) && \
	mkdir -p $(OPENV2_DIR) && \
	protoc -I ./proto --go_out=$(GEN_DIR) --go-grpc_out=$(GEN_DIR) --grpc-gateway_out=$(GEN_DIR) --openapiv2_out=$(OPENV2_DIR) ./proto/*.proto

# .PHONY: gengateway
# gengateway: gen
# 	protoc -I ./proto --grpc-gateway_out=$(GEN_DIR) --grpc-gateway_opt grpc_api_configuration=./proto/trip.yaml ./proto/trip.proto

.PHONY: genclean
genclean:
	rm -rf $(GEN_DIR) && \
	rm -rf $(OPENV2_DIR)