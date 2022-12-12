package inventory

import (
	"context"
	"encoding/json"
	sh "github.com/geb/aweproj/book-store/shared"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
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
	var (
		ctx      = context.Background()
		pubIdStr = ec.Param("pubId")
	)
	pubId, err := strconv.Atoi(pubIdStr)
	if err != nil {
		return sh.Response(ec, "", nil, sh.New(sh.BAD_REQUEST, err))
	}
	token := ec.Request().Header.Get("X-Token")

	response, err := c.InterfacesHolder.InventoryViewService.FindBookByPubId(ctx, pubId, token)

	if err != nil {
		return sh.Response(ec, "", nil, sh.New(sh.BAD_REQUEST, err))
	}
	var data []map[string]interface{}
	jsonResult, _ := json.Marshal(response)
	json.Unmarshal(jsonResult, &data)

	log.Println(data)

	return sh.Response(ec, "show_book", map[string]interface{}{
		"dataBook": data,
	}, nil)
}
