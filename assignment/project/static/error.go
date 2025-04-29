package static

import "errors"

var (
	ErrReadTagID   = errors.New("error get tag detail")
	ErrHasPosts    = errors.New("error delete tag because it has associated posts")
	ErrTagNotFound = errors.New("error tag id not found")

	// Favourite errors - User following
	ErrUserNotFound = errors.New("error user id not found")

	// Post errors
	ErrInsertPost           = errors.New("error creating post")
	ErrTagNotFoundOrDeleted = errors.New("error one or more tags not found or deleted")
	ErrInsertPostTags       = errors.New("error creating post tag")
	ErrFetchPostDetail      = errors.New("error fetching post detail")
)
