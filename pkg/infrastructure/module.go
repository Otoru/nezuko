package infrastructure

import (
	"github.com/otoru/nezuko/pkg/infrastructure/database"
	"github.com/otoru/nezuko/pkg/infrastructure/http"
	"github.com/otoru/nezuko/pkg/infrastructure/logger"
	"github.com/otoru/nezuko/pkg/infrastructure/settings"
	"go.uber.org/fx"
)

var Module = fx.Options(
	logger.Module,
	settings.Module,
	database.Module,
	http.Module,
)
