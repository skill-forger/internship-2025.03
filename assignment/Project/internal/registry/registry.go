package registry

import (
	"github.com/labstack/echo/v4"
	swagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"

	_ "golang-project/docs/swagger"
	"golang-project/internal/handler"
	"golang-project/internal/registry/authentication"
	"golang-project/internal/registry/health"
	"golang-project/internal/registry/profile"
	"golang-project/server"
)

// NewHandlerRegistries returns all server handler registries
func NewHandlerRegistries(db *gorm.DB) ([]server.HandlerRegistry, error) {
	registries := []server.HandlerRegistry{
		initSwaggerRegistry(),
		initHealthCheckHandler(db).RegisterRoutes(),
	}

	for _, hdl := range initResourceHandlers(db) {
		registries = append(registries, hdl.RegisterRoutes())
	}

	return registries, nil
}

// initSwaggerRegistry returns the swagger handler registry
func initSwaggerRegistry() server.HandlerRegistry {
	return server.HandlerRegistry{
		Route: "/swagger",
		Register: func(group *echo.Group) {
			group.GET("/*", swagger.WrapHandler)
		},
	}
}

// initHealthCheckHandler returns the health check handler registry
func initHealthCheckHandler(db *gorm.DB) handler.ResourceHandler {
	return health.NewRegistry("/health", db)
}

// initResourceHandlers returns the service resource handler registry
func initResourceHandlers(db *gorm.DB) []handler.ResourceHandler {
	return []handler.ResourceHandler{
		authentication.NewRegistry("/auth", db),
		profile.NewRegistry("/profile", db),
	}
}
