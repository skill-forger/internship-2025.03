package profile

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	hdl "golang-project/internal/handler"
	"golang-project/internal/middleware"
	svc "golang-project/internal/service"
	"golang-project/server"
)

// handler represents the implementation of handler.Profile
type handler struct {
	route      string
	profileSvc svc.Profile
}

// NewHandler returns a new implementation of handler.Profile
func NewHandler(route string, profileSvc svc.Profile) hdl.Profile {
	return &handler{
		route:      route,
		profileSvc: profileSvc,
	}
}

// RegisterRoutes registers the handler routes and returns the server.HandlerRegistry
func (h *handler) RegisterRoutes() server.HandlerRegistry {
	return server.HandlerRegistry{
		Route: h.route,
		Register: func(group *echo.Group) {
			group.GET("/:userId", h.Get)
			group.GET("/posts", h.ListBloggerPosts, middleware.Authentication(nil))
			group.GET("/posts/:postId", h.GetPostDetail, middleware.Authentication(nil))
			group.PUT("", h.Update, middleware.Authentication(nil))
			group.PUT("/change-password", h.ChangePassword, middleware.Authentication(nil))
		},
	}
}

// Get   handles the profile detail request
//
//	@Summary		Respond profile detail information
//	@Description	Respond profile detail information
//	@Tags			profile
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		int	true	"User ID"
//	@Success		200	{object}	contract.ProfileResponse
//	@Failure		400	{object}	error
//	@Router			/profile/{userId} [get]
func (h *handler) Get(e echo.Context) error {
	// Get user ID from URL param
	userID := e.Param("userId")
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "user ID is required")
	}

	// Convert ID from string to int
	id, err := strconv.Atoi(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	response, err := h.profileSvc.GetByID(id)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, response)
}

// ListBloggerPosts returns all posts (published and draft) of the current blogger
//
//	@Summary      View all blogger's posts
//	@Description  Blogger can view all their posts. Use query parameters to filter (e.g., is_published=false to view drafts).
//	@Tags         profile
//	@Produce      json
//	@Security     BearerToken
//	@Param        is_published  query     bool   false  "Filter by status post"
//	@Success      200  {object}   contract.ListPostResponse
//	@Failure      401  {object}  error
//	@Router       /profile/posts [get]
func (h *handler) ListBloggerPosts(e echo.Context) error {
	return nil
}

// GetPostDetail returns the details of a specific post (published or draft) of the current blogger
//
//	@Summary      View a specific blogger's post
//	@Description  Blogger can view the detail of their own post, whether it's published or draft
//	@Tags         profile
//	@Produce      json
//	@Security     BearerToken
//	@Param        postId  path      int  true  "Post ID"
//	@Success      200     {object}  contract.PostResponse
//	@Failure      404     {object}  error
//	@Router       /profile/posts/{postId} [get]
func (h *handler) GetPostDetail(e echo.Context) error {
	return nil
}

// Update handles the request to update blogger profile information
//
//	@Summary      Update blogger profile
//	@Description  Blogger can update their profile information
//	@Tags         profile
//	@Accept       json
//	@Produce      json
//	@Param        request  body      contract.UpdateProfileRequest  true  "Update profile request"
//	@Security     BearerToken
//	@Success      200      {object}  contract.ProfileResponse
//	@Failure      400      {object}  error
//	@Router       /profile [put]
func (h *handler) Update(e echo.Context) error {
	return nil
}

// ChangePassword handles the request to change blogger password
//
//	@Summary      Change blogger password
//	@Description  Blogger can change their password
//	@Tags         profile
//	@Accept       json
//	@Produce      json
//	@Param        request  body      contract.ChangePasswordRequest  true  "Change password request"
//	@Security     BearerToken
//	@Success      200      {object}  contract.ChangePasswordResponse
//	@Failure      400      {object}  error
//	@Router       /profile/change-password [put]
func (h *handler) ChangePassword(e echo.Context) error {
	return nil
}
