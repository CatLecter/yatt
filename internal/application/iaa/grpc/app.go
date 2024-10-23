package grpcapp

import (
	"fmt"
	usergrpc "github.com/CatLecter/yatt/internal/presentation/iaa/grpc/user"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	log        *zerolog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *zerolog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()
	usergrpc.NewUserServer(gRPCServer)
	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
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
