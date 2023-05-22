.PHONY: proto
gen-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  api/dropper/filedrop.proto

.PHONY: run
run:
	go run cmd/dropper/main.go
