package handlers

import (
	"github.com/CatLecter/yatt/internal/services"
	"github.com/rs/zerolog"
)

type Handler struct {
	service *services.Service
	log     *zerolog.Logger
}

func New(service *services.Service, log *zerolog.Logger) *Handler {
	return &Handler{service: service, log: log}
}
