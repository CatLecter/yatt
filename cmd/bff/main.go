package main

import (
	"context"
	"github.com/CatLecter/yatt/internal/application/bff"
	bffconfig "github.com/CatLecter/yatt/internal/config/bff"
	iaagrpcclient "github.com/CatLecter/yatt/internal/infrastructure/clients/bff/grpc"
	"github.com/CatLecter/yatt/internal/presentation/bff/handlers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	config := bffconfig.MustNew()

	iaaClient, err := iaagrpcclient.New(&log, config.Iaa.URI, config.Iaa.RetryTimeout, config.Iaa.Retries)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create gRPC iaa client")
	}

	service := user.New(iaaClient, &log)

	handler := handlers.New(service, &log)

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
