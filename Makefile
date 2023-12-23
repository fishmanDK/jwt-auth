start-grpc:
	go run cmd/jwt-auth/grpc/main.go

start:
	go run cmd/jwt-auth/http/main.go


generate-proto:
	protoc --go_out=internal/convert-proto --go_opt=paths=source_relative \
    		--go-grpc_out=internal/convert-proto --go-grpc_opt=paths=source_relative \
    		proto/auth.proto