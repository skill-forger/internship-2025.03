package profile

import (
	"net/http"

	"github.com/labstack/echo/v4"

	hdl "golang-project/internal/handler"
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
		Route:           h.route,
		IsAuthenticated: true,
		Register: func(group *echo.Group) {
			group.GET("/:userId", h.Get)
			group.GET("/posts", h.ListPosts)
			group.GET("/posts/:postId", h.GetPost)
			group.PUT("", h.Update)
			group.PUT("/change-password", h.ChangePassword)
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
//	@Security		BearerToken
//	@Success		200	{object}	contract.ProfileResponse
//	@Failure		400	{object}	error
//	@Router			/profile [get]
func (h *handler) Get(e echo.Context) error {
	ctxUser, err := hdl.GetContextUser(e)
	if err != nil {
		return err
	}

	response, err := h.profileSvc.GetByID(ctxUser.ID)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, response)
}

// ListPosts returns all posts (published and draft) of the current blogger
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
func (h *handler) ListPosts(e echo.Context) error {
	return nil
}

// GetPost returns the details of a specific post (published or draft) of the current blogger
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
func (h *handler) GetPost(e echo.Context) error {
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
