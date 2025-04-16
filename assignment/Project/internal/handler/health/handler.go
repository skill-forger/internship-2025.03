package health

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	ct "golang-project/internal/contract"
	hdl "golang-project/internal/handler"
	"golang-project/server"
)

// handler represents the implementation of handler.ResourceHandler
type handler struct {
	route string
	db    *gorm.DB
}

// NewHandler returns a new implementation of handler.ResourceHandler
func NewHandler(route string, db *gorm.DB) hdl.ResourceHandler {
	return &handler{
		route: route,
		db:    db,
	}
}

// RegisterRoutes registers the handler routes and returns the server.HandlerRegistry
func (h *handler) RegisterRoutes() server.HandlerRegistry {
	return server.HandlerRegistry{
		Route: h.route,
		Register: func(group *echo.Group) {
			group.GET("", h.HealthCheck)
		},
	}
}

// HealthCheck   handles the checking of server and database liveness
//	@Summary		Show server liveness
//	@Description	Perform server and dependent resource liveness check
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	contract.HealthCheckResponse
//	@Failure		400	{object}	error
//	@Router			/health [get]
func (h *handler) HealthCheck(e echo.Context) error {
	response := []*ct.HealthCheckResponse{{Resource: "server", Status: "ok"}}
	response = append(response, h.CheckDatabase())

	return e.JSON(http.StatusOK, response)
}

// CheckDatabase return contract.HealthCheckResponse with the server and database result
func (h *handler) CheckDatabase() *ct.HealthCheckResponse {
	result := &ct.HealthCheckResponse{Resource: "database", Status: "ok"}

	if h.db == nil {
		result.Status = "error: database instance is missing"
		return result
	}

	sqlDB, err := h.db.DB()
	if err != nil {
		result.Status = fmt.Sprintf("error: %s", err)
		return result
	}

	err = sqlDB.Ping()
	if err != nil {
		result.Status = fmt.Sprintf("error: %s", err)
		return result
	}

	return result
}
