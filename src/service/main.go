package main

import (
	"fmt"
	stg2 "github.com/Foundation-13/mwarehouse/src/service/storage"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/Foundation-13/mwarehouse/src/service/api"
	"github.com/Foundation-13/mwarehouse/src/service/aws"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	aws, err := aws.NewClient()
	if err != nil {
		panic(err)
	}

	fmt.Printf("AWS opened !!!")

	stg := stg2.NewAWSClient("foundation-13-temporary", aws.S3)

	m := api.NewManager(stg)

	api.Assemble(e, m)

	e.GET("/", func(c echo.Context) error {
		c.JSON(http.StatusOK, map[string]string{"status": "green"})
		return nil
	})

	e.Logger.Fatal(e.Start(":8765"))
}
