package post

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"golang-project/internal/contract"
	hdl "golang-project/internal/handler"
	"golang-project/internal/middleware"
	svc "golang-project/internal/service"
	"golang-project/server"
	"golang-project/static"
)

// handler represents the implementation of handler.Post
type handler struct {
	route   string
	postSvc svc.Post
}

// NewHandler returns a new implementation of handler.Post
func NewHandler(route string, postSvc svc.Post) hdl.Post {
	return &handler{
		route:   route,
		postSvc: postSvc,
	}
}

func (h *handler) RegisterRoutes() server.HandlerRegistry {
	return server.HandlerRegistry{
		Route: h.route,
		Register: func(group *echo.Group) {
			group.GET("", h.List)
			group.GET("/:postId", h.Get)
			group.POST("", h.Create, middleware.Authentication(nil))
			group.PUT("/:postId", h.Update, middleware.Authentication(nil))
			group.DELETE("/:postId", h.Delete, middleware.Authentication(nil))
		},
	}
}

// Get handles the post detail request
//
//	@Summary		Respond post detail information
//	@Description	Respond post detail information
//	@Tags			post
//	@Accept			json
//	@Produce		json
//	@Param			postId	path		int	true	"Post ID"
//	@Success		200		{object}	contract.PostResponse
//	@Failure		400		{object}	error
//	@Router			/posts/{postId} [get]
func (h *handler) Get(e echo.Context) error {
	// Get post ID from URL param
	postID := e.Param("postId")
	if postID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Post ID is required")
	}

	// Convert ID from string to int
	id, err := strconv.Atoi(postID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid post ID")
	}

	// Get data from service
	response, err := h.postSvc.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Post not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving post")
	}

	return e.JSON(http.StatusOK, response)
}

// List handles the request to retrieve all published posts
//
//	@Summary		View all published posts
//	@Description	Reader/Blogger can view all published posts and filter by specific condition (e.g. tag, author)
//	@Tags			post
//	@Accept			json
//	@Produce		json
//	@Param			filter	query		contract.ListPostRequest	false	"Filtering parameters"
//	@Success		200		{object}	contract.ListPostResponse
//	@Failure		400		{object}	error
//	@Router			/posts [get]
func (h *handler) List(e echo.Context) error {
	// Get filter from query params
	filter := contract.ListPostRequest{}
	if err := e.Bind(&filter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid filter parameters")
	}

	// Set default pagination if not provided
	if filter.Page <= 0 {
		filter.Page = static.Pagination.DefaultPage
	}
	if filter.PageSize <= 0 {
		filter.PageSize = static.Pagination.DefaultPageSize
	}

	// Get data from service
	response, err := h.postSvc.List(&filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving posts")
	}

	return e.JSON(http.StatusOK, response)
}

// Create handles the request to create a new post
//
//	@Summary		Create a new post
//	@Description	Blogger can create a new post (default as draft)
//	@Tags			post
//	@Accept			json
//	@Produce		json
//	@Security		BearerToken
//	@Param			request	body		contract.CreatePostRequest	true	"Create post request"
//	@Success		200		{object}	contract.PostResponse
//	@Failure		400		{object}	error
//	@Router			/posts [post]
func (h *handler) Create(e echo.Context) error {
	return nil
}

// Update handles the request to update an existing post
//
//	@Summary		Update an existing post
//	@Description	Blogger can update their post content and toggle publish/draft status
//	@Tags			post
//	@Accept			json
//	@Produce		json
//	@Security		BearerToken
//	@Param			postId	path		int							true	"Post ID"
//	@Param			request	body		contract.UpdatePostRequest	true	"Update post request"
//	@Success		200		{object}	contract.PostResponse
//	@Failure		400		{object}	error
//	@Router			/posts/{postId} [put]
func (h *handler) Update(e echo.Context) error {
	var req contract.UpdatePostRequest

	postId, err := strconv.Atoi(e.Param("postId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, static.ErrInvalidPostID)
	}

	if err := e.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if len(strings.TrimSpace(req.Title)) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, static.ErrPostTitleRequired)
	}

	if len(strings.TrimSpace(req.Body)) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, static.ErrPostBodyRequired)
	}

	updatePost, err := h.postSvc.Update(postId, &req, req.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return e.JSON(http.StatusOK, updatePost)
}

// Delete handles the request to delete a post
//
//	@Summary		Delete a post
//	@Description	Blogger can delete their own post
//	@Tags			post
//	@Accept			json
//	@Produce		json
//	@Security		BearerToken
//	@Param			postId	path		int	true	"Post ID"
//	@Success		204		{object}	nil
//	@Failure		400		{object}	error
//	@Router			/posts/{postId} [delete]
func (h *handler) Delete(e echo.Context) error {
	return nil
}
