package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/Foundation-13/mwarehouse/src/service/api"
	"github.com/Foundation-13/mwarehouse/src/service/aws"
	"github.com/Foundation-13/mwarehouse/src/service/db"
	"github.com/Foundation-13/mwarehouse/src/service/storage"
	"github.com/Foundation-13/mwarehouse/src/service/utils"
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

	stg := storage.NewAWSClient("foundation-13-temporary", aws.S3)
	db := db.NewDynamoDBClient(aws.Dynamo)

	m := api.NewManager(stg, db, utils.XID{})

	api.Assemble(e, m)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "green"})
	})

	e.Logger.Fatal(e.Start(":8765"))
}
