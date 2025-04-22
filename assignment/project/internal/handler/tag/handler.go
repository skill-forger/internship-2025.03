package tag

import (
	"github.com/labstack/echo/v4"
	ct "golang-project/internal/contract"
	hdl "golang-project/internal/handler"
	svc "golang-project/internal/service"
	"golang-project/server"
	"net/http"
	"strings"
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
// @Router      /tags [post]
func (h *handler) Create(c echo.Context) error {
	var req ct.CreateTagRequest
	// Bind request body to the struct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//check
	if req.Name == "" || len(strings.TrimSpace(req.Name)) == 0 {
		return c.JSON(http.StatusBadRequest, "Name is required")
	}

	// Call the service to create the tag
	createdTag, err := h.tagSvc.Create(req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Unable to create tag")
	}

	// Return created tag
	return c.JSON(http.StatusOK, createdTag)
}
