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
	shopping_http "github.com/geb/aweproj/ses3/infrastructure/shopping"
)

type (
	Holder struct {
		dig.In

		// - healthcheck-http-start
		HealthcheckController *healthcheck_http.Controller
		// - healthcheck-http-end

		// - shopping-http-start
		ShoppingController *shopping_http.Controller
		// - shopping-http-end

		SharedHolder shared.Holder

		// - infrastructure-end
	}
)

func (h *Holder) ListenMessaging() {

}

func (h *Holder) ListenHttp() {

	h.SharedHolder.Echo.GET("/swagger/*", swagger.WrapHandler)

	// - healthcheck-check-http-start
	h.SharedHolder.Echo.GET("/application/health", h.HealthcheckController.Check)
	// - healthcheck-check-http-end

	// - shopping-http-start

	h.SharedHolder.Echo.GET("/shopping/book/find/:pubId", h.ShoppingController.FindBookByPubId)
	h.SharedHolder.Echo.POST("/shopping/book/bulk/create", h.ShoppingController.BulkCreateBook)

	// - shopping-http-end

	if err := h.SharedHolder.Echo.Start(fmt.Sprintf(":%d", h.SharedHolder.Config.EchoServerPort)); err != nil {
		if err.Error() == "http: Server closed" {
			log.Printf("closing echo http server")
		} else {
			log.Printf("failed to start echo http server %s", err)
		}
	}

	// - http-end

}
