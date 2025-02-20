package kit

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

// Compatibility check.
var _ logging.Logger = &Logger{}

// Logger is a go-kit/log logging adapter compatible with logging middlewares.
type Logger struct {
	log.Logger
}

// InterceptorLogger converts go-kit/log logger to Logger adapter.
func InterceptorLogger(logger log.Logger) *Logger {
	return &Logger{logger}
}

// Log implements logging.Logger interface.
func (l *Logger) Log(lvl logging.Level, msg string) {
	switch lvl {
	case logging.DEBUG:
		_ = level.Debug(l.Logger).Log(msg)
	case logging.INFO:
		_ = level.Info(l.Logger).Log(msg)
	case logging.WARNING:
		_ = level.Warn(l.Logger).Log(msg)
	case logging.ERROR:
		_ = level.Error(l.Logger).Log(msg)
	default:
		panic(fmt.Sprintf("kit: unknown level %s", lvl))
	}
}

// Log implements logging.Logger interface.
func (l *Logger) With(fields ...string) logging.Logger {
	vals := make([]interface{}, 0, len(fields))
	for _, v := range fields {
		vals = append(vals, v)
	}
	return InterceptorLogger(log.With(l.Logger, vals...))
}
