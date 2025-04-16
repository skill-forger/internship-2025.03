package server

import (
	"errors"

	"github.com/labstack/echo/v4"
)

var (
	ErrMissingServerAddress = errors.New("server address is missing")
	ErrUninitializedEngine  = errors.New("server engine is not initialized")
)

// ConfigProvider represents the config provider for Echo server
type ConfigProvider func(*echo.Echo)

// HandlerRegistry represents the route group and function that are used by each handler
type HandlerRegistry struct {
	// Route is the route group name of URI
	Route string
	// IsAuthenticated indicates where this route group needs authenticated access
	IsAuthenticated bool
	// Register the function to register the handler for each rout
	Register func(*echo.Group)
}
