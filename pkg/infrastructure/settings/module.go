package settings

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		func() (*Settings, error) {
			viper.AutomaticEnv()

			database := new(Database)
			database.URI = viper.GetString("DATABASE_URI")
			database.Name = viper.GetString("DATABASE_NAME")

			settings := new(Settings)
			settings.Database = database
			settings.Address = viper.GetString("API_ADDRESS")

			return settings, nil
		},
	),
)
