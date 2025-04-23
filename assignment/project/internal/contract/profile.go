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

// ListProfileResponse contains the list of profiles that the current user is following
type ListProfileResponse struct {
	Bloggers []*ProfileResponse `json:"bloggers,omitempty"`
}

// BloggerFollowStatusResponse represents the response when a user follows/unfollows a blogger,
// containing the user ID and the current following status
type BloggerFollowStatusResponse struct {
	UserID      int  `json:"user_id"`
	IsFollowing bool `json:"is_following"`
}
