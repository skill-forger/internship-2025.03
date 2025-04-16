package server

import (
	"context"

	"github.com/labstack/echo/v4"
)

// Engine represents the server engine
type Engine interface {
	Address() string
	Startup(handlers ...HandlerRegistry) error
	Shutdown(ctx context.Context) error
}

// engine is an implementation of the Echo Engine
type engine struct {
	address string
	server  *echo.Echo
}

// NewEngine creates and returns an Echo server engine instance
func NewEngine(address string, configs ...ConfigProvider) Engine {
	echoServer := echo.New()

	for _, provide := range configs {
		provide(echoServer)
	}

	return &engine{address: address, server: echoServer}
}

// Address returns the server address
func (e *engine) Address() string {
	return e.address
}

// Startup registers routes handlers and serves the request
func (e *engine) Startup(handlers ...HandlerRegistry) error {
	if e.Address() == "" {
		return ErrMissingServerAddress
	}

	for _, handler := range handlers {
		handler.Register(e.server.Group(handler.Route))
	}

	return e.server.Start(e.Address())
}

// Shutdown cuts the server connection and stops serving request
func (e *engine) Shutdown(ctx context.Context) error {
	if e.server == nil {
		return ErrUninitializedEngine
	}

	return e.server.Shutdown(ctx)
}
