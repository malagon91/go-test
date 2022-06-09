package controllers

import (
	"api-bootstrap-echo/models"

	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type mockBootstrapRepository struct {
	mock.Mock
}

// Get :
func (mock *mockBootstrapRepository) Get(id int) (bootstrap models.Bootstrap, err error) {
	args := mock.Called(id)
	return args.Get(0).(models.Bootstrap), args.Error(1)
}

// Save :
func (mock *mockBootstrapRepository) Save(bootstrap models.Bootstrap) (Saved bool, err error) {
	args := mock.Called(bootstrap)
	return args.Get(0).(bool), args.Error(1)
}

// https://stackoverflow.com/questions/45126312/how-do-i-test-an-error-on-reading-from-a-request-body
// https://echo.labstack.com/guide/testing
func TestGetOk(t *testing.T) {
	e := echo.New()
	bootstrap := models.Bootstrap{ID: 123, Title: "Test", Description: "More tests"}
	mockRepository := new(mockBootstrapRepository)

	mockRepository.On("Get", 123).Return(bootstrap, nil)

	bootstrapController := BootstrapController{Repository: mockRepository}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/bootstrap/endpoint/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")

	h := &bootstrapController

	if assert.NoError(t, h.Get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
