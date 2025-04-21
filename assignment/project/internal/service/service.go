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
	GetAll() ([]*ct.TagResponse, error)
	GetByID(int) (*ct.TagResponse, error)
	GetPostsByID(int) ([]*ct.PostResponse, error)
	Create(request *ct.TagRequest) (*ct.TagResponse, error)
	Delete(int) error
}
