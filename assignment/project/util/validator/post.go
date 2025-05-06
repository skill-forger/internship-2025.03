package validator

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	ct "golang-project/internal/contract"
)

// ValidatePost validates the title and body of a post request.
func ValidatePost(req *ct.CreatePostRequest) error {
	if len(strings.TrimSpace(req.Title)) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Title is required")
	}

	if len(req.Title) > 255 {
		return echo.NewHTTPError(http.StatusBadRequest, "Title is too long (maximum 255 characters)")
	}

	if len(strings.TrimSpace(req.Body)) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Body is required")
	}

	return nil
}
