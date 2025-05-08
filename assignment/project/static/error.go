package static

import "errors"

var (
	// User Permission errors
	ErrUserPermission = errors.New("error user do not have permission")
	ErrPostOwner      = errors.New("error user do not own the requested post")

	//Profile errors
	ErrListBloggerPosts = errors.New("error retrieving blogger posts")
	ErrParamInvalid     = errors.New("error invalid param")

	// Tags errors
	ErrReadTagID   = errors.New("error get tag detail")
	ErrHasPosts    = errors.New("error delete tag because it has associated posts")
	ErrTagNotFound = errors.New("error tag id not found")

	// Favourite errors - User following
	ErrUserNotFound            = errors.New("error user id not found")
	ErrSelfFollow              = errors.New("error cannot follow yourself")
	ErrDatabaseOperation       = errors.New("error occurred during database operation")
	ErrFollowStatusUpdate      = errors.New("error failed to update follow status")
	ErrUnsupportedFollowAction = errors.New("error unsupported follow action")

	// Favourite errors - Post favourites
	ErrGetFavouritePosts          = errors.New("error retrieving favourite posts")
	ErrGetFollowedBloggerPosts    = errors.New("error retrieving followed blogger posts")
	ErrFavouriteStatusUpdate      = errors.New("error failed to update favourite status")
	ErrUnsupportedFavouriteAction = errors.New("error unsupported favourite action")

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
	ErrCommentNotFound  = errors.New("error comment not found")
	ErrInvalidCommentID = errors.New("error invalid comment id")

	// Change Password errors
	ErrInvalidPassword = errors.New("invalid password")
	ErrComfirmPassword = errors.New("comfirm new passwords do not match")
)
