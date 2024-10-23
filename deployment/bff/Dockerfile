FROM golang:1.23-bookworm as builder

RUN mkdir "/tmp/build"

WORKDIR /tmp/build

COPY ../app app
COPY ../cmd cmd
COPY ../docs docs
COPY ../configs configs
COPY ../database database
COPY ../internal internal
COPY ../vendor vendor
COPY ../go.mod .
COPY ../go.sum .

RUN go mod tidy
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -o bff cmd/main.go

FROM debian:12-slim

WORKDIR /usr/local/bin/

COPY --from=builder /tmp/build/bff ./bff
COPY ../docs docs

CMD bff
