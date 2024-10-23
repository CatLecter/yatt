package iaa

import (
	grpcapp "github.com/CatLecter/yatt/internal/application/iaa/grpc"
	"github.com/rs/zerolog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(logger *zerolog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	//TODO: инициализировать хранилище
	//TODO: инициализировать сервер
	gRPCSrv := grpcapp.New(logger, grpcPort)

	return &App{GRPCSrv: gRPCSrv}
}
