package inventory

import (

	sh "github.com/geb/aweproj/book-store/shared"
	"github.com/labstack/echo/v4"

)

func (c *Controller) LandingPage(ec echo.Context) error {

	return sh.Response(ec,"landing_page", nil, nil)
}
