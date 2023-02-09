package cmd

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/otoru/nezuko/pkg/controllers"
	"github.com/otoru/nezuko/pkg/infrastructure"
	"github.com/otoru/nezuko/pkg/infrastructure/settings"
	"github.com/otoru/nezuko/pkg/services"
	"go.uber.org/dig"
	"go.uber.org/fx"
)

func Execute() error {
	app := fx.New(
		infrastructure.Module,
		services.Module,
		controllers.Module,
		fx.Invoke(func(config *settings.Settings, server *gin.Engine, lc fx.Lifecycle) {
			lc.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						return server.Run(config.Address)
					},
				},
			)
		}),
	)

	ctx := context.Background()
	err := app.Start(ctx)

	return dig.RootCause(err)
}
