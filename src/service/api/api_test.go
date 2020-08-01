package api_test

import (
	"bytes"
	"github.com/Foundation-13/mwarehouse/src/service/types"
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/buger/jsonparser"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Foundation-13/mwarehouse/src/service/api"
	"github.com/Foundation-13/mwarehouse/src/service/api/apimocks"
)

func TestStatus(t *testing.T) {
	t.Run("key is empty", func(t *testing.T) {
		a := newApi()

		a.r.GET("/media//status").SetDebug(true).
			Run(a.e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusInternalServerError, r.Code)
		})
	})

	t.Run("succeeded", func(t *testing.T) {
		a := newApi()

		a.m.On("GetJobStatus", mock.Anything, mock.Anything).Return(types.JobStatus(0), nil)

		a.r.GET("/media/123/status").SetDebug(true).
			Run(a.e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)

				assertID(t, "123", r.Body, "key")
				assertID(t, "0", r.Body, "status")
		})
	})
}

func TestUpload(t *testing.T) {
	a := newApi()

	a.m.On("UploadMedia", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("123", nil)

	a.r.POST("/media").
		SetDebug(true).
		SetFileFromPath([]gofight.UploadFile{
			{
				Path:    "/images/media.png",
				Name:    "file",
				Content: []byte("123"),
			},
		}).
		Run(a.e, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusCreated, r.Code)

			assertID(t, "123", r.Body, "id")
		})
}

// helpers

type apiMocks struct {
	r *gofight.RequestConfig
	e *echo.Echo
	m *apimocks.Manager
}

func newApi() apiMocks {
	r := gofight.New()

	e := echo.New()
	m := &apimocks.Manager{}

	api.Assemble(e, m)

	return apiMocks{
		r: r,
		e: e,
		m: m,
	}
}

func assertID(t assert.TestingT, expected string, body *bytes.Buffer, keys ...string) {
	data := body.Bytes()
	id, _ := jsonparser.GetString(data, keys...)
	assert.Equal(t, expected, id)
}