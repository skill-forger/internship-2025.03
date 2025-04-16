package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	ct "golang-project/internal/contract"
	"golang-project/server"
	"golang-project/static"
)

// Authentication provides the middleware for any API requires user authentication
func Authentication(registries []server.HandlerRegistry) echo.MiddlewareFunc {
	pathSkipper := mapPathSkipper(registries)

	return echoJwt.WithConfig(echoJwt.Config{
		Skipper: func(c echo.Context) bool {
			return pathSkipper[getRouteGroup(c.Request().URL.Path)]
		},
		SigningKey:    []byte(viper.GetString(static.EnvAuthSecret)),
		SigningMethod: echoJwt.AlgorithmHS256,
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			keyFunc := func(token *jwt.Token) (interface{}, error) {
				return []byte(viper.GetString(static.EnvAuthSecret)), nil
			}

			token, err := jwt.ParseWithClaims(auth, &ct.CustomClaim{}, keyFunc)
			if err != nil {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			claim, ok := token.Claims.(*ct.CustomClaim)
			if !ok || !token.Valid {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "parse jwt custom claim failed")
			}

			return &ct.ContextUser{ID: claim.UserID, Email: claim.UserEmail}, nil
		},
	})
}

func mapPathSkipper(registries []server.HandlerRegistry) map[string]bool {
	result := map[string]bool{"/": true, "/favicon.ico": true}

	for _, r := range registries {
		if r.IsAuthenticated {
			continue
		}
		result[r.Route] = true
	}

	return result
}

func getRouteGroup(path string) string {
	paths := strings.Split(path, "/")
	if len(paths) < 2 {
		return ""
	}

	return fmt.Sprintf("/%s", paths[1])
}
