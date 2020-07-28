package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"

	"github.com/Foundation-13/mwarehouse/src/service/api"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	m := api.NewManager()

	api.Assemble(e, m)

	e.GET("/", func(c echo.Context) error {
		c.JSON(http.StatusOK, map[string]string{"status": "green"})
		return nil
	})

	e.Logger.Fatal(e.Start(":8765"))
}
