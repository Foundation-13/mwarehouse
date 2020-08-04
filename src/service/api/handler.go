package api

import (
	"fmt"
	"github.com/Foundation-13/mwarehouse/src/service/types"
	"github.com/labstack/echo"
	"net/http"


	"github.com/Foundation-13/mwarehouse/src/service/log"
)

func Assemble(e *echo.Echo, m Manager) {
	h := &handler{
		manager: m,
	}

	g := e.Group("/media")

	g.POST("", h.upload)

	g.PUT("/:key/process", h.process)

	g.GET("/:key/status", h.status)
	g.GET("/jobs", h.jobs)
}

type handler struct {
	manager Manager
}

// handlers

func (h *handler) upload(c echo.Context) error {
	ctx := c.Request().Context()

	// Source
	fileHeader, err := c.FormFile("file")
	if err != nil {
		log.FromContext(ctx).WithError(err).Error("failed to get file")
		return err
	}

	f, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	id, err := h.manager.UploadMedia(ctx, f, fileHeader.Filename)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{"id": id})
}

func (h *handler) status(c echo.Context) error {
	ctx := c.Request().Context()

	key := c.Param("key")
	if key == "" {
		return fmt.Errorf("key is empty")
	}

	result, err := h.manager.GetJobStatus(ctx, key)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"key": result.Key,
		"status": fmt.Sprint(result.Status),
	})
}

func (h *handler) process(c echo.Context) error {
	ctx := c.Request().Context()
	key := c.Param("key")

	filters, err := types.UnmarshalFilters(c.Request().Body)
	if err != nil {
		return err
	}

	err = h.manager.ProcessMedia(ctx, key, filters)
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) jobs(c echo.Context) error {
	return fmt.Errorf("not implemented")
}