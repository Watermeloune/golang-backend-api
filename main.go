package main

import (
	"github.com/labstack/echo"
	"github.com/Watermeloune/golang-backend-api/router"
)

func main() {
	e := echo.New()
	e = router.SetRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
