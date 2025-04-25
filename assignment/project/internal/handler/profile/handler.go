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
			group.GET("", h.Get)
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

// ListDraftPosts returns the list of draft posts of the current blogger
//
//	@Summary      View all blogger's draft posts
//	@Description  Blogger can view all their draft posts
//	@Tags         profile
//	@Produce      json
//	@Security     BearerToken
//	@Success      200  {object}   contract.ListPostResponse
//	@Failure      401  {object}  error
//	@Router       /profile/drafts [get]
func (h *handler) ListDraftPosts(e echo.Context) error {
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
