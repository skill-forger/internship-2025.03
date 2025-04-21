package tag

import (
	"github.com/labstack/echo/v4"
	ct "golang-project/internal/contract"
	hdl "golang-project/internal/handler"
	svc "golang-project/internal/service"
	"golang-project/server"
	"net/http"
	"strconv"
)

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
		IsAuthenticated: true,
		Register: func(group *echo.Group) {
			group.GET("", h.GetAll)
			group.GET("/:id/posts", h.GetPosts)
			group.POST("", h.Create)
			group.DELETE("/:id", h.Delete)
		},
	}
}

// GetAll handles the request to get all tags
// @Summary     Get all tags
// @Description  Get all blog tags
// @Tags        tag
// @Accept      json
// @Produce     json
// @Success     200 {array}  contract.TagResponse
// @Failure     400 {object} error
// @Router      /tag [get]
func (h *handler) GetAll(e echo.Context) error {
	response, err := h.tagSvc.GetAll()

	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, response)
}

// GetPosts handles the request to get all posts for a tag
// @Summary     Get all posts for a tag
// @Description  Get all blog posts belong to a particular tag
// @Tags        tag
// @Accept      json
// @Produce     json
// @Security    BearerToken
// @Param       id  path     int  true  "Tag ID"
// @Success     200 {array}  contract.PostResponse
// @Failure     400 {object} error
// @Router      /tag/{id}/posts [get]
func (h *handler) GetPosts(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid tag id")
	}

	response, err := h.tagSvc.GetPostsByID(id)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, response)
}

// Create handles the request to create a new tag
// @Summary     Create a new tag
// @Description  Blogger can create new blog tag
// @Tags        tag
// @Accept      json
// @Produce     json
// @Security    BearerToken
// @Param       request body contract.TagRequest true "Tag Request"
// @Success     201 {object} contract.TagResponse
// @Failure     400 {object} error
// @Router      /tag [post]
func (h *handler) Create(e echo.Context) error {

	request := new(ct.TagRequest)
	if err := e.Bind(request); err != nil {
		return err
	}

	response, err := h.tagSvc.Create(request)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, response)
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
// @Router      /tag/{id} [delete]
func (h *handler) Delete(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid tag id")
	}

	err = h.tagSvc.Delete(id)
	if err != nil {
		return err
	}

	return e.NoContent(http.StatusNoContent)
}
