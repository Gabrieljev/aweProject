package infrastructure

import (
	"fmt"
	"github.com/labstack/echo/v4"
	swagger "github.com/swaggo/echo-swagger"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func ProtectedAPI(h *Holder) {
	// - register your middleware here

	protectedApi := h.SharedHolder.Echo.Group("/book-store/inventory")
	{
		protectedApi.Use(CustomMiddlewareAuth())

		// - inventory-http-start

		protectedApi.GET("/book/landing_page", h.InventoryController.LandingPage)

		protectedApi.GET("/book/find/:pubId", h.InventoryController.FindBookByPubId)

		protectedApi.POST("/book/bulk/create", h.InventoryController.BulkCreateBook)

		protectedApi.PUT("/book/update/:id", h.InventoryController.UpdateBook)

		protectedApi.DELETE("/book/delete/:id", h.InventoryController.DeleteBook)

		// - inventory-http-end
	}

}

func RendererPage(h *Holder) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	h.SharedHolder.Echo.Static("/", path+"/shared/renderer/assets")

	h.SharedHolder.Template.Add("login", template.Must(template.ParseFiles(path+"/shared/renderer/member-renderer/index.html")))

	h.SharedHolder.Template.Add("show_book", template.Must(template.ParseFiles(path+"/shared/renderer/inventory-renderer/find_book/index.html")))

	h.SharedHolder.Template.Add("landing_page", template.Must(template.ParseFiles(path+"/shared/renderer/inventory-renderer/landing_page/index.html")))

	h.SharedHolder.Echo.Renderer = h.SharedHolder.Template
}

func PublicAPI(h *Holder) {

	h.SharedHolder.Echo.GET("/swagger/*", swagger.WrapHandler)

	// - healthcheck-check-http-start
	h.SharedHolder.Echo.GET("/application/health", h.HealthcheckController.Check)
	// - healthcheck-check-http-end

	// - member-http-start
	h.SharedHolder.Echo.GET("/book-store/member/user/login", h.MemberController.Login)

	h.SharedHolder.Echo.POST("/book-store/member/user/isUserValid", h.MemberController.IsUserValid)

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

			token, err := ctx.Cookie("token")

			token1 := ctx.Request().Header.Get("X-Token")

			if err != nil && token1 == "" {
				return ctx.Redirect(http.StatusTemporaryRedirect, "/book-store/member/user/login")
			}
			if token != nil{
				ctx.Request().Header.Add("X-Token", token.Value)
			}
			_ = next(ctx)
			return
		}
	}
}
