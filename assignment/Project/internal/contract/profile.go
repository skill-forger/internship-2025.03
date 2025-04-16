package contract

// ProfileResponse specifies the data and types for profile API response
type ProfileResponse struct {
	ID           int    `json:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email,omitempty"`
	Pseudonym    string `json:"display_name,omitempty"`
	ProfileImage string `json:"profile_image,omitempty"`
	Biography    string `json:"biography,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}
