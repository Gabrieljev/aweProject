package infrastructure

import (
	"fmt"
	"github.com/geb/aweproj/ses3/shared"
	"github.com/geb/aweproj/ses3/shared/dto"
	"github.com/labstack/echo/v4"
	swagger "github.com/swaggo/echo-swagger"
	"log"
	"net/http"
	"strconv"
	"time"
)

func ProtectedAPI(h *Holder) {
	// - register your middleware here

	protectedApi := h.SharedHolder.Echo.Group("/book-store")
	{
		protectedApi.Use(CustomMiddlewareAuth())
		inventory := protectedApi.Group("/inventory")
		// - inventory-http-start

		inventory.GET("/book/find/:pubId", h.InventoryController.FindBookByPubId)

		inventory.POST("/book/bulk/create", h.InventoryController.BulkCreateBook)

		inventory.PUT("/book/update/:id", h.InventoryController.UpdateBook)

		inventory.DELETE("/book/delete/:id", h.InventoryController.DeleteBook)

		// - inventory-http-end
	}

}

func PublicAPI(h *Holder) {

	h.SharedHolder.Echo.GET("/swagger/*", swagger.WrapHandler)

	// - healthcheck-check-http-start
	h.SharedHolder.Echo.GET("/application/health", h.HealthcheckController.Check)
	// - healthcheck-check-http-end

	// - member-http-start

	h.SharedHolder.Echo.POST("/book-store/member/user/login", h.MemberController.Login)

	h.SharedHolder.Echo.POST("/book-store/member/user/create", h.MemberController.CreateUser)
	// - member-http-end

}

func CustomMiddlewareLatency() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			rw := ctx.Response().Writer

			txName := fmt.Sprintf("path %s method [%s]", ctx.Path(), ctx.Request().Method)
			// before
			startTime := time.Now()
			err := next(ctx)
			// after
			latency := time.Since(startTime)
			log.Printf("%s, latency %d milliseconds", txName, latency.Milliseconds())
			if err != nil {
				if httperr, ok := err.(*echo.HTTPError); ok {
					rw.WriteHeader(httperr.Code)
				} else {
					rw.WriteHeader(http.StatusInternalServerError)
				}
			}
			return err
		}
	}
}

func CustomMiddlewareAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			token := ctx.Request().Header.Get("X-Token")

			// do your authentication level here
			username, err := shared.ClaimToken(token)
			if err != nil {
				badReq := strconv.Itoa(http.StatusBadRequest)
				return ctx.JSON(http.StatusBadRequest, &dto.BaseResponseDto{
					Code:       badReq,
					Message:    badReq,
					Data:       nil,
					Errors:     []string{err.Error()},
					ServerTime: time.Now().Unix(),
				})
			}
			_ = next(ctx)

			log.Printf("you are %s", username)
			return
		}
	}
}
