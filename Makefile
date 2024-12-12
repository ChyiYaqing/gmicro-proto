.PHONY: proto

proto: 
	@echo "generate order service code"
	protoc -I order \
		--go_out=golang/order \
		--go_opt=paths=source_relative \
		--go-grpc_out=golang/order \
		--go-grpc_opt=paths=source_relative \
		order/order.proto
	
	@echo "generate payment service code"
	protoc -I payment \
		--go_out=golang/payment \
		--go_opt=paths=source_relative \
		--go-grpc_out=golang/payment \
		--go-grpc_opt=paths=source_relative \
		payment/payment.proto

	@echo "generate shipping service code"
	protoc -I shipping \
		--go_out=golang/shipping \
		--go_opt=paths=source_relative \
		--go-grpc_out=golang/shipping \
		--go-grpc_opt=paths=source_relative \
		shipping/shipping.proto