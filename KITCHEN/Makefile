run-orders:
	@go run services/orders/*.go

run-kitchen:
	@go run services/kitchen/*.go

PROTO_DIR=protobuf
OUT_DIR=services/common/genproto/orders
PROTO_FILE=$(PROTO_DIR)/orders.proto

gen:
	mkdir -p $(OUT_DIR)
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILE)
