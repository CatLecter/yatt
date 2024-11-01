gen-proto:
	protoc api/v1/proto/* --go_out=./pkg/gen/go/ --go-grpc_out=./pkg/gen/go/

gen-swag:
	swag init -g ./cmd/bff/main.go -o api/v1/swagger

build-bff:
	go build -o ./build/bff ./cmd/bff/main.go

build-iaa:
	go build -o ./build/iaa ./cmd/iaa/main.go

build-all:
	go build -o ./build/bff ./cmd/bff/main.go
	go build -o ./build/iaa ./cmd/iaa/main.go

run-bff:
	./build/bff --config configs/iaa/local.yaml

run-iaa:
	docker compose up -d --build db
	./build/iaa --config configs/iaa/local.yaml

clean:
	rm ./build/*
