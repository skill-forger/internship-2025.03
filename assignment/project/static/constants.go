package static

// BloggerFollowAction defines action types for follow/unfollow operations
type BloggerFollowAction string

const (
	Follow BloggerFollowAction = "follow"
	Unfollow BloggerFollowAction = "unfollow"
)

// PostFavouriteAction defines action types for favourite/unfavourite operations
type PostFavouriteAction string

const (
	Favourite PostFavouriteAction = "favourite"
	Unfavourite PostFavouriteAction = "unfavourite"
)
