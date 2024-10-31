package app

import (
	"context"
	grpcapp "github.com/CatLecter/yatt/internal/application/iaa/grpc"
	iaaconfig "github.com/CatLecter/yatt/internal/config/iaa"
	"github.com/rs/zerolog"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(ctx *context.Context, log *zerolog.Logger, config *iaaconfig.Config) *App {
	return &App{GRPCSrv: grpcapp.New(ctx, log, config)}
}
