package healthcheck

import (
	"context"
	sh "github.com/geb/aweproj/book-store/shared"
	healthcheck_dto "github.com/geb/aweproj/book-store/shared/dto/healthcheck"
	"github.com/labstack/echo/v4"
)

// All godoc
// @Tags healthcheck-controller
// @Summary Healthcheck
// @Description Put all mandatory parameter
// @Param Accept-Language header string true "EN" default(EN)
// @Accept json
// @Produce json
// @Success 200 {object} healthcheck.CheckResponseDto
// @Router /application/health [get]
func (c *Controller) Check(ec echo.Context) error {
	var (
		ctx     = context.Background()
		request healthcheck_dto.CheckRequestDto
	)

	response, err := c.InterfacesHolder.HealthCheckViewService.Check(ctx, request)

	if err != nil {
		return sh.Response(ec, "", nil, sh.New(sh.BAD_REQUEST, err))
	}

	return sh.Response(ec, "", response, nil)
}
