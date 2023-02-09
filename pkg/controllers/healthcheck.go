package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otoru/nezuko/pkg/services/healthcheck"
	"go.uber.org/zap"
)

func HealthCheck(router *gin.Engine, service *healthcheck.HealthCheck, logger *zap.Logger) {
	router.GET("/healthz", func(ctx *gin.Context) {
		status := http.StatusNoContent

		if err := service.Validate(logger); err != nil {
			status = http.StatusInternalServerError
		}

		ctx.Writer.WriteHeader(status)
	})
}
