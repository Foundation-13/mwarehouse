package api_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/buger/jsonparser"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Foundation-13/mwarehouse/src/service/api"
	"github.com/Foundation-13/mwarehouse/src/service/api/apimocks"
	"github.com/Foundation-13/mwarehouse/src/service/types"
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

		a.m.On("GetJobStatus", mock.Anything, mock.Anything).Return(types.Job{
			Status: types.JobStatus(0),
			Key:    "123",
		}, nil)

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

func TestProcess(t *testing.T) {
	t.Run("succeeded", func(t *testing.T) {
		a := newApi()

		a.m.On("ProcessMedia", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		body := `
		{
			"filters": [
				{
					"name": "monochrome"
				},
				{
					"name": "scale",
					"percentage": 70
				}
			] 	
		}
		`

		a.r.PUT("/media/123/process").
			SetDebug(true).
			SetBody(body).
			Run(a.e, func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, resp.Code)

				checker := func(f types.Filters) bool {
					if len(f.Filters) != 2 {
						return false
					}

					f1 := f.Filters[0] == types.NamedFilter{Name: "monochrome"}
					f2 := f.Filters[1] == types.FilterWithPercentage{
						NamedFilter: types.NamedFilter{Name: "scale"},
						Percentage:  70,
					}
					return f1 && f2
				}

				a.m.AssertCalled(t, "ProcessMedia", mock.Anything, "123", mock.MatchedBy(checker))
			})
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
