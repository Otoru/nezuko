package logger

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(
		func(lc fx.Lifecycle) (*zap.Logger, error) {
			logger, err := zap.NewProduction()

			lc.Append(
				fx.Hook{
					OnStop: func(ctx context.Context) error {
						return logger.Sync()
					},
				},
			)

			return logger, err
		},
	),
	fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: logger}
	}),
)
