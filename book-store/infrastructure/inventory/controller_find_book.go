package inventory

import (
	sh "github.com/geb/aweproj/book-store/shared"
	"github.com/labstack/echo/v4"
)

// All godoc
// @Tags inventory-controller
// @Summary Inventory
// @Description Put all mandatory parameter
// @Param X-Token header string true "token" default(token)
// @Param Accept-Language header string true "EN" default(EN)
// @Param pubId path string true "publisherId" default(1)
// @Accept json
// @Produce json
// @Success 200 {object} inventory.BookDto
// @Router /book-store/inventory/book/find/{pubId} [get]
func (c *Controller) FindBookByPubId(ec echo.Context) error {
	//var (
	//	ctx      = context.Background()
	//	pubIdStr = ec.Param("pubId")
	//)
	//pubId, err := strconv.Atoi(pubIdStr)
	//if err != nil {
	//	return sh.Response(ec, "",nil, sh.New(sh.BAD_REQUEST, err))
	//}
	//
	//response, err := c.InterfacesHolder.InventoryViewService.FindBookByPubId(ctx, pubId)
	//
	//if err != nil {
	//	return sh.Response(ec, "",nil, sh.New(sh.BAD_REQUEST, err))
	//}

	return sh.Response(ec,"show_book", nil, nil)
}
