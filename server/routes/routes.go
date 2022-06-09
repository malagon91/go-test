package routes

import (
	"api-bootstrap-echo/constants"
	"api-bootstrap-echo/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Route :
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc echo.HandlerFunc
}

// ServiceRoutes :
type ServiceRoutes []Route

// Routes :
var Routes ServiceRoutes

// RegisterRoutes :
func RegisterRoutes(bootstrapController controllers.BootstrapController, healthController controllers.HealthController) {
	Routes = ServiceRoutes{
		Route{
			Method:      http.MethodGet,
			Name:        "SampleGet",
			Pattern:     "/" + constants.ServiceName + "/endpoint/:id",
			HandlerFunc: bootstrapController.Get,
		},
		Route{
			Method:      http.MethodPost,
			Name:        "SamplePost",
			Pattern:     "/" + constants.ServiceName + "/endpoint",
			HandlerFunc: bootstrapController.Post,
		},
		Route{
			Method:      http.MethodGet,
			Name:        "HealthGet",
			Pattern:     constants.HealthResource,
			HandlerFunc: healthController.Get,
		},
	}
}
