package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	ct "golang-project/internal/contract"
	"golang-project/server"
)

// ResourceHandler             represents all API resource handler
//
//	@title						golang project layout server swagger API
//	@version					1.0
//	@description				This is the swagger API for golang project layout.
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				API Support
//	@contact.url				http://www.swagger.io/support
//	@contact.email				support@swagger.io
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@host						localhost:3000
//	@BasePath					/
//	@securityDefinitions.apikey	BearerToken
//	@in							header
//	@name						Authorization
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
type ResourceHandler interface {
	RegisterRoutes() server.HandlerRegistry
}

// Authentication represents all authentication resource handler
type Authentication interface {
	ResourceHandler
	SignIn(echo.Context) error
	SignUp(echo.Context) error
	VerifyEmail(echo.Context) error
}

// Profile represents all profile resource handler
type Profile interface {
	ResourceHandler
	Get(echo.Context) error
	ListBloggerPosts(echo.Context) error
	GetPostDetail(echo.Context) error
	Update(echo.Context) error
	ChangePassword(echo.Context) error
}

// Tag represents all tag resource handler
type Tag interface {
	ResourceHandler
	Create(echo.Context) error
	Delete(echo.Context) error
	List(echo.Context) error
	ListPosts(echo.Context) error
}

// Post represents all Post resource handler
type Post interface {
	ResourceHandler
	Create(echo.Context) error
	Get(echo.Context) error
	List(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

// Favourite represents all favourite resource handler
type Favourite interface {
	ResourceHandler
	UpdateBlogger(echo.Context) error
	ListBloggers(echo.Context) error
	ListBloggerPosts(echo.Context) error
	UpdatePost(echo.Context) error
	ListPosts(echo.Context) error
}

// Comment represents all comment resource handler
type Comment interface {
	ResourceHandler
	Create(echo.Context) error
	List(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

// GetContextUser returns the authenticated user in echo Context
func GetContextUser(e echo.Context) (*ct.ContextUser, error) {
	ctxUser, ok := e.Get("user").(*ct.ContextUser)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "context user is malformed or missing")
	}

	return ctxUser, nil
}
