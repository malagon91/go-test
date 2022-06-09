package main

import (
	"api-bootstrap-echo/controllers"
	"api-bootstrap-echo/libs/logger"
	"api-bootstrap-echo/models"
	"api-bootstrap-echo/repositories"
	"api-bootstrap-echo/server"
	"api-bootstrap-echo/server/di"
)

func main() {

	err := di.GetContainer().Invoke(func(
		configuration *models.Configuration,
		bootstrapRepository *repositories.BootstrapRepository,
		bootstrapController *controllers.BootstrapController,
		healthController *controllers.HealthController) {
		server.Start(configuration, bootstrapController, healthController)
	})
	if err != nil {
		logger.Fatal("main", "main", "di", err.Error())
	}

}
