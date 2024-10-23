package main

import (
	"context"
	"github.com/CatLecter/yatt/internal/application/iaa"
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

	app := iaa.New(
		&log,
		5080,
		"postgresql://admin:admin@127.0.0.1:5432/db?sslmode=disable",
		600,
	)

	go app.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sgn := <-stop

	log.Info().Msgf("Received signal: %s", sgn.String())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	app.GRPCSrv.Stop()

	select {
	case <-ctx.Done():
		log.Info().Msg("gRPC server stopped")
	}
}
