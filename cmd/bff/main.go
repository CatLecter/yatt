package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"os/signal"
	"syscall"
	"yatt/internal/application/bff"
	bffConfig "yatt/internal/config/bff"
	grpcClient "yatt/internal/infrastructure/clients/bff/grpc"
	"yatt/internal/presentation/bff/handlers"
	service "yatt/internal/services/bff"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	config := bffConfig.MustNew()

	iaaClient, err := grpcClient.New(&log, config.Iaa.URI, config.Iaa.RetryTimeout, config.Iaa.Retries)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create gRPC iaa client")
	}

	srv := service.New(&log, iaaClient)

	handler := handlers.New(srv, &log)

	bffApp := bff.New(
		handler,
		&log,
		config.App.Host,
		config.App.Port,
		config.App.ReadTimeout,
		config.App.WriteTimeout,
		config.App.AllowOrigins,
	)

	go bffApp.MustRun()

	stopCh := make(chan os.Signal, 1)

	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	stopSignal := <-stopCh

	log.Info().Msgf("Received signal: %s", stopSignal.String())

	ctx, cancel := context.WithTimeout(context.Background(), config.App.StopTimeout)

	defer cancel()

	if err := bffApp.Stop(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server Shutdown")
	}

	select {
	case <-ctx.Done():
		log.Info().Msg("Server stopped")
	}
}
