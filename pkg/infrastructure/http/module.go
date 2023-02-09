package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		func() (*gin.Engine, error) {
			gin.SetMode(gin.ReleaseMode)

			router := gin.New()

			return router, nil
		},
	),
)
