GEN_DIR='./pb'

.PHONY: server
server:
	@go run cmd/server/main.go

.PHONY: client
client:
	@go run cmd/client/main.go

.PHONY: gen
gen: genclean
	mkdir -p $(GEN_DIR) && \
	protoc -I ./proto --go_out=$(GEN_DIR) --go-grpc_out=$(GEN_DIR) ./proto/*.proto

.PHONY: genclean
genclean:
	rm -rf $(GEN_DIR)