// Code generated by ntaps. DO NOT EDIT.
package infrastructure

import (
	"github.com/geb/aweproj/ses3/shared"
	"log"

	"go.uber.org/dig"

	"fmt"

	_ "github.com/geb/aweproj/ses3/docs"
	swagger "github.com/swaggo/echo-swagger"

	healthcheck_http "github.com/geb/aweproj/ses3/infrastructure/healthcheck"
	inventory_http "github.com/geb/aweproj/ses3/infrastructure/inventory"
)

type (
	Holder struct {
		dig.In

		// - healthcheck-http-start
		HealthcheckController *healthcheck_http.Controller
		// - healthcheck-http-end

		// - inventory-http-start
		InventoryController *inventory_http.Controller
		// - inventory-http-end

		SharedHolder shared.Holder

		// - infrastructure-end
	}
)

func (h *Holder) ListenHttp() {

	RegisterMiddleware(h)

	h.SharedHolder.Echo.GET("/swagger/*", swagger.WrapHandler)

	// - healthcheck-check-http-start
	h.SharedHolder.Echo.GET("/application/health", h.HealthcheckController.Check)
	// - healthcheck-check-http-end

	// - inventory-http-start

	h.SharedHolder.Echo.GET("/inventory/book/find/:pubId", h.InventoryController.FindBookByPubId)

	h.SharedHolder.Echo.POST("/inventory/book/bulk/create", h.InventoryController.BulkCreateBook)

	h.SharedHolder.Echo.PUT("/inventory/book/update/:id", h.InventoryController.UpdateBook)

	h.SharedHolder.Echo.DELETE("/inventory/book/delete/:id", h.InventoryController.DeleteBook)

	// - inventory-http-end

	if err := h.SharedHolder.Echo.Start(fmt.Sprintf(":%d", h.SharedHolder.Config.EchoServerPort)); err != nil {
		if err.Error() == "http: Server closed" {
			log.Printf("closing echo http server")
		} else {
			log.Printf("failed to start echo http server %s", err)
		}
	}

	// - http-end

}
