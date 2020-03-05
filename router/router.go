package router

import (
	"net/http"

	"github.com/labstack/echo"
)

//SetRoutes ...
func SetRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}
