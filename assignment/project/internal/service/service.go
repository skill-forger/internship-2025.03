package service

import (
	"errors"
	ct "golang-project/internal/contract"
)

var (
	ErrNotFound               = errors.New("Not found")
	ErrTagAssociatedWithPosts = errors.New("cannot delete tag that is associated with posts")
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
}
