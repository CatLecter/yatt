package main

import (
	"context"
	app "github.com/CatLecter/yatt/internal/application/iaa"
	iaaconfig "github.com/CatLecter/yatt/internal/config/iaa"
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

	config := iaaconfig.MustNew()

	ctx, cancel := context.WithTimeout(context.Background(), config.App.StopTimeout)

	defer cancel()

	a := app.New(&ctx, &log, config)

	go a.GRPCSrv.MustRun()

	stopCh := make(chan os.Signal, 1)

	signal.Notify(stopCh, syscall.SIGTERM, syscall.SIGINT)

	stopSignal := <-stopCh

	log.Info().Msgf("Received signal: %s", stopSignal.String())

	a.GRPCSrv.Stop()

	select {
	case <-ctx.Done():
		log.Info().Msg("gRPC server stopped")
	}
}
