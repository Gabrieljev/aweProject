package inventory

import (
	"context"
	sh "github.com/geb/aweproj/ses3/shared"
	inventory_dto "github.com/geb/aweproj/ses3/shared/dto/inventory"
	"github.com/labstack/echo/v4"
	"strconv"
)

// All godoc
// @Tags inventory-controller
// @Summary Inventory
// @Description Put all mandatory parameter
// @Param X-Token header string true "token" default(token)
// @Param Accept-Language header string true "EN" default(EN)
// @Param BookDto body inventory.BookReq true "BookRequest"
// @Param id path string true "objectId"
// @Accept json
// @Produce json
// @Success 200
// @Router /book-store/inventory/book/update/{id} [put]
func (c *Controller) UpdateBook(ec echo.Context) error {
	var (
		ctx     = context.Background()
		request inventory_dto.BookReq
	)

	if err := ec.Bind(&request); err != nil {
		return sh.Response(ec, nil, sh.New(sh.BAD_REQUEST, err))
	}

	idStr := ec.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		return sh.Response(ec, nil, sh.New(sh.BAD_REQUEST, err))
	}

	if err := ec.Validate(&request); err != nil {
		return sh.Response(ec, nil, sh.New(sh.BAD_REQUEST, err))
	}

	err = c.InterfacesHolder.InventoryViewService.UpdateBook(ctx, id, request)

	if err != nil {
		return sh.Response(ec, nil, sh.New(sh.BAD_REQUEST, err))
	}

	return sh.Response(ec, nil, nil)

}
