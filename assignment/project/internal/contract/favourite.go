package contract

// BloggerFollowStatusResponse represents the response when a user follows/unfollows a blogger,
// containing the user ID and the current following status
type BloggerFollowStatusResponse struct {
	UserID      int  `json:"user_id"`
	IsFollowing bool `json:"is_following"`
}

// BloggerListResponse contains the list of bloggers that the current user is following
type BloggerListResponse struct {
	Bloggers []ProfileResponse `json:"bloggers,omitempty"`
}

// PostFavouriteStatusResponse represents the response when a user marks/unmarks a post as favourite,
// containing the post ID and the current favourite status
type PostFavouriteStatusResponse struct {
	PostID      int  `json:"post_id"`
	IsFavourite bool `json:"is_favourite"`
}

// PostListResponse contains the list of posts, used for both:
// - Posts from followed bloggers (GetFollowedBloggersPosts)
// - Posts marked as favourite by the current user (GetFavouritePosts)
type PostListResponse struct {
	Posts []PostDetailResponse `json:"posts,omitempty"`
}
