package service

import (
	ct "golang-project/internal/contract"
)

// Authentication represents the service logic of Authentication
type Authentication interface {
	SignIn(*ct.SignInRequest) (*ct.SignInResponse, error)
	SignUp(*ct.SignUpRequest) (*ct.SignUpResponse, error)
}

// Profile represents the service logic of Profile
type Profile interface {
	GetByID(int) (*ct.ProfileResponse, error)
	Update(int, *ct.UpdateProfileRequest) (*ct.ProfileResponse, error)
}

type Tag interface {
	Create(string) (*ct.TagResponse, error)
	Delete(int) error
	List() (*ct.ListTagResponse, error)
	ListPosts(int) (*ct.ListPostResponse, error)
}

type Comment interface {
	Create(*ct.CreateCommentRequest, int) (*ct.CommentResponse, error)
	List(*ct.ListCommentRequest) (*ct.ListCommentResponse, error)
}

type Post interface {
	GetByID(int) (*ct.PostResponse, error)
	List(*ct.ListPostRequest) (*ct.ListPostResponse, error)
	Create(*ct.CreatePostRequest, int) (*ct.PostResponse, error)
	Update(int, *ct.UpdatePostRequest, []int) (*ct.PostResponse, error)
}

// Favourite represents the service logic of Favourite features
type Favourite interface {
	// User following operations
	UpdateFollowStatus(userID int, req *ct.BloggerFollowRequest) (*ct.BloggerFollowStatusResponse, error)
	ListFollowingUsers(userID int) (*ct.ListProfileResponse, error)
	ListUserPosts(userID int) (*ct.ListPostResponse, error)
	// Post favorite operations
	Favourite(userID, postID int, isFavourite bool) (*ct.PostFavouriteStatusResponse, error)
	ListFavouritePosts(userID int) (*ct.ListPostResponse, error)
}
