package database

import (
	"context"

	"github.com/otoru/nezuko/pkg/infrastructure/settings"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		func(lc fx.Lifecycle, config *settings.Settings) (*Client, error) {
			client, err := NewClient(config.Database.URI)

			if err != nil {
				return nil, err
			}

			lc.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						return client.Connect(ctx)
					},
					OnStop: func(ctx context.Context) error {
						return client.Disconnect(ctx)
					},
				},
			)

			return client, nil
		},
	),
)
