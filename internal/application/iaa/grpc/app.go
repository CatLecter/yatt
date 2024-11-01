package grpcapp

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
	iaaconfig "yatt/internal/config/iaa"
	"yatt/internal/database"
	repository "yatt/internal/infrastructure/repositories/iaa"
	usergrpc "yatt/internal/presentation/iaa/grpc"
	service "yatt/internal/services/iaa"
)

type App struct {
	log        *zerolog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(ctx *context.Context, log *zerolog.Logger, config *iaaconfig.Config) *App {
	gRPCServer := grpc.NewServer()
	db := database.MustNew(
		*ctx,
		config.Storage.URI,
		config.Storage.MaxConnections,
		config.Storage.MinConnections,
		config.Storage.MaxConnLifetime,
		config.Storage.MaxConnIdleTime,
	)
	repo := repository.New(db, log)
	userService := service.New(log, repo.UserStorage)
	usergrpc.New(gRPCServer, userService)
	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       config.App.Port,
	}
}

func (a *App) Run() error {
	a.log.Info().Msg("Starting gRPC server...")
	var listener net.Listener
	var err error
	if listener, err = net.Listen("tcp", fmt.Sprintf(":%d", a.port)); err != nil {
		return err
	}
	a.log.Info().Msg("gRPC server is running: " + listener.Addr().String())
	if err = a.gRPCServer.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	a.log.Info().Msg("Stopping gRPC server...")
	a.gRPCServer.GracefulStop()
}
