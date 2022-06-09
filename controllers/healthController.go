package controllers

import (
	"api-bootstrap-echo/constants"

	"api-bootstrap-echo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// NewHealthController :
func NewHealthController(config *models.Configuration) *HealthController {
	return &HealthController{Configuration: config}
}

// HealthController :
type HealthController struct {
	Configuration *models.Configuration
}

// Get :
func (controller *HealthController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, models.HealthResponse{
		ComponentName: constants.ServiceName,
		Status:        constants.HealthPass,
		Version:       controller.Configuration.Version,
	})

}
