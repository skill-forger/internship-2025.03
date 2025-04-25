package authentication

import (
	"net/http"

	"github.com/labstack/echo/v4"

	ct "golang-project/internal/contract"
	hdl "golang-project/internal/handler"
	svc "golang-project/internal/service"
	"golang-project/server"
)

// handler represents the implementation of handler.Authentication
type handler struct {
	route   string
	authSvc svc.Authentication
}

// NewHandler returns a new implementation of handler.Authentication
func NewHandler(route string, authSvc svc.Authentication) hdl.Authentication {
	return &handler{
		route:   route,
		authSvc: authSvc,
	}
}

// RegisterRoutes registers the handler routes and returns the server.HandlerRegistry
func (h *handler) RegisterRoutes() server.HandlerRegistry {
	return server.HandlerRegistry{
		Route: h.route,
		Register: func(group *echo.Group) {
			group.POST("/sign-in", h.SignIn)
		},
	}
}

// SignIn handles the authentication request via predefined credentials
//
//	@Summary		Signs In user into the system
//	@Description	Authenticates user via predefined credentials and return JWT Token
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			SignInRequest	body		ct.SignInRequest	true "Sign In Request Payload"
//	@Success		200				{array}		ct.SignInResponse
//	@Failure		400				{object}	error
//	@Router			/auth/sign-in [post]
func (h *handler) SignIn(e echo.Context) error {
	request := new(ct.SignInRequest)
	if err := e.Bind(request); err != nil {
		return err
	}

	response, err := h.authSvc.SignIn(request)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, response)
}

// SignUp handles the request to register a new user
//
//	@Summary      Register a new user
//	@Description  Reader can sign up to become a blogger
//	@Tags         authentication
//	@Accept       json
//	@Produce      json
//	@Param        request  body      ct.SignUpRequest  true  "Sign up request"
//	@Success      200      {object}  ct.SignUpResponse
//	@Failure      400      {object}  error
//	@Router       /auth/sign-up [post]
func (h *handler) SignUp(e echo.Context) error {
	return nil
}

// VerifyEmail handles the request to verify email address
//
//	@Summary		Verify email address
//	@Description	Blogger can verify their email address upon signing up
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ct.VerifyEmailRequest	true	"Email verification request"
//	@Success		200		{object}	ct.VerifyEmailResponse
//	@Failure		400		{object}	error
//	@Router			/auth/verify [post]
func (h *handler) VerifyEmail(e echo.Context) error {
	return nil
}
