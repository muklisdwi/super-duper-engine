package log

import (
	"go-boiler/config"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// referennce : https://pkg.go.dev/github.com/rs/zerolog#section-readme
// initialize log
func InitializeLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	writer := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	log.Logger = log.Output(writer)
	log.Trace().Msg("zerolog initilaized")
}

func ErrorWithStack(err error) {
	log.Error().Msgf("%+v", errors.WithStack(err))
}

func SetLogLevel(config *config.Config) {
	level, err := zerolog.ParseLevel(config.Server.LogLevel)
	if err != nil {
		level = zerolog.TraceLevel
		log.Trace().Str("loglevel", level.String()).Msg("Environment has no log level set up, using default.")
	} else {
		log.Trace().Str("loglevel", level.String()).Msg("Set log level detected.")
	}
	zerolog.SetGlobalLevel(level)
}
