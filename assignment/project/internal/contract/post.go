package contract

// PostDetailResponse specifies the data and types for post API response
type PostDetailResponse struct {
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
	UserID    int    `json:"user_id,omitempty"`
	Tags      []int  `json:"tags,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// ListPostResponse specifies the data and types for list post API response
type ListPostResponse struct {
	Posts []*PostDetailResponse `json:"posts"`
}

// PostFavouriteStatusResponse represents the response when a user marks/unmarks a post as favourite,
// containing the post ID and the current favourite status
type PostFavouriteStatusResponse struct {
	PostID      int  `json:"post_id"`
	IsFavourite bool `json:"is_favourite"`
}

// PostFavouriteAction defines the possible actions for adding/removing posts from favourites
type PostFavouriteAction int

const (
	Unfavourite PostFavouriteAction = 0
	Favourite   PostFavouriteAction = 1
)

// PostFavouriteRequest represents the request payload for add to/remove from favourites actions
type PostFavouriteRequest struct {
	Action PostFavouriteAction `json:"action" validate:"required,oneof=0 1"`
	PostID int                 `json:"post_id"`
}
