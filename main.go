package main

import (
	"github.com/labstack/echo"
	"github.com/Watermeloune/golang-backend-api/route"
)

func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
}
