package contract

import (
	"github.com/golang-jwt/jwt"
)

// CustomClaim represents the authentication custom claim payload
type CustomClaim struct {
	jwt.StandardClaims
	UserID    int    `json:"user_id,omitempty"`
	UserEmail string `json:"user_email,omitempty"`
}

// ContextUser represents the authenticated user in API context
type ContextUser struct {
	ID    int    `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

// SignInRequest represents the request payload for Sign In API
type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// SignInResponse specifies the data and types for Sign In API response
type SignInResponse struct {
	UserID       int    `json:"user_id,omitempty"`
	Token        string `json:"token,omitempty"`
	Type         string `json:"type,omitempty"`
	ExpiredAfter int    `json:"expired_at,omitempty"`
}
