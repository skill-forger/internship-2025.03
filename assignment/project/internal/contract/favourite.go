package contract

// BloggerFollowStatusResponse represents the response when a user follows/unfollows a blogger,
// containing the user ID and the current following status
type BloggerFollowStatusResponse struct {
	UserID      int  `json:"user_id"`
	IsFollowing bool `json:"is_following"`
}

// PostFavouriteStatusResponse represents the response when a user marks/unmarks a post as favourite,
// containing the post ID and the current favourite status
type PostFavouriteStatusResponse struct {
	PostID      int  `json:"post_id"`
	IsFavourite bool `json:"is_favourite"`
}