package static

import "errors"

var (
	// Tags errors
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

	// SignUp errors
	ErrEmailAlreadyExists    = errors.New("error email already exists")
	ErrInvalidEmail          = errors.New("error invalid email format")
	ErrPasswordHashingFailed = errors.New("error hashing password")
	ErrSaveUserFailed        = errors.New("error saving user to database")
	ErrInvalidName           = errors.New("error invalid name format")
	ErrCheckEmailFailed      = errors.New("error checking email failed")

	// Post errors
	ErrInsertPost           = errors.New("error creating post")
	ErrTagNotFoundOrDeleted = errors.New("error one or more tags not found or deleted")
	ErrInsertPostTags       = errors.New("error creating post tag")
	ErrFetchPostDetail      = errors.New("error fetching post detail")
)
