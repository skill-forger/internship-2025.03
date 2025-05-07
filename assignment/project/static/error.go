package static

import "errors"

var (
	// User Permission errors
	ErrUserPermission = errors.New("error user do not have permission")

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

	// Favourite errors - Post favourites
	ErrGetFavouritePosts    = errors.New("error retrieving favourite posts")
	
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
	ErrPostNotFound         = errors.New("error post not found")
	ErrInvalidPostID        = errors.New("error invalid post id")

	// Comment errors
	ErrCommentNotFound = errors.New("error comment not found")

	// Change Password errors
	ErrInvalidPassword = errors.New("invalid password")
	ErrComfirmPassword = errors.New("comfirm new passwords do not match")
)
