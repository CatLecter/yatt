package handlers

import (
	grpcclient "github.com/CatLecter/yatt/internal/infrastructure/clients/bff/grpc"
	"github.com/rs/zerolog"
)

type Handler struct {
	client *grpcclient.Client
	log    *zerolog.Logger
}

func New(client *grpcclient.Client, log *zerolog.Logger) *Handler {
	return &Handler{client: client, log: log}
}
