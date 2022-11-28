package inventory

import (
	"context"
	sh "github.com/geb/aweproj/ses3/shared"
	"github.com/labstack/echo/v4"
	"strconv"
)

// All godoc
// @Tags inventory-controller
// @Summary Inventory
// @Description Put all mandatory parameter
// @Param X-Username header string true "guest" default(guest)
// @Param Accept-Language header string true "EN" default(EN)
// @Param pubId path string true "publisherId" default(1)
// @Accept json
// @Produce json
// @Success 200 {object} inventory.BookDto
// @Router /inventory/book/find/{pubId} [get]
func (c *Controller) FindBookByPubId(ec echo.Context) error {
	var (
		ctx      = context.Background()
		pubIdStr = ec.Param("pubId")
	)
	pubId, err := strconv.Atoi(pubIdStr)
	if err != nil {
		return sh.Response(ec, nil, sh.New(sh.BAD_REQUEST, err))
	}

	response, err := c.InterfacesHolder.InventoryViewService.FindBookByPubId(ctx, pubId)

	if err != nil {
		return sh.Response(ec, nil, sh.New(sh.BAD_REQUEST, err))
	}

	return sh.Response(ec, response, nil)
}
