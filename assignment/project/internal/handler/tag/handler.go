package tag

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	ct "golang-project/internal/contract"
	hdl "golang-project/internal/handler"
	svc "golang-project/internal/service"
	"golang-project/server"
)

// handler represents the implementation of hdl.Tag
type handler struct {
	route  string
	tagSvc svc.Tag
}

// NewHandler returns a new implementation of hdl.Tag
func NewHandler(route string, tagSvc svc.Tag) hdl.Tag {
	return &handler{
		route:  route,
		tagSvc: tagSvc,
	}
}

// RegisterRoutes registers the handler routes and returns the server.HandlerRegistry
func (h *handler) RegisterRoutes() server.HandlerRegistry {
	return server.HandlerRegistry{
		Route: h.route,
		Register: func(group *echo.Group) {
			group.GET("", h.GetAll)
			group.POST("", h.Create)
		},
	}
}

// Create handles the request to create a tag
// @Summary     Create a new tag
// @Description Create a new tag with the provided name
// @Tags        tag
// @Accept      json
// @Produce     json
// @Param       request  body     contract.CreateTagRequest  true  "Create Tag Request"
// @Success     200      {object} contract.TagDetailResponse "Tag created successfully"
// @Failure     400      {object} string                    "Invalid request"
// @Failure     422      {object} string                    "Unprocessable entity"
// @Router      /tags [post]
func (h *handler) Create(c echo.Context) error {
	var req ct.CreateTagRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid request")
	}

	if len(strings.TrimSpace(req.Name)) == 0 {
		return c.JSON(http.StatusUnprocessableEntity, "Name is required")
	}

	createdTag, err := h.tagSvc.Create(req.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Unable to create tag")
	}

	// Return created tag
	return c.JSON(http.StatusOK, createdTag)
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
