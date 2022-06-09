package di

import (
	"api-bootstrap-echo/config"
	"api-bootstrap-echo/controllers"
	"api-bootstrap-echo/repositories"
	"sync"

	"go.uber.org/dig"
)

var (
	container *dig.Container
	once      sync.Once
)

// GetContainer :
func GetContainer() *dig.Container {
	once.Do(func() {
		container = buildContainer()
	})
	return container
}

// BuildContainer :
func buildContainer() *dig.Container {
	container := dig.New()
	handlerContainerErrors(
		container.Provide(config.NewConfiguration),
		container.Provide(repositories.NewBootstrapRepository),
		container.Provide(controllers.NewBootstrapController),
		container.Provide(controllers.NewHealthController))
	return container
}

func handlerContainerErrors(errors ...error) {
	for _, err := range errors {
		if err != nil {
			panic(err)
		}
	}
}
