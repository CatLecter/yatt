package app

import (
	"context"
	"github.com/rs/zerolog"
	grpcapp "yatt/internal/application/iaa/grpc"
	iaaconfig "yatt/internal/config/iaa"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(ctx *context.Context, log *zerolog.Logger, config *iaaconfig.Config) *App {
	return &App{GRPCSrv: grpcapp.New(ctx, log, config)}
}
