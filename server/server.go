package server

import (
	"api-bootstrap-echo/constants"
	"api-bootstrap-echo/controllers"
	"api-bootstrap-echo/libs/logger"
	"api-bootstrap-echo/models"
	"api-bootstrap-echo/server/routes"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var server *echo.Echo

// Start :
func Start(configuration *models.Configuration, bootstrapController *controllers.BootstrapController,
	healthController *controllers.HealthController) {
	initServer()
	setupRoutes(*bootstrapController, *healthController)
	setupMiddleware(configuration)
	setupNotFoundHandler()
	setupMethodNotAllowedHandler()
	setUpPrometheus()
	startServer(configuration)
}

func initServer() {
	server = echo.New()
	server.HideBanner = true
}

func setupRoutes(bootstrapController controllers.BootstrapController,
	healthController controllers.HealthController) {
	routes.RegisterRoutes(bootstrapController, healthController)
	for _, r := range routes.Routes {
		server.Add(r.Method, r.Pattern, r.HandlerFunc).Name = r.Name
	}
}

func setupMiddleware(configuration *models.Configuration) {
	server.Use(
		middleware.Recover(),
		middleware.LoggerWithConfig(logger.GetLoggerConfig()),
		middleware.RequestIDWithConfig(middleware.RequestIDConfig{
			Generator: func() string {
				return uuid.NewString()
			},
		}),
		middleware.GzipWithConfig(middleware.GzipConfig{
			Skipper: func(c echo.Context) bool {
				return c.Path() == constants.MetricsResource
			},
			Level: 5, // enabled by default, remove for short responses
		}),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: configuration.CorsAllowedOrigins,
			AllowMethods: configuration.CorsAllowedMethods,
			AllowHeaders: configuration.CorsAllowedHeaders,
			MaxAge:       7200,
		}),
	)
}

func setupNotFoundHandler() {
	echo.NotFoundHandler = func(c echo.Context) error {
		msg := fmt.Sprintf("Resource not found :%s", c.Request().URL)
		return c.String(http.StatusNotFound, msg)
	}
}

func setupMethodNotAllowedHandler() {
	echo.MethodNotAllowedHandler = func(c echo.Context) error {
		msg := fmt.Sprintf("Method not allowed: %s: %s", c.Request().Method, c.Request().URL)
		return c.String(http.StatusNotFound, msg)
	}
}

func setUpPrometheus() {
	p := prometheus.NewPrometheus("echo", func(c echo.Context) bool {
		return c.Path() == constants.MetricsResource
	})
	p.Use(server)
}

func startServer(configuration *models.Configuration) {
	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", configuration.Port)))
}
