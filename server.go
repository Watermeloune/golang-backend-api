package main

import (
	"github.com/Watermeloune/golang-backend-api/route"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e = SetRoutes()
	e.Logger.Fatal(e.Start(":1323"))
}
gi