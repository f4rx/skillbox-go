package misc

import (
	logger "github.com/f4rx/logger-zap-wrapper"
	"go.uber.org/zap"
)

var Slog *zap.SugaredLogger //nolint:gochecknoglobals

func NewLogger() {
	Slog = logger.NewSugaredLogger()
	Slog.Sync() //nolint:errcheck
}
