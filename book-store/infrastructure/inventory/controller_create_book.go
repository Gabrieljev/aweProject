package inventory

import (
	"context"
	sh "github.com/geb/aweproj/book-store/shared"
	inventory_dto "github.com/geb/aweproj/book-store/shared/dto/inventory"
	"github.com/labstack/echo/v4"
)

// All godoc
// @Tags inventory-controller
// @Summary Inventory
// @Description Put all mandatory parameter
// @Param X-Token header string true "token" default(token)
// @Param Accept-Language header string true "EN" default(EN)
// @Param BookDto body inventory.BookBulkReq true "BulkBookRequest"
// @Accept json
// @Produce json
// @Success 200
// @Router /book-store/inventory/book/bulk/create [post]
func (c *Controller) BulkCreateBook(ec echo.Context) error {
	var (
		ctx     = context.Background()
		request inventory_dto.BookBulkReq
	)

	if err := ec.Bind(&request); err != nil {
		return sh.Response(ec, "", nil, sh.New(sh.BAD_REQUEST, err))
	}

	err := c.InterfacesHolder.InventoryViewService.BulkCreateBook(ctx, request.BookBulk)

	if err != nil {
		return sh.Response(ec, "", nil, sh.New(sh.BAD_REQUEST, err))
	}

	return sh.Response(ec, "", nil, nil)

}
