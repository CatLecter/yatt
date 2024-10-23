package main

import (
	"context"
	"github.com/CatLecter/yatt/configs"
	"github.com/CatLecter/yatt/internal/application/bff"
	identitygrpc "github.com/CatLecter/yatt/internal/infrastructure/clients/identity/grpc"
	"github.com/CatLecter/yatt/internal/presentation/bff/handlers"
	"github.com/CatLecter/yatt/internal/services"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg := configs.NewConfig()

	identityClient, err := identitygrpc.New(&log, "127.0.0.1:5080", time.Duration(300)*time.Millisecond, 3)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create gRPC identity client")
	}

	service := services.New(identityClient, &log)

	handler := handlers.New(service, &log)

	app := bff.New(handler, &log, cfg.Host, cfg.Port)

	go app.MustRun()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	sgn := <-quit

	log.Info().Msgf("Received signal: %s", sgn.String())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	if err := app.Stop(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server Shutdown")
	}

	select {
	case <-ctx.Done():
		log.Info().Msg("Server stopped")
	}
}
