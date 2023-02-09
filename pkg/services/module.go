package services

import (
	"github.com/otoru/nezuko/pkg/infrastructure/database"
	"github.com/otoru/nezuko/pkg/services/healthcheck"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(func(database *database.Client) *healthcheck.HealthCheck {
		instance := healthcheck.New()

		instance.Append(
			healthcheck.Runner{
				Name:       "database",
				Evaluation: database.Ping,
			},
		)

		return instance
	}),
)
