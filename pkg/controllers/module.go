package controllers

import (
	"github.com/otoru/nezuko/pkg/controllers/healthcheck"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(healthcheck.Controller),
)
