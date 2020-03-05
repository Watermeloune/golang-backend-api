package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func TestHandler(c echo.Context) error {
	return c.String(http.StatusOK, "TEST")
}
