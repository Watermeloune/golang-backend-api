package router

import (
	"net/http"

	h "github.com/Watermeloune/golang-backend-api/handlers"
	"github.com/labstack/echo"
)

//SetRoutes ...
func SetRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/test", h.TestHandler)
	return e
}
