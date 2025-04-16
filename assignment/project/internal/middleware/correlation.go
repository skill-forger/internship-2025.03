package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Correlation provides the middleware for any request and response correlation ID
func Correlation() echo.MiddlewareFunc {
	config := middleware.RequestIDConfig{
		Skipper:      func(c echo.Context) bool { return c.Request().URL.String() == "/health" },
		Generator:    func() string { return uuid.New().String() },
		TargetHeader: echo.HeaderXCorrelationID,
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()

			cid := req.Header.Get(config.TargetHeader)
			if cid == "" {
				cid = config.Generator()
				req.Header.Set(config.TargetHeader, cid)
			}

			res.Header().Set(config.TargetHeader, cid)

			if config.RequestIDHandler != nil {
				config.RequestIDHandler(c, cid)
			}

			return next(c)
		}
	}
}
