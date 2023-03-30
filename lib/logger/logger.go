package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var logger zerolog.Logger

func New(level string) zerolog.Logger {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	logLevel, err := zerolog.ParseLevel(level)

	if err != nil {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	zerolog.SetGlobalLevel(logLevel)

	logger = zerolog.New(os.Stdout).
		With().Timestamp().CallerWithSkipFrameCount(2).Logger()

	zerolog.DefaultContextLogger = &logger

	return logger
}
