package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/Foundation-13/mwarehouse/src/service/api"
	"github.com/Foundation-13/mwarehouse/src/service/aws"
	"github.com/Foundation-13/mwarehouse/src/service/config"
	"github.com/Foundation-13/mwarehouse/src/service/db"
	"github.com/Foundation-13/mwarehouse/src/service/log"
	"github.com/Foundation-13/mwarehouse/src/service/storage"
	"github.com/Foundation-13/mwarehouse/src/service/utils"
)

func main() {
	cfg, err := config.FromEnvironment()
	if err != nil {
		panic(err)
	}

	log.InitLog(cfg.LocalRun)

	ctx := context.Background()
	logger := log.FromContext(ctx)

	logger.Infof("service starting, region = %s, temp-bucket-name = %s", cfg.Region, cfg.TempBucketName)

	e := echo.New()
	e.Use(middleware.Recover())

	aws, err := aws.NewClient(cfg.Region)
	if err != nil {
		logger.WithError(err).Panic("failed to create aws client")
	}

	logger.Infof("AWS opened !!!")

	stg := storage.NewAWSClient(cfg.TempBucketName, aws.S3)
	db := db.NewDynamoDBClient(aws.Dynamo)

	m := api.NewManager(stg, db, utils.XID{})

	api.Assemble(e, m)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "green"})
	})

	logger.Fatal(e.Start(":8765"))
}
