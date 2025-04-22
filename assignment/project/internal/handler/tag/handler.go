package tag

import (
	"net/http"

	"github.com/labstack/echo/v4"

	hdl "golang-project/internal/handler"
	svc "golang-project/internal/service"
	"golang-project/server"
)

// handler represents the implementation of handler.Tag
type handler struct {
	route  string
	tagSvc svc.Tag
}

// NewHandler returns a new implementation of handler.Tag
func NewHandler(route string, tagSvc svc.Tag) hdl.Tag {
	return &handler{
		route:  route,
		tagSvc: tagSvc,
	}
}

// RegisterRoutes registers the handler routes and returns the server.HandlerRegistry
func (h *handler) RegisterRoutes() server.HandlerRegistry {
	return server.HandlerRegistry{
		Route:           h.route,
		IsAuthenticated: false,
		Register: func(group *echo.Group) {
			group.GET("", h.GetAll)
		},
	}
}

// GetAll handles the request to get all tags
// @Summary     Get all tags
// @Description  Get all blog tags
// @Tags        tag
// @Accept      json
// @Produce     json
// @Success     200 {object} contract.ListTagResponse
// @Failure     400 {object} error
// @Router      /tags [get]
func (h *handler) GetAll(e echo.Context) error {
	response, err := h.tagSvc.GetAllTags()

	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, response)
}
