gen-proto:
	protoc api/proto/v1/* --go_out=./pkg/gen/go/ --go-grpc_out=./pkg/gen/go/

gen-docs:
	swag init -g cmd/bff/main.go

build-bff:
	go build -o ./build/bff ./cmd/bff/main.go

build-iaa:
	go build -o ./build/iaa ./cmd/iaa/main.go

build-initializer:
	go build -o ./build/initializer ./cmd/initializer/main.go

build-all:
	go build -o ./build/bff ./cmd/bff/main.go
	go build -o ./build/iaa ./cmd/iaa/main.go
	go build -o ./build/initializer ./cmd/initializer/main.go

run-initializer:
	./build/initializer

run-bff:
	./build/bff

run-iaa:
	./build/iaa

clean:
	rm ./build/*
