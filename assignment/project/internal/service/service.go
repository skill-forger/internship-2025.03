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
	ListComments(*ct.ListCommentRequest) (*ct.ListCommentResponse, error)
}
