FROM golang:1.23-bookworm as builder

RUN mkdir "/tmp/build"

WORKDIR /tmp/build

COPY .. .

RUN go mod tidy
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -o iaa cmd/main.go

FROM debian:12-slim

WORKDIR /usr/local/bin/

COPY --from=builder /tmp/build/identity ./identity

CMD identity
