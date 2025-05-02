package authentication

import (
	"golang-project/static"
	"net/http"
	"strings"

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
			group.POST("/sign-up", h.SignUp)
			group.POST("/verify", h.VerifyEmail)
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
	var req ct.SignUpRequest

	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err)
	}

	if len(strings.TrimSpace(req.Email)) == 0 || len(strings.TrimSpace(req.Password)) == 0 || len(strings.TrimSpace(req.LastName)) == 0 || len(strings.TrimSpace(req.FirstName)) == 0 {
		return e.JSON(http.StatusUnprocessableEntity, "All fields are required")
	}

	resp, err := h.authSvc.SignUp(&req)
	if err != nil {
		// Xử lý lỗi cụ thể
		switch err {
		case static.ErrEmailAlreadyExists:
			return e.JSON(http.StatusBadRequest, "Email already exists")
		case static.ErrInvalidEmail:
			return e.JSON(http.StatusBadRequest, "Invalid email format")
		case static.ErrPasswordHashingFailed:
			return e.JSON(http.StatusInternalServerError, "Password hashing failed")
		case static.ErrSaveUserFailed:
			return e.JSON(http.StatusInternalServerError, "Could not save user to database")
		default:
			return e.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	return e.JSON(http.StatusOK, resp)
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
