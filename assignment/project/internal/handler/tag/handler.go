package tag

import (
	"net/http"
	"strconv"
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
		Route:           h.route,
		IsAuthenticated: true,
		Register: func(group *echo.Group) {
			group.POST("", h.Create)
			group.GET("", h.List)
			group.GET("/:id/posts", h.ListPosts)
			group.DELETE("/:id", h.Delete)
		},
	}
}

// List handles the request to get all tags
// @Summary     Get all tags
// @Description  Readers/Bloggers can view all blog tags
// @Tags        tag
// @Accept      json
// @Produce     json
// @Success     200 {object} contract.ListTagResponse
// @Failure     400 {object} error
// @Router      /tags [get]
func (h *handler) List(e echo.Context) error {
	response, err := h.tagSvc.List()

	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, response)
}

// ListPosts handles the request to get all posts for a tag
// @Summary     Get all posts for a tag
// @Description  Readers/Bloggers can view all blog posts belong to a particular tag
// @Tags        tag
// @Accept      json
// @Produce     json
// @Param       tagId  path     int  true  "Tag ID"
// @Success     200 {object}  contract.ListPostResponse
// @Failure     400 {object} error
// @Router      /tags/:tagId/posts [get]
func (h *handler) ListPosts(e echo.Context) error {
	return nil
}

// Create handles the request to create a tag
//
//	@Summary		Create a new tag
//	@Description	Create a new tag with the provided name
//	@Tags			tag
//	@Accept			json
//	@Produce		json
//	@Param			request	body		contract.CreateTagRequest	true	"Create Tag Request"
//	@Success		200		{object}	contract.TagDetailResponse	"Tag created successfully"
//	@Failure		400		{object}	string						"Invalid request"
//	@Failure		422		{object}	string						"Unprocessable entity"
//	@Router			/tags [post]
func (h *handler) Create(c echo.Context) error {
	var req ct.CreateTagRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
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

// Delete handles the request to delete a tag
// @Summary     Delete a tag
// @Description  Blogger can delete a tag that does not contain any blog
// @Tags        tag
// @Accept      json
// @Produce     json
// @Security    BearerToken
// @Param       id  path     int  true  "Tag ID"
// @Success     204 "No Content"
// @Failure     400 {object} error
// @Router      /tags/{id} [delete]
func (h *handler) Delete(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid tag ID")
	}

	err = h.tagSvc.Delete(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Unable to delete tag")
	}

	return e.NoContent(http.StatusNoContent)
}
