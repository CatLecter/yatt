package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"os/signal"
	"syscall"
	app "yatt/internal/application/iaa"
	iaaconfig "yatt/internal/config/iaa"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	config := iaaconfig.MustNew()

	ctx := context.Background()

	a := app.New(&ctx, &log, config)

	go a.GRPCSrv.MustRun()

	stopCh := make(chan os.Signal, 1)

	signal.Notify(stopCh, syscall.SIGTERM, syscall.SIGINT)

	stopSignal := <-stopCh

	log.Info().Msgf("Received signal: %s", stopSignal.String())

	a.GRPCSrv.Stop()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), config.App.StopTimeout)
	defer cancel()

	select {
	case <-timeoutCtx.Done():
		log.Info().Msg("gRPC server stopped")
	}
}
