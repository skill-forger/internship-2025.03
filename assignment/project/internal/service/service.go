package service

import (
	ct "golang-project/internal/contract"
)

// Authentication represents the service logic of Authentication
type Authentication interface {
	SignIn(*ct.SignInRequest) (*ct.SignInResponse, error)
}

// Profile represents the service logic of Profile
type Profile interface {
	GetByID(int) (*ct.ProfileResponse, error)
}

type Tag interface {
	Create(string) (*ct.TagDetailResponse, error)
	Delete(int) error
	List() (*ct.ListTagResponse, error)
	ListPosts(int) (*ct.ListPostResponse, error)
}

type Comment interface {
}

// Favourite represents the service logic of Favourite features
type Favourite interface {
	// User following operations
	FollowUser(userID, targetUserID int, isFollow bool) (*ct.BloggerFollowStatusResponse, error)
	ListFollowers(userID int) (*ct.ListProfileResponse, error)
	ListFollowerPosts(userID int) (*ct.ListPostResponse, error)

	// Post favorite operations
	FavouritePost(userID, postID int, isFavourite bool) (*ct.PostFavouriteStatusResponse, error)
	ListFavourites(userID int) (*ct.ListPostResponse, error)
}