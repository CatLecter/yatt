package handlers

import (
	"github.com/rs/zerolog"
	service "yatt/internal/services/bff"
)

type Handler struct {
	service *service.Service
	log     *zerolog.Logger
}

func New(service *service.Service, log *zerolog.Logger) *Handler {
	return &Handler{service: service, log: log}
}
