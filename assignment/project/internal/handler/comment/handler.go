package comment

import (
	"github.com/labstack/echo/v4"

	hdl "golang-project/internal/handler"
	svc "golang-project/internal/service"
	"golang-project/server"
)

// handler represents the implementation of hdl.Comment
type handler struct {
	route      string
	commentSvc svc.Comment
}

// NewHandler returns a new implementation of hdl.Comment
func NewHandler(route string, commentSvc svc.Comment) hdl.Comment {
	return &handler{
		route:      route,
		commentSvc: commentSvc,
	}
}

// RegisterRoutes registers the handler routes and returns the server.HandlerRegistry
func (h *handler) RegisterRoutes() server.HandlerRegistry {
	return server.HandlerRegistry{
		Route: h.route,
		Register: func(group *echo.Group) {
			group.GET("", h.List)
			group.POST("", h.Create)
			group.PUT("/:commentId", h.Update)
			group.DELETE("/:commentId", h.Delete)
		},
	}
}

// List handles the request to get all comments for a post
// @Summary     Get all comments for a post
// @Description  Reader/Blogger can view all comments in the blog posts
// @Tags        comment
// @Accept      json
// @Produce     json
// @Param       filter  query  contract.ListCommentRequest  false	"Filtering parameters"
// @Success     200 {object} contract.ListCommentResponse
// @Failure     400 {object} error
// @Router      /comments [get]
func (h *handler) List(e echo.Context) error {
	return nil
}

// Update handles the request to update a comment
// @Summary     update a comment
// @Description Blogger can update their comment
// @Tags        comment
// @Accept      json
// @Produce     json
// @Security    BearerToken
// @Param       request body contract.UpdateCommentRequest true "Update Comment Request"
// @Param       commentId  path     int  true  "Comment ID"
// @Success     200 {object} contract.CommentResponse
// @Failure     400 {object} error
// @Router      /comments/{commentId} [put]
func (h *handler) Update(e echo.Context) error {
	return nil
}

// Create handles the request to create a new comment
// @Summary     Create a new comment
// @Description  Blogger can make a new comment/ Blogger can reply to another comment (with parentCMTID in request body)
// @Tags        comment
// @Accept      json
// @Produce     json
// @Security    BearerToken
// @Param       request body contract.CreateCommentRequest true "Create Comment Request"
// @Success     200 {object} contract.CommentResponse
// @Failure     400 {object} error
// @Router      /comments [post]
func (h *handler) Create(e echo.Context) error {
	return nil
}

// Delete handles the request to delete a comment
// @Summary     Delete a comment
// @Description  Blogger can delete their comment
// @Tags        comment
// @Accept      json
// @Produce     json
// @Security    BearerToken
// @Param       commentId  path     int  true  "comment ID"
// @Success     204 "No Content"
// @Failure     400 {object} error
// @Router      /comments/{commentId} [delete]
func (h *handler) Delete(e echo.Context) error {
	return nil
}
