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
	h.SharedHolder.Echo.Use(CustomMiddlewareLatency())
	h.SharedHolder.Echo.Use(CustomMiddlewareAuth())

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
			userName := ctx.Request().Header.Get("X-Username")
			// do your authentication level here
			_ = next(ctx)

			log.Printf("you are %s", userName)
			return
		}
	}
}
