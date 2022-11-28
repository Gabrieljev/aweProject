package shopping

import (
	"context"
	sh "github.com/geb/aweproj/ses3/shared"
	shopping_dto "github.com/geb/aweproj/ses3/shared/dto/shopping"
	"github.com/labstack/echo/v4"
)

// All godoc
// @Tags shopping-controller
// @Summary Shopping
// @Description Put all mandatory parameter
// @Param X-Username header string true "guest" default(guest)
// @Param Accept-Language header string true "EN" default(EN)
// @Param BookDto body shopping.BookBulkReq true "BulkBookRequest"
// @Accept json
// @Produce json
// @Success 200
// @Router /shopping/book/bulk/create [post]
func (c *Controller) BulkCreateBook(ec echo.Context) error {
	var (
		ctx     = context.Background()
		request shopping_dto.BookBulkReq
	)

	if err := ec.Bind(&request); err != nil {
		return sh.Response(ec, nil, sh.New(sh.BAD_REQUEST, err))
	}

	err := c.InterfacesHolder.ShoppingViewService.BulkCreateBook(ctx, request.BookBulk)

	if err != nil {
		return sh.Response(ec, nil, sh.New(sh.BAD_REQUEST, err))
	}

	return sh.Response(ec, nil, nil)

}
