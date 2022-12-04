package shared

import (
	"github.com/geb/aweproj/book-store/shared/config"
	"github.com/geb/aweproj/book-store/shared/renderer"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In
		Echo     *echo.Echo
		Config   *config.EnvConfiguration
		Template *renderer.Template
		WebClient WebClientHolder
	}
)
