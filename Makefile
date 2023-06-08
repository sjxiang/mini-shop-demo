

proto:
	protoc ./go-grpc-api-gateway/pkg/**/pb/*.proto \
		--go_out=.
		--go-grpc_out=.


# protoc  --go_out=. --go-grpc_out=. .\go-grpc-api-gateway\pkg\auth\pb\auth.proto
# protoc  --go_out=. --go-grpc_out=. .\go-grpc-api-gateway\pkg\order\pb\order.proto
# protoc  --go_out=. --go-grpc_out=. .\go-grpc-api-gateway\pkg\product\pb\product.proto

gateway-server:
	go run ./go-grpc-api-gateway/cmd/main.go

