package static

import "errors"

var (
	// Tags errors
	ErrReadTagID   = errors.New("error get tag detail")
	ErrHasPosts    = errors.New("error delete tag because it has associated posts")
	ErrTagNotFound = errors.New("error tag id not found")
	// Favourite errors - User following
	ErrUserNotFound       = errors.New("error user id not found")
	ErrSelfFollow         = errors.New("error cannot follow yourself")
	ErrAlreadyFollowing   = errors.New("error already following this user")
	ErrNotFollowing       = errors.New("error not following this user")
	ErrDatabaseOperation  = errors.New("error occurred during database operation")
	ErrFollowStatusUpdate = errors.New("error failed to update follow status")

	// SignUp errors
	ErrEmailAlreadyExists    = errors.New("error email already exists")
	ErrInvalidEmail          = errors.New("error invalid email format")
	ErrPasswordHashingFailed = errors.New("error hashing password")
	ErrSaveUserFailed        = errors.New("error saving user to database")
	ErrInvalidName           = errors.New("error invalid name format")
)
