package shared

import (
	"github.com/geb/aweproj/ses3/shared/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In
		Echo *echo.Echo
		Config *config.EnvConfiguration
	}
)
