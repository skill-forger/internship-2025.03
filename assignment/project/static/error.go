package static

import "errors"

var (
	// Tag errors
	ErrReadTagID   = errors.New("error get tag detail")
	ErrHasPosts    = errors.New("error delete tag because it has associated posts")
	ErrTagNotFound = errors.New("error tag id not found")

	// Favourite errors - User following
	ErrUserNotFound     = errors.New("error user id not found")
	ErrSelfFollow       = errors.New("error cannot follow yourself")
	ErrAlreadyFollowing = errors.New("error already following this user")
	ErrNotFollowing     = errors.New("error not following this user")
	ErrReadUserID       = errors.New("error get user detail")
	ErrReadFollowStatus = errors.New("error get follow status")

	// Post favourites errors - commented out as requested
	// ErrPostNotFound        = errors.New("error post id not found")
	// ErrAlreadyFavourited   = errors.New("error already favourited this post")
	// ErrNotFavourited       = errors.New("error post not in favourites")
	// ErrReadPostID          = errors.New("error get post detail")
	// ErrReadFavouriteStatus = errors.New("error get favourite status")
)
