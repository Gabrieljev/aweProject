package infrastructure

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func RegisterMiddleware(h *Holder) {
	// - register your middleware here
	h.SharedHolder.Echo.Use(CustomMiddleware())

}

func CustomMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			rw := ctx.Response().Writer

			txName := fmt.Sprintf("%s [%s]", ctx.Path(), ctx.Request().Method)

			startTime := time.Now()
			err := next(ctx)
			latency := time.Since(startTime)
			log.Printf("request and response %s, latency %d milliseconds", txName, latency.Milliseconds())
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
