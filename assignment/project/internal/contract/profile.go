package contract

// ProfileResponse specifies the data and types for profile API response
type ProfileResponse struct {
	ID           int    `json:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email,omitempty"`
	Pseudonym    string `json:"pseudonym,omitempty"`
	ProfileImage string `json:"profile_image,omitempty"`
	Biography    string `json:"biography,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

// UpdateProfileRequest defines the payload for updating a user's profile information.
type UpdateProfileRequest struct {
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Pseudonym    string `json:"pseudonym,omitempty"`
	ProfileImage string `json:"profile_image,omitempty"`
	Biography    string `json:"biography,omitempty"`
}

// ChangePasswordRequest defines the payload required to change a user's password.
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8"`
}
