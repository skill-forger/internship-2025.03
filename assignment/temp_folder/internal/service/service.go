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
