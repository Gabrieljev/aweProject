package inventory

import (
	"context"
	sh "github.com/geb/aweproj/book-store/shared"
	"github.com/labstack/echo/v4"
	"strconv"
)

// All godoc
// @Tags inventory-controller
// @Summary Inventory
// @Description Put all mandatory parameter
// @Param X-Token header string true "token" default(token)
// @Param Accept-Language header string true "EN" default(EN)
// @Param id path string true "objectId"
// @Accept json
// @Produce json
// @Success 200
// @Router /book-store/inventory/book/delete/{id} [delete]
func (c *Controller) DeleteBook(ec echo.Context) error {
	var (
		ctx = context.Background()
	)

	idStr := ec.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		return sh.Response(ec, "",nil, sh.New(sh.BAD_REQUEST, err))
	}

	err = c.InterfacesHolder.InventoryViewService.DeleteBook(ctx, id)

	if err != nil {
		return sh.Response(ec, "",nil, sh.New(sh.BAD_REQUEST, err))
	}

	return sh.Response(ec, "",nil, nil)

}
