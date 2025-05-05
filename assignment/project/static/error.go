package static

import "errors"

var (
	ErrReadTagID   = errors.New("error get tag detail")
	ErrHasPosts    = errors.New("error delete tag because it has associated posts")
	ErrTagNotFound = errors.New("error tag id not found")

	// Post errors
	ErrInvalidPostID     = errors.New("error invalid post id")
	ErrPostTitleRequired = errors.New("error published post required title")
	ErrPostBodyRequired  = errors.New("error published post required body")
	ErrPostInvalidField  = errors.New("error invalid field")

	// Favourite errors - User following
	ErrUserNotFound = errors.New("error user id not found")
)
