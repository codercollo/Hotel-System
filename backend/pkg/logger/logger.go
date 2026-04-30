// Package logger wraps zerolog and provides a globally accessible structured
// logger. Initialise once at startup via Init; use Log() or the package-level
// helpers (Info, Error, etc.) everywhere else.
package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var log zerolog.Logger

// Init configures the global logger. Call once from main.
// level: "debug" | "info" | "warn" | "error"
// format: "json" | "pretty"
func Init(level, format string) {
	var w io.Writer = os.Stdout

	if format == "pretty" {
		w = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
	}

	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}

	log = zerolog.New(w).
		Level(lvl).
		With().
		Timestamp().
		Caller().
		Logger()

	zerolog.TimeFieldFormat = time.RFC3339
}

// Log returns the global logger for direct use (e.g. adding fields).
func Log() *zerolog.Logger { return &log }

// Debug logs at DEBUG level.
func Debug(msg string) { log.Debug().Msg(msg) }

// Info logs at INFO level.
func Info(msg string) { log.Info().Msg(msg) }

// Warn logs at WARN level.
func Warn(msg string) { log.Warn().Msg(msg) }

// Error logs at ERROR level with an attached error.
func Error(err error, msg string) { log.Error().Err(err).Msg(msg) }

// Fatal logs at FATAL level and exits.
func Fatal(err error, msg string) { log.Fatal().Err(err).Msg(msg) }

// WithRequestID returns a child logger with a request ID field attached.
func WithRequestID(requestID string) zerolog.Logger {
	return log.With().Str("request_id", requestID).Logger()
}

// WithField returns a child logger with an arbitrary string field.
func WithField(key, value string) zerolog.Logger {
	return log.With().Str(key, value).Logger()
}
